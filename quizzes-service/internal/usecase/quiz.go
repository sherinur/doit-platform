package usecase

import (
	"context"
	"quizzes-service/internal/model"
)

type QuizUsecase struct {
	Repo QuizRepo
}

func NewQuizUsecase(repo QuizRepo) *QuizUsecase {
	return &QuizUsecase{Repo: repo}
}

func (uc QuizUsecase) CreateQuiz(ctx context.Context, request model.Quiz) (model.Quiz, error) {
	res, err := uc.Repo.CreateQuiz(ctx, request)
	if err != nil {
		return model.Quiz{}, err
	}

	return res, nil
}

func (uc QuizUsecase) GetQuizById(ctx context.Context, id string) (model.Quiz, error) {
	res, err := uc.Repo.GetQuizById(ctx, id)
	if err != nil {
		return model.Quiz{}, err
	}

	return res, nil
}

func (uc QuizUsecase) GetQuizAll(ctx context.Context) ([]model.Quiz, error) {
	res, err := uc.Repo.GetQuizAll(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc QuizUsecase) UpdateQuiz(ctx context.Context, request model.Quiz) (model.Quiz, error) {
	err := uc.Repo.UpdateQuiz(ctx, request)
	if err != nil {
		return model.Quiz{}, err
	}

	return model.Quiz{
		ID: request.ID,
	}, nil
}

func (uc QuizUsecase) DeleteQuiz(ctx context.Context, id string) (model.Quiz, error) {
	err := uc.Repo.DeleteQuiz(ctx, id)
	if err != nil {
		return model.Quiz{}, err
	}

	return model.Quiz{ID: id}, nil
}
