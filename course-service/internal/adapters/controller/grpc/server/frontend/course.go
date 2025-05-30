package frontend

import (
	"context"

	"github.com/sherinur/doit-platform/course-service/internal/model"

	coursesvc "github.com/sherinur/doit-platform/apis/gen/course-service/service/frontend/course/v1"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	coursesvc.UnimplementedCourseServiceServer
	uc CourseUsecase
}

func NewCourse(usecase CourseUsecase) *Course {
	return &Course{uc: usecase}
}

func (h *Course) CreateCourse(ctx context.Context, req *coursesvc.CreateCourseRequest) (*coursesvc.CourseResponse, error) {
	course := &model.Course{
		Title:        req.Title,
		Description:  req.Description,
		InstructorID: req.InstructorId,
		CategoryID:   req.CategoryId,
		Tags:         req.Tags,
	}
	err := h.uc.CreateCourse(ctx, course)
	if err != nil {
		return nil, err
	}
	return &coursesvc.CourseResponse{Course: mapToProtoCourse(course)}, nil
}

func (h *Course) GetCourse(ctx context.Context, req *coursesvc.GetCourseRequest) (*coursesvc.CourseResponse, error) {
	course, err := h.uc.GetCourse(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &coursesvc.CourseResponse{Course: mapToProtoCourse(course)}, nil
}

func (h *Course) ListCourses(ctx context.Context, req *coursesvc.ListCoursesRequest) (*coursesvc.ListCoursesResponse, error) {
	courses, err := h.uc.SearchCourses(ctx, req.Search, int64(req.Page), int64(req.PageSize))
	if err != nil {
		return nil, err
	}

	var protoCourses []*coursesvc.Course
	for _, c := range courses {
		protoCourses = append(protoCourses, mapToProtoCourse(c))
	}
	return &coursesvc.ListCoursesResponse{Courses: protoCourses}, nil
}

func parseObjectID(id string) primitive.ObjectID {
	oid, _ := primitive.ObjectIDFromHex(id)
	return oid
}

func parseObjectIDs(ids []string) []primitive.ObjectID {
	var result []primitive.ObjectID
	for _, id := range ids {
		result = append(result, parseObjectID(id))
	}
	return result
}

func mapToProtoCourse(c *model.Course) *coursesvc.Course {
	return &coursesvc.Course{
		Id:           c.ID,
		Title:        c.Title,
		Description:  c.Description,
		InstructorId: c.InstructorID,
		CategoryId:   c.CategoryID,
		Tags:         c.Tags,
	}
}
