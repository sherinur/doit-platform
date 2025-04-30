package usecase

import (
	"context"
	"errors"
	"quizzes-service/internal/model"
)

type AnswerUsecase struct {
	Repo AnswerRepo
}

func NewAnswerUsecase(repo AnswerRepo) *AnswerUsecase {
	return &AnswerUsecase{Repo: repo}
}

func (uc AnswerUsecase) CreateAnswer(ctx context.Context, request model.Answer) (model.Answer, error) {
	if request.Text == "" || request.QuestionID == "" {
		return model.Answer{}, errors.New("invalid input data")
	}

	res, err := uc.Repo.CreateAnswer(ctx, request)
	if err != nil {
		return model.Answer{}, err
	}

	return res, nil
}

func (uc AnswerUsecase) CreateAnswers(ctx context.Context, request []model.Answer) ([]model.Answer, error) {
	for _, answer := range request {
		if answer.Text == "" || answer.QuestionID == "" {
			return nil, errors.New("invalid input data")
		}
	}

	res, err := uc.Repo.CreateAnswers(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc AnswerUsecase) GetAnswerById(ctx context.Context, id string) (model.Answer, error) {
	res, err := uc.Repo.GetAnswerById(ctx, id)
	if err != nil {
		return model.Answer{}, err
	}

	return res, nil
}

func (uc AnswerUsecase) GetAnswersByQuestionId(ctx context.Context, id string) ([]model.Answer, error) {
	res, err := uc.Repo.GetAnswersByQuestionId(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc AnswerUsecase) UpdateAnswer(ctx context.Context, request model.Answer) (model.Answer, error) {
	err := uc.Repo.UpdateAnswer(ctx, request)
	if err != nil {
		return model.Answer{}, err
	}

	return model.Answer{
		ID: request.ID,
	}, nil
}

func (uc AnswerUsecase) DeleteAnswer(ctx context.Context, id string) (model.Answer, error) {
	err := uc.Repo.DeleteAnswer(ctx, id)
	if err != nil {
		return model.Answer{}, err
	}

	return model.Answer{ID: id}, nil
}
