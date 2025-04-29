package usecase

import (
	"context"
	"fmt"
	"quizzes-service/internal/model"
)

type ResultUsecase struct {
	resultRepo   ResultRepo
	quizRepo     QuizRepo
	questionRepo QuestionRepo
	answerRepo   AnswerRepo
}

func NewResultUsecase(rrepo ResultRepo, qrepo QuizRepo, querepo QuestionRepo, arepo AnswerRepo) *ResultUsecase {
	return &ResultUsecase{resultRepo: rrepo, quizRepo: qrepo, questionRepo: querepo, answerRepo: arepo}
}

func (uc ResultUsecase) CreateResult(ctx context.Context, request model.Result) (model.Result, error) {
	quiz, _ := uc.quizRepo.GetQuizById(ctx, request.QuizID)

	totalPoints := 0.0
	score := 0.0

	for i := 0; i < len(request.Questions); i++ {
		question, _ := uc.questionRepo.GetQuestionById(ctx, request.Questions[i].ID)
		request.Questions[i].Points = question.Points
		totalPoints += question.Points

		for _, answer := range request.Questions[i].Answers {
			ans, _ := uc.answerRepo.GetAnswerById(ctx, answer.ID)
			if ans.IsCorrect {
				score++
			}
		}
	}

	request.Score = score / totalPoints
	fmt.Println(request.Score, score, totalPoints)

	if totalPoints != quiz.TotalPoints {
		return model.Result{}, fmt.Errorf("totalPoints does not match")
	}

	res, err := uc.resultRepo.CreateResult(ctx, request)
	if err != nil {
		return model.Result{}, err
	}

	return res, nil
}

func (uc ResultUsecase) GetResultById(ctx context.Context, id string) (model.Result, error) {
	result, err := uc.resultRepo.GetResultById(ctx, id)
	if err != nil {
		return model.Result{}, err
	}

	questions, err := uc.questionRepo.GetQuestionsByQuizId(ctx, result.QuizID)
	if err != nil {
		return model.Result{}, err
	}

	questionIds := make([]string, len(questions))
	for _, question := range questions {
		questionIds = append(questionIds, question.ID)
	}

	answers, err := uc.answerRepo.GetAnswersByQuestionIds(ctx, questionIds)
	if err != nil {
		return model.Result{}, err
	}

	answerMap := make(map[string][]model.Answer)
	for _, answer := range answers {
		answerMap[answer.QuestionID] = append(answerMap[answer.QuestionID], answer)
	}

	for i := range questions {
		questions[i].Answers = answerMap[questions[i].ID]
	}

	result.Questions = questions

	return result, nil
}

func (uc ResultUsecase) GetResultsByQuizId(ctx context.Context, id string) ([]model.Result, error) {
	results, err := uc.resultRepo.GetResultsByQuizId(ctx, id)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(results); i++ {
		questions, err := uc.questionRepo.GetQuestionsByQuizId(ctx, results[i].QuizID)
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

		results[i].Questions = questions
	}

	return results, nil
}

func (uc ResultUsecase) GetResultsByUserId(ctx context.Context, id string) ([]model.Result, error) {
	results, err := uc.resultRepo.GetResultsByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(results); i++ {
		questions, err := uc.questionRepo.GetQuestionsByQuizId(ctx, results[i].QuizID)
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

		results[i].Questions = questions
	}

	return results, nil
}

func (uc ResultUsecase) DeleteResult(ctx context.Context, id string) (model.Result, error) {
	err := uc.resultRepo.DeleteResult(ctx, id)
	if err != nil {
		return model.Result{}, err
	}

	return model.Result{ID: id}, nil
}
