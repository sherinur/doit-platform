//rmv file: app.go

package app

import (
	"context"
	"course-service/internal/adapters/repo/mongo"
	"course-service/internal/app/logger"
	"course-service/internal/usecase"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type App struct {
	// DB                *mongo.Database
	Logger *zap.Logger
	// CourseUsecase     *usecase.CourseUsecase
	// TagUsecase        *usecase.TagUsecase
	// CategoryUsecase   *usecase.CategoryUsecase
	// InstructorUsecase *usecase.InstructorUsecase
}

// TODO:
// remove usecase from app structure
// remove db from app structure
// add config to app structure
// add metrics(otel, prometheus) to app structure

func NewApp() *App {
	// Initialize logger
	zapLogger := logger.NewZapLogger()

	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatalf("failed to create mongo client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("failed to connect to mongo: %v", err)
	}

	db := client.Database("course_service")

	// Initialize Repositories
	courseRepo := mongoRepo.NewCourseRepository(db)
	tagRepo := mongoRepo.NewTagRepository(db)
	categoryRepo := mongoRepo.NewCategoryRepository(db)
	instructorRepo := mongoRepo.NewInstructorRepository(db)

	// Initialize Usecases
	courseUC := usecase.NewCourseUsecase(courseRepo)
	tagUC := usecase.NewTagUsecase(tagRepo)
	categoryUC := usecase.NewCategoryUsecase(categoryRepo)
	instructorUC := usecase.NewInstructorUsecase(instructorRepo)

	return &App{
		DB:                db,
		Logger:            zapLogger,
		CourseUsecase:     courseUC,
		TagUsecase:        tagUC,
		CategoryUsecase:   categoryUC,
		InstructorUsecase: instructorUC,
	}
}
