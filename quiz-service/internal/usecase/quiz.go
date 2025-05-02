package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
)

type QuizUsecase struct {
	quizRepo     QuizRepo
	questionRepo QuestionRepo
}

func NewQuizUsecase(qrepo QuizRepo, queRepo QuestionRepo) *QuizUsecase {
	return &QuizUsecase{quizRepo: qrepo, questionRepo: queRepo}
}

func (uc QuizUsecase) CreateQuiz(ctx context.Context, request model.Quiz) (model.Quiz, error) {
	if request.Title == "" || request.Description == "" || request.CreatedBy == "" || request.Status == "" {
		return model.Quiz{}, errors.New("invalid input data")
	}

	res, err := uc.quizRepo.CreateQuiz(ctx, request)
	if err != nil {
		return model.Quiz{}, err
	}

	return res, nil
}

func (uc QuizUsecase) GetQuizById(ctx context.Context, id string) (model.Quiz, error) {
	quiz, err := uc.quizRepo.GetQuizById(ctx, id)
	if err != nil {
		return model.Quiz{}, err
	}

	totalPoints := 0.0
	questions, err := uc.questionRepo.GetQuestionsByQuizId(ctx, id)
	if err != nil {
		return model.Quiz{}, err
	}

	for _, question := range questions {
		totalPoints += question.Points
	}

	quiz.Questions = questions

	err = uc.quizRepo.UpdateQuiz(ctx, model.Quiz{ID: id, TotalPoints: totalPoints})
	if err != nil {
		return model.Quiz{}, fmt.Errorf("failed to re-count total points od quiz: %w", err)
	}
	quiz.TotalPoints = totalPoints

	return quiz, nil
}

func (uc QuizUsecase) UpdateQuiz(ctx context.Context, request model.Quiz) (model.Quiz, error) {
	err := uc.quizRepo.UpdateQuiz(ctx, request)
	if err != nil {
		return model.Quiz{}, err
	}

	return model.Quiz{
		ID: request.ID,
	}, nil
}

func (uc QuizUsecase) DeleteQuiz(ctx context.Context, id string) (model.Quiz, error) {
	err := uc.quizRepo.DeleteQuiz(ctx, id)
	if err != nil {
		return model.Quiz{}, err
	}

	return model.Quiz{ID: id}, nil
}
