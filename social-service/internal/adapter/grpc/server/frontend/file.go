package frontend

import (
	feedbacksvc "github.com/sherinur/doit-platform/apis/gen/social-service/service/frontend/feedback/v1/feedbacksvc"
	"go.uber.org/zap"
)

type Feedback struct {
	feedbacksvc.UnimplementedFileServiceServer

	log *zap.Logger
	uc  FeedbackUsecase
}

func NewFile(uc FeedbackUsecase, logger *zap.Logger) *Feedback {
	return &Feedback{
		log: logger,
		uc:  uc,
	}
}
