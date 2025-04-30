package usecase

import (
	"context"
	"errors"
	"fmt"
	"quizzes-service/internal/model"
)

type QuestionUsecase struct {
	quizRepo     QuizRepo
	questionRepo QuestionRepo
	answerRepo   AnswerRepo
}

func NewQuestionUsecase(qrepo QuizRepo, querepo QuestionRepo, arepo AnswerRepo) *QuestionUsecase {
	return &QuestionUsecase{quizRepo: qrepo, questionRepo: querepo, answerRepo: arepo}
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

	answers, err := uc.answerRepo.GetAnswersByQuestionId(ctx, id)
	if err != nil {
		return model.Question{}, fmt.Errorf("failed to get answers for the question: %w", err)
	}

	question.Answers = answers

	return question, nil
}

func (uc QuestionUsecase) GetQuestionsByQuizId(ctx context.Context, id string) ([]model.Question, error) {
	questions, err := uc.questionRepo.GetQuestionsByQuizId(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("wrong QuizID or quiz does not exist: %w", err)
	}

	questionIds := make([]string, len(questions))
	for _, question := range questions {
		questionIds = append(questionIds, question.ID)
	}

	answers, err := uc.answerRepo.GetAnswersByQuestionIds(ctx, questionIds)
	if err != nil {
		return nil, fmt.Errorf("failed to get answers for questions: %w", err)
	}

	answerMap := make(map[string][]model.Answer)
	for _, answer := range answers {
		answerMap[answer.QuestionID] = append(answerMap[answer.QuestionID], answer)
	}

	for i := range questions {
		questions[i].Answers = answerMap[questions[i].ID]
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
