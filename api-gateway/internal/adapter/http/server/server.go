package server

import (
	"fmt"
	"net/http"

	"github.com/sherinur/doit-platform/api-gateway/config"
	"github.com/sherinur/doit-platform/api-gateway/internal/adapter/http/server/handler"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const serverIPAddress = "0.0.0.0:%d"

type API struct {
	server *gin.Engine
	cfg    config.HTTPServer
	addr   string
	log    *zap.Logger

	userHandler *handler.User
	fileHandler *handler.File
}

func New(cfg config.HTTPServer, logger *zap.Logger, fileUsecase FileUsecase, userUsecase UserUsecase) *API {
	gin.SetMode(cfg.Mode)
	server := gin.New()
	server.Use(gin.Recovery())

	// Binding presenter
	userHandler := handler.NewUser(userUsecase)

	api := &API{
		server:      server,
		cfg:         cfg,
		addr:        fmt.Sprintf(serverIPAddress, cfg.Port),
		userHandler: userHandler,
		log:         logger,
	}

	api.setupRoutes()

	return api
}

func (a *API) setupRoutes() {
	a.server.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	v1 := a.server.Group("/api/v1")
	{
		file := v1.Group("/file")
		{
			file.PUT("/", a.fileHandler.Create)
			file.GET("/:key", a.fileHandler.Get)
			file.DELETE("/:key", a.fileHandler.Delete)
		}
		user := v1.Group("/user-service")
		{
			user.POST("/register", a.userHandler.Register)
			user.POST("/login", a.userHandler.Login)
			user.POST("/refresh-token", a.userHandler.RefreshToken)
			user.POST("/logout", a.userHandler.Logout)
			user.PUT("/update-profile", a.userHandler.UpdateUserInfo)
			user.PUT("/update-password", a.userHandler.UpdateUserPassword)
			user.GET("getallusers", a.userHandler.GetAllUsers)
			user.POST("/send-verify-code", a.userHandler.SendVerificationCode)
			user.POST("/verify-email", a.userHandler.VerifyEmail)
		}
	}
}

func (a *API) Run() error {
	a.log.Info("Running http server", zap.String("addr", a.addr), zap.String("gin mode", a.cfg.Mode))
	return a.server.Run(a.addr)
}
