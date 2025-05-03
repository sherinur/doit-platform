package dto

import (
	svc "github.com/sherinur/doit-platform/apis/gen/quiz-service/service/frontend/quiz/v1"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ToQuizFromCreateRequest(req *svc.CreateQuizRequest) (model.Quiz, error) {
	return model.Quiz{
		Title:       req.Title,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
		Status:      req.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func FromQuizToCreateResponse(quiz model.Quiz) (*svc.CreateQuizResponse, error) {
	return &svc.CreateQuizResponse{
		CreatedId: quiz.ID,
	}, nil
}

func FromQuizToGetResponse(quiz model.Quiz) (*svc.GetQuizResponse, error) {
	response := &svc.GetQuizResponse{
		Id:          quiz.ID,
		Title:       quiz.Title,
		Description: quiz.Description,
		CreatedAt:   timestamppb.New(quiz.CreatedAt),
		UpdatedAt:   timestamppb.New(quiz.UpdatedAt),
		Status:      quiz.Status,
		TotalPoints: quiz.TotalPoints,
	}

	for _, question := range quiz.Questions {
		que, _ := FromQuestionToGetResponse(question)
		response.Questions = append(response.Questions, que)
	}

	return response, nil
}

func ToQuizFromUpdateRequest(req *svc.UpdateQuizRequest) (model.Quiz, error) {
	quiz := req.Quiz
	var response model.Quiz

	response.ID = quiz.Id
	response.Title = quiz.Title
	response.Description = quiz.Description
	response.CreatedBy = quiz.CreatedBy
	response.Status = quiz.Status
	response.TotalPoints = quiz.TotalPoints
	response.UpdatedAt = quiz.UpdatedAt.AsTime()

	return response, nil
}

func FromQuizToUpdateResponse(quiz model.Quiz) (*svc.UpdateQuizResponse, error) {
	return &svc.UpdateQuizResponse{UpdatedId: quiz.ID}, nil
}

func FromQuizToDeleteResponse(quiz model.Quiz) (*svc.DeleteQuizResponse, error) {
	return &svc.DeleteQuizResponse{DeletedId: quiz.ID}, nil
}
