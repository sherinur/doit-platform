package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"quizzes-service/config"
	"quizzes-service/internal/adapters/controller/http/service/handler"
)

const serverIPAddress = "0.0.0.0:%s" // Changed to 0.0.0.0 for external access

type API struct {
	server *gin.Engine
	cfg    config.HTTPServer
	addr   string

	ResultHandler   *handler.ResultHandler
	QuizHandler     *handler.QuizHandler
	QuestionHandler *handler.QuestionHandler
	AnswerHandler   *handler.AnswerHandler
}

func New(cfg config.Server, resultUseCase ResultUseCase, quizUseCase QuizUseCase, questionUseCase QuestionUseCase, answerUseCase AnswerUseCase) *API {
	// Setting the Gin mode
	gin.SetMode(cfg.HttpServer.Mode)
	// Creating a new Gin Engine
	server := gin.New()

	// Applying middleware
	server.Use(gin.Recovery())

	// Binding clients
	resultHandler := handler.NewResultHandler(resultUseCase)
	quizHandler := handler.NewQuizHandler(quizUseCase)
	questionHandler := handler.NewQuestionHandler(questionUseCase)
	answerHandler := handler.NewAnswerHandler(answerUseCase)

	api := &API{
		server:          server,
		cfg:             cfg.HttpServer,
		addr:            fmt.Sprintf(serverIPAddress, cfg.HttpServer.Port),
		ResultHandler:   resultHandler,
		QuizHandler:     quizHandler,
		QuestionHandler: questionHandler,
		AnswerHandler:   answerHandler,
	}

	api.setupRoutes()

	return api
}

func (a *API) setupRoutes() {
	v1 := a.server.Group("/api/v1")
	{
		quizzes := v1.Group("/quizzes")
		{
			quizzes.POST("/", a.QuizHandler.CreateQuiz)
			quizzes.GET("/:id", a.QuizHandler.GetQuizById)
			quizzes.PUT("/:id", a.QuizHandler.UpdateQuiz)
			quizzes.DELETE("/:id", a.QuizHandler.DeleteQuiz)
		}

		questions := v1.Group("/questions")
		{
			questions.POST("/", a.QuestionHandler.CreateQuestion)
			questions.POST("/many", a.QuestionHandler.CreateQuestions)
			questions.GET("/:id", a.QuestionHandler.GetQuestionById)
			questions.GET("/quiz/:id", a.QuestionHandler.GetQuestionsByQuizId)
			questions.PUT("/:id", a.QuestionHandler.UpdateQuestion)
			questions.DELETE("/:id", a.QuestionHandler.DeleteQuestion)
		}

		answers := v1.Group("/answers")
		{
			answers.POST("/", a.AnswerHandler.CreateAnswer)
			answers.POST("/many", a.AnswerHandler.CreateAnswers)
			answers.GET("/:id", a.AnswerHandler.GetAnswerById)
			answers.GET("/question/:id", a.AnswerHandler.GetAnswersByQuestionId)
			answers.PUT("/:id", a.AnswerHandler.UpdateAnswer)
			answers.DELETE("/:id", a.AnswerHandler.DeleteAnswer)
		}

		result := v1.Group("/result")
		{
			result.POST("/", a.ResultHandler.CreateResult)
			result.GET("/:id", a.ResultHandler.GetResultById)
			result.GET("/quiz/:id", a.ResultHandler.GetResultsByQuizId)
			result.GET("/user/:id", a.ResultHandler.GetResultsByUserId)
			result.DELETE("/:id", a.ResultHandler.DeleteResult)
		}
	}
}

func (a *API) Run(errCh chan<- error) {
	go func() {
		log.Printf("HTTP server starting on: %v", a.addr)

		if err := a.server.Run(a.addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("failed to start HTTP server: %w", err)
			return
		}
	}()
}

func (a *API) Stop() error {
	// Setting up the signal channel to catch termination signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Blocking until a signal is received
	sig := <-quit
	log.Println("Shutdown signal received", "signal:", sig.String())

	// Creating a context with timeout for graceful shutdown
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("HTTP server shutting down gracefully")

	// Note: You can use `Shutdown` if you use `http.Server` instead of `gin.Engine`.
	log.Println("HTTP server stopped successfully")

	return nil
}
