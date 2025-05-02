package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
)

type QuestionUsecase struct {
	quizRepo     QuizRepo
	questionRepo QuestionRepo
}

func NewQuestionUsecase(qrepo QuizRepo, querepo QuestionRepo) *QuestionUsecase {
	return &QuestionUsecase{quizRepo: qrepo, questionRepo: querepo}
}

func (uc QuestionUsecase) CreateQuestion(ctx context.Context, request model.Question) (model.Question, error) {
	if request.Text == "" || request.Type == "" || request.QuizID == "" || request.Points <= 0 {
		return model.Question{}, errors.New("invalid input data")
	}

	res, err := uc.questionRepo.CreateQuestion(ctx, request)
	if err != nil {
		return model.Question{}, err
	}

	return res, nil
}

func (uc QuestionUsecase) CreateQuestions(ctx context.Context, request []model.Question) ([]model.Question, error) {
	for _, question := range request {
		if question.Text == "" || question.Type == "" || question.QuizID == "" || question.Points <= 0 {
			return nil, errors.New("invalid input data")
		}
	}

	res, err := uc.questionRepo.CreateQuestions(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc QuestionUsecase) GetQuestionById(ctx context.Context, id string) (model.Question, error) {
	question, err := uc.questionRepo.GetQuestionById(ctx, id)
	if err != nil {
		return model.Question{}, err
	}

	return question, nil
}

func (uc QuestionUsecase) GetQuestionsByQuizId(ctx context.Context, id string) ([]model.Question, error) {
	questions, err := uc.questionRepo.GetQuestionsByQuizId(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("wrong QuizID or quiz does not exist: %w", err)
	}

	return questions, nil
}

func (uc QuestionUsecase) UpdateQuestion(ctx context.Context, request model.Question) (model.Question, error) {
	err := uc.questionRepo.UpdateQuestion(ctx, request)
	if err != nil {
		return model.Question{}, err
	}

	return model.Question{
		ID: request.ID,
	}, nil
}

func (uc QuestionUsecase) DeleteQuestion(ctx context.Context, id string) (model.Question, error) {
	err := uc.questionRepo.DeleteQuestion(ctx, id)
	if err != nil {
		return model.Question{}, err
	}

	return model.Question{ID: id}, nil
}
