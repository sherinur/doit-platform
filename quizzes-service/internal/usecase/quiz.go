package usecase

import (
	"context"
	"quizzes-service/internal/model"
)

type QuizUsecase struct {
	quizRepo     QuizRepo
	questionRepo QuestionRepo
	answerRepo   AnswerRepo
}

func NewQuizUsecase(qrepo QuizRepo, queRepo QuestionRepo, arepo AnswerRepo) *QuizUsecase {
	return &QuizUsecase{quizRepo: qrepo, questionRepo: queRepo, answerRepo: arepo}
}

func (uc QuizUsecase) CreateQuiz(ctx context.Context, request model.Quiz) (model.Quiz, error) {
	res, err := uc.quizRepo.CreateQuiz(ctx, request)
	if err != nil {
		return model.Quiz{}, err
	}

	return res, nil
}

func (uc QuizUsecase) GetQuizById(ctx context.Context, id string) (model.Quiz, []model.Question, []model.Answer, error) {
	quiz, err := uc.quizRepo.GetQuizById(ctx, id)
	if err != nil {
		return model.Quiz{}, nil, nil, err
	}

	questions, err := uc.questionRepo.GetQuestionsByQuizId(ctx, id)
	if err != nil {
		return model.Quiz{}, nil, nil, err
	}

	questionIds := make([]string, len(questions))
	for _, question := range questions {
		questionIds = append(questionIds, question.ID)
	}

	answers, err := uc.answerRepo.GetAnswersByQuestionIds(ctx, questionIds)

	return quiz, questions, answers, nil
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
