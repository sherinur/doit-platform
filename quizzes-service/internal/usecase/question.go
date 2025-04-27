package usecase

import (
	"context"
	"quizzes-service/internal/model"
)

type QuestionUsecase struct {
	questionRepo QuestionRepo
	answerRepo   AnswerRepo
}

func NewQuestionUsecase(qrepo QuestionRepo, arepo AnswerRepo) *QuestionUsecase {
	return &QuestionUsecase{questionRepo: qrepo, answerRepo: arepo}
}

func (uc QuestionUsecase) CreateQuestion(ctx context.Context, request model.Question) (model.Question, error) {
	res, err := uc.questionRepo.CreateQuestion(ctx, request)
	if err != nil {
		return model.Question{}, err
	}

	return res, nil
}

func (uc QuestionUsecase) GetQuestionById(ctx context.Context, id string) (model.Question, []model.Answer, error) {
	question, err := uc.questionRepo.GetQuestionById(ctx, id)
	if err != nil {
		return model.Question{}, nil, err
	}

	answers, err := uc.answerRepo.GetAnswersByQuestionId(ctx, id)
	if err != nil {
		return model.Question{}, nil, err
	}

	return question, answers, nil
}

func (uc QuestionUsecase) GetQuestionsByQuizId(ctx context.Context, id string) ([]model.Question, []model.Answer, error) {
	questions, err := uc.questionRepo.GetQuestionsByQuizId(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	questionIds := make([]string, len(questions))
	for _, question := range questions {
		questionIds = append(questionIds, question.ID)
	}

	answers, err := uc.answerRepo.GetAnswersByQuestionIds(ctx, questionIds)

	return questions, answers, nil
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
