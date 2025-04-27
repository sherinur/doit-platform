package usecase

import (
	"context"
	"quizzes-service/internal/model"
)

type QuestionUsecase struct {
	Repo QuestionRepo
}

func NewQuestionUsecase(repo QuestionRepo) *QuestionUsecase {
	return &QuestionUsecase{Repo: repo}
}

func (uc QuestionUsecase) CreateQuestion(ctx context.Context, request model.Question) (model.Question, error) {
	res, err := uc.Repo.CreateQuestion(ctx, request)
	if err != nil {
		return model.Question{}, err
	}

	return res, nil
}

func (uc QuestionUsecase) GetQuestionById(ctx context.Context, id string) (model.Question, error) {
	res, err := uc.Repo.GetQuestionById(ctx, id)
	if err != nil {
		return model.Question{}, err
	}

	return res, nil
}

func (uc QuestionUsecase) GetQuestionAll(ctx context.Context) ([]model.Question, error) {
	res, err := uc.Repo.GetQuestionAll(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc QuestionUsecase) UpdateQuestion(ctx context.Context, request model.Question) (model.Question, error) {
	err := uc.Repo.UpdateQuestion(ctx, request)
	if err != nil {
		return model.Question{}, err
	}

	return model.Question{
		ID: request.ID,
	}, nil
}

func (uc QuestionUsecase) DeleteQuestion(ctx context.Context, id string) (model.Question, error) {
	err := uc.Repo.DeleteQuestion(ctx, id)
	if err != nil {
		return model.Question{}, err
	}

	return model.Question{ID: id}, nil
}
