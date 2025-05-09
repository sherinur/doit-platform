package frontend

import (
	"context"

	svc "github.com/sherinur/doit-platform/apis/gen/content-service/service/frontend/file/v1"
	"github.com/sherinur/doit-platform/content-service/internal/adapter/controller/grpc/server/frontend/dto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type File struct {
	svc.UnimplementedFileServiceServer

	log *zap.Logger
	uc  FileUsecase
}

func NewFile(uc FileUsecase, logger *zap.Logger) *File {
	return &File{
		log: logger,
		uc:  uc,
	}
}

func (f *File) Create(ctx context.Context, req *svc.CreateFileRequest) (*svc.CreateFileResponse, error) {
	file, err := dto.ToFileFromCreateRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	url, err := f.uc.Create(ctx, file)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.CreateFileResponse{Url: url}, nil
}

func (f *File) Get(ctx context.Context, req *svc.GetFileRequest) (*svc.GetFileResponse, error) {
	file, err := f.uc.Get(ctx, req.Key)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.GetFileResponse{File: dto.FromFile(*file)}, nil
}

func (f *File) Delete(ctx context.Context, req *svc.DeleteFileRequest) (*svc.DeleteFileResponse, error) {
	err := f.uc.Delete(ctx, req.Key)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.DeleteFileResponse{}, nil
}
