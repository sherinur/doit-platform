package security

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sherinur/doit-platform/user-service/internal/domain/model"
)

type JWTManager struct {
	jwtAccessSecret      []byte
	jwtRefreshSecret     []byte
	jwtAccessExpiration  int
	jwtRefreshExpiration int
}

func NewJWTManager(accessSecret string, refreshSecret string, accessExpiration int, refreshExpiration int) *JWTManager {
	return &JWTManager{
		jwtAccessSecret:      []byte(accessSecret),
		jwtRefreshSecret:     []byte(refreshSecret),
		jwtAccessExpiration:  accessExpiration,
		jwtRefreshExpiration: refreshExpiration,
	}
}

func (s *JWTManager) GenerateTokens(accessPayload jwt.MapClaims, refreshPayload jwt.MapClaims) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessPayload)
	accessTokenStr, err := accessToken.SignedString(s.jwtAccessSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshPayload)
	refreshTokenStr, err := refreshToken.SignedString(s.jwtRefreshSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenStr, refreshTokenStr, nil
}

func (s *JWTManager) CreateAccessPayload(user *model.User) jwt.MapClaims {
	return jwt.MapClaims{
		"role":    user.Role,
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Second * time.Duration(s.jwtAccessExpiration)).Unix(),
	}
}

func (s *JWTManager) CreateRefreshPayload(user *model.User) jwt.MapClaims {
	return jwt.MapClaims{
		"role":    user.Role,
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Second * time.Duration(s.jwtRefreshExpiration)).Unix(),
	}
}

func (s *JWTManager) ValidateAccessToken(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.jwtAccessSecret, nil
	})

	return err == nil && token.Valid
}

func (s *JWTManager) ParseRefreshToken(tokenStr string) *model.User {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.jwtRefreshSecret, nil
	})
	if err != nil {
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	}

	if claims["user_id"] == nil || claims["role"] == nil {
		return nil
	}

	return &model.User{
		ID:   int64(claims["user_id"].(float64)),
		Role: claims["role"].(string),
	}
}

func (s *JWTManager) ExtractUserIDAndRole(tokenStr string, isRefresh bool) (int64, string, error) {
	var secret []byte
	if isRefresh {
		secret = s.jwtRefreshSecret
	} else {
		secret = s.jwtAccessSecret
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, "", errors.New("invalid token claims")
	}

	userIDFloat, okID := claims["user_id"].(float64)
	role, okRole := claims["role"].(string)
	if !okID || !okRole {
		return 0, "", errors.New("user_id or role not found in token")
	}

	return int64(userIDFloat), role, nil
}
