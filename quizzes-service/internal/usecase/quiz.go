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

func (uc QuizUsecase) GetQuizById(ctx context.Context, id string) (model.Quiz, error) {
	quiz, err := uc.quizRepo.GetQuizById(ctx, id)
	if err != nil {
		return model.Quiz{}, err
	}

	questions, err := uc.questionRepo.GetQuestionsByQuizId(ctx, id)
	if err != nil {
		return model.Quiz{}, err
	}

	questionIds := make([]string, len(questions))
	for _, question := range questions {
		questionIds = append(questionIds, question.ID)
	}

	answers, err := uc.answerRepo.GetAnswersByQuestionIds(ctx, questionIds)

	answerMap := make(map[string][]model.Answer)
	for _, answer := range answers {
		answerMap[answer.QuestionID] = append(answerMap[answer.QuestionID], answer)
	}

	for i := range questions {
		questions[i].Answers = answerMap[questions[i].ID]
	}

	quiz.Questions = questions

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
