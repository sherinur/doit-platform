package frontend

import (
	feedbacksvc "github.com/sherinur/doit-platform/apis/gen/social-service/service/frontend/feedback/v1"
	"go.uber.org/zap"
)

type Feedback struct {
	feedbacksvc.UnimplementedFeedbackServiceServer

	log *zap.Logger
	uc  FeedbackUsecase
}

func NewFile(uc FeedbackUsecase, logger *zap.Logger) *Feedback {
	return &Feedback{
		log: logger,
		uc:  uc,
	}
}
