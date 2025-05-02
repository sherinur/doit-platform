package usecase

import (
	"context"
	"fmt"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
)

type ResultUsecase struct {
	resultRepo   ResultRepo
	quizRepo     QuizRepo
	questionRepo QuestionRepo
}

func NewResultUsecase(rrepo ResultRepo, qrepo QuizRepo, querepo QuestionRepo) *ResultUsecase {
	return &ResultUsecase{resultRepo: rrepo, quizRepo: qrepo, questionRepo: querepo}
}

func (uc ResultUsecase) CreateResult(ctx context.Context, request model.Result) (model.Result, error) {
	if request.UserID == "" || request.Status == "" || request.Questions == nil {
		return model.Result{}, fmt.Errorf("invalid request")
	}

	quiz, err := uc.quizRepo.GetQuizById(ctx, request.QuizID)
	if err != nil {
		return model.Result{}, fmt.Errorf("wrong QuizID format or Quiz with this ID does no exist")
	}

	totalPoints := 0.0
	score := 0.0

	for i := 0; i < len(request.Questions); i++ {
		question, _ := uc.questionRepo.GetQuestionById(ctx, request.Questions[i].ID)
		request.Questions[i].Points = question.Points
		totalPoints += question.Points

		for _, answer := range request.Questions[i].Answers {
			if answer.IsCorrect {
				score++
			}
		}
	}

	request.Score = score / totalPoints

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

	for i := range result.Questions {
		question, _ := uc.questionRepo.GetQuestionById(ctx, result.Questions[i].ID)
		result.Questions[i].Text = question.Text
		result.Questions[i].Type = question.Type
		result.Questions[i].Points = question.Points
		result.Questions[i].QuizID = question.QuizID

		for j := range result.Questions[i].Answers {
			for _, answer := range question.Answers {
				if result.Questions[i].Answers[j].ID == answer.ID {
					result.Questions[i].Answers[j].Text = answer.Text
				}
			}
		}
	}

	return result, nil
}

func (uc ResultUsecase) GetResultsByQuizId(ctx context.Context, id string) ([]model.Result, error) {
	results, err := uc.resultRepo.GetResultsByQuizId(ctx, id)
	if err != nil {
		return nil, err
	}

	for k := 0; k < len(results); k++ {
		for i := range results[k].Questions {
			question, _ := uc.questionRepo.GetQuestionById(ctx, results[k].Questions[i].ID)
			results[k].Questions[i].Text = question.Text
			results[k].Questions[i].Type = question.Type
			results[k].Questions[i].Points = question.Points
			results[k].Questions[i].QuizID = question.QuizID

			for j := range results[k].Questions[i].Answers {
				for _, answer := range question.Answers {
					if results[k].Questions[i].Answers[j].ID == answer.ID {
						results[k].Questions[i].Answers[j].Text = answer.Text
					}
				}
			}
		}
	}

	return results, nil
}

func (uc ResultUsecase) GetResultsByUserId(ctx context.Context, id string) ([]model.Result, error) {
	results, err := uc.resultRepo.GetResultsByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	for k := 0; k < len(results); k++ {
		for i := range results[k].Questions {
			question, _ := uc.questionRepo.GetQuestionById(ctx, results[k].Questions[i].ID)
			results[k].Questions[i].Text = question.Text
			results[k].Questions[i].Type = question.Type
			results[k].Questions[i].Points = question.Points
			results[k].Questions[i].QuizID = question.QuizID

			for j := range results[k].Questions[i].Answers {
				for _, answer := range question.Answers {
					if results[k].Questions[i].Answers[j].ID == answer.ID {
						results[k].Questions[i].Answers[j].Text = answer.Text
					}
				}
			}
		}
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
