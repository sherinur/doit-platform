package frontend

import (
	"context"
	"course-service/internal/model"

	coursesvc "github.com/sherinur/doit-platform/apis/gen/course-service/service/frontend/course/v1"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CourseHandler struct {
	coursesvc.UnimplementedCourseServiceServer
	uc CourseUsecase
}

func NewCourseHandler(usecase CourseUsecase) *CourseHandler {
	return &CourseHandler{uc: usecase}
}

func (h *CourseHandler) CreateCourse(ctx context.Context, req *proto.CreateCourseRequest) (*proto.CourseResponse, error) {
	course := &model.Course{
		Title:        req.Title,
		Description:  req.Description,
		InstructorID: parseObjectID(req.InstructorId),
		CategoryID:   parseObjectID(req.CategoryId),
		Tags:         parseObjectIDs(req.Tags),
	}
	err := h.uc.CreateCourse(ctx, course)
	if err != nil {
		return nil, err
	}
	return &proto.CourseResponse{Course: mapToProtoCourse(course)}, nil
}

func (h *CourseHandler) GetCourse(ctx context.Context, req *proto.GetCourseRequest) (*proto.CourseResponse, error) {
	id := parseObjectID(req.Id)
	course, err := h.uc.GetCourse(ctx, id)
	if err != nil {
		return nil, err
	}
	return &proto.CourseResponse{Course: mapToProtoCourse(course)}, nil
}

func (h *CourseHandler) ListCourses(ctx context.Context, req *proto.ListCoursesRequest) (*proto.ListCoursesResponse, error) {
	courses, err := h.uc.SearchCourses(ctx, req.Search, int64(req.Page), int64(req.PageSize))
	if err != nil {
		return nil, err
	}

	var protoCourses []*proto.Course
	for _, c := range courses {
		protoCourses = append(protoCourses, mapToProtoCourse(c))
	}
	return &proto.ListCoursesResponse{Courses: protoCourses}, nil
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

func mapToProtoCourse(c *model.Course) *proto.Course {
	return &proto.Course{
		Id:           c.ID.Hex(),
		Title:        c.Title,
		Description:  c.Description,
		InstructorId: c.InstructorID.Hex(),
		CategoryId:   c.CategoryID.Hex(),
		Tags:         toHexList(c.Tags),
	}
}

func toHexList(objs []primitive.ObjectID) []string {
	var list []string
	for _, obj := range objs {
		list = append(list, obj.Hex())
	}
	return list
}
