package security

import (
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
