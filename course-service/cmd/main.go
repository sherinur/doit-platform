package main

import (
	"net"
	"os"

	grpcController "course-service/internal/adapters/controller/grpc"
	"course-service/internal/adapters/repo/mongo"
	"course-service/internal/app"
	"course-service/internal/usecase"
	"course-service/pkg/mongo"

	grpcService "course-service/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	app.InitLogger()
	logger := app.Log
	logger.Info("Starting Course Service")

	mongoClient := mongo.Connect(os.Getenv("MONGO_URI"))
	db := mongoClient.Database("course_service")

	courseRepo := mongoRepo.NewCourseRepository(db)
	courseUsecase := usecase.NewCourseUsecase(courseRepo)
	grpcHandler := grpcController.NewCourseHandler(courseUsecase)

	grpcServer := grpc.NewServer()
	grpcService.RegisterCourseServiceServer(grpcServer, grpcHandler)

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		logger.Fatal("Failed to listen on port 50053", zap.Error(err))
	}
	logger.Info("gRPC server listening on port 50053")

	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatal("Failed to serve gRPC server", zap.Error(err))
	}
}
