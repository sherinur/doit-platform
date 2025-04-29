package usecase

import (
	"context"
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
	res, err := uc.questionRepo.CreateQuestion(ctx, request)
	if err != nil {
		return model.Question{}, err
	}

	err = uc.quizRepo.ChangeTotalPointsQuiz(ctx, request.QuizID, request.Points)
	if err != nil {
		return model.Question{}, fmt.Errorf("error updating quiz total points quiz: %w", err)
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
		return model.Question{}, err
	}

	question.Answers = answers

	return question, nil
}

func (uc QuestionUsecase) GetQuestionsByQuizId(ctx context.Context, id string) ([]model.Question, error) {
	questions, err := uc.questionRepo.GetQuestionsByQuizId(ctx, id)
	if err != nil {
		return nil, err
	}

	questionIds := make([]string, len(questions))
	for _, question := range questions {
		questionIds = append(questionIds, question.ID)
	}

	answers, err := uc.answerRepo.GetAnswersByQuestionIds(ctx, questionIds)
	if err != nil {
		return nil, err
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
	oldQuestion, err := uc.questionRepo.GetQuestionById(ctx, request.ID)
	if err != nil {
		return model.Question{}, err
	}

	err = uc.questionRepo.UpdateQuestion(ctx, request)
	if err != nil {
		return model.Question{}, err
	}

	if oldQuestion.QuizID != request.QuizID {
		err = uc.quizRepo.ChangeTotalPointsQuiz(ctx, oldQuestion.QuizID, 0-request.Points)
		if err != nil {
			return model.Question{}, fmt.Errorf("error updating quiz total points quiz: %w", err)
		}

		err = uc.quizRepo.ChangeTotalPointsQuiz(ctx, request.QuizID, request.Points)
		if err != nil {
			return model.Question{}, fmt.Errorf("error updating quiz total points quiz: %w", err)
		}
	}

	return model.Question{
		ID: request.ID,
	}, nil
}

func (uc QuestionUsecase) DeleteQuestion(ctx context.Context, id string) (model.Question, error) {
	question, err := uc.questionRepo.GetQuestionById(ctx, id)
	if err != nil {
		return model.Question{}, err
	}

	err = uc.questionRepo.DeleteQuestion(ctx, id)
	if err != nil {
		return model.Question{}, err
	}

	err = uc.quizRepo.ChangeTotalPointsQuiz(ctx, question.QuizID, 0-question.Points)
	if err != nil {
		return model.Question{}, fmt.Errorf("error updating quiz total points quiz: %w", err)
	}

	return model.Question{ID: id}, nil
}
