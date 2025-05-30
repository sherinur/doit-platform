package dto

import (
	feedbacksvc "github.com/sherinur/doit-platform/apis/gen/social-service/service/frontend/feedback/v1"
	"github.com/sherinur/doit-platform/social-service/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToModelFromAddRequest(req *feedbacksvc.AddFeedbackRequest) *model.Feedback {
	return &model.Feedback{
		UserID:   req.UserId,
		CourseID: req.CourseId,
		Comment:  req.Comment,
		Rating:   req.Rating,
	}
}

func ToModelFromUpdateRequest(req *feedbacksvc.UpdateFeedbackRequest) *model.Feedback {
	return &model.Feedback{
		Comment: req.Comment,
		Rating:  req.Rating,
	}
}

func ToProto(fb *model.Feedback) *feedbacksvc.Feedback {
	return &feedbacksvc.Feedback{
		Id:        fb.ID,
		UserId:    fb.UserID,
		CourseId:  fb.CourseID,
		Comment:   fb.Comment,
		Rating:    fb.Rating,
		CreatedAt: timestamppb.New(fb.CreatedAt),
	}
}
