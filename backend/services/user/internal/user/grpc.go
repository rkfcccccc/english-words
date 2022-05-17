package user

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/rkfcccccc/english_words/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	service *Service
	pb.UnimplementedUserServiceServer
}

func NewServer(service *Service) *Server {
	return &Server{service: service}
}

func (server *Server) Register(s *grpc.Server) {
	pb.RegisterUserServiceServer(s, server)
}

func (server *Server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	userId, err := server.service.Create(ctx, in.Email, in.Password)

	if errors.Is(err, ErrAlreadyExists) {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	if errors.Is(err, ErrInvalidEmail) {
		grpc.SetTrailer(ctx, metadata.Pairs("ERROR_NAME", "INVALID_EMAIL"))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if errors.Is(err, ErrTooLongPassword) {
		grpc.SetTrailer(ctx, metadata.Pairs("ERROR_NAME", "TOO_LONG_PASSWORD"))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if errors.Is(err, ErrTooLongEmail) {
		grpc.SetTrailer(ctx, metadata.Pairs("ERROR_NAME", "TOO_LONG_EMAIL"))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err != nil {
		return nil, fmt.Errorf("service.Create: %v", err)
	}

	return &pb.CreateResponse{UserId: int32(userId)}, nil
}

func (server *Server) GetById(ctx context.Context, in *pb.GetByIdRequest) (*pb.User, error) {
	u, err := server.service.GetById(ctx, int(in.UserId))

	if errors.Is(err, ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if err != nil {
		return nil, err
	}

	return &pb.User{Id: int32(u.Id), Email: u.Email, Password: u.Password, RegisterDate: u.RegisterDate.Unix()}, nil
}

func (server *Server) GetByEmail(ctx context.Context, in *pb.GetByEmailRequest) (*pb.User, error) {
	u, err := server.service.GetByEmail(ctx, in.Email)

	if errors.Is(err, ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if err != nil {
		return nil, err
	}

	return &pb.User{Id: int32(u.Id), Email: u.Email, Password: u.Password, RegisterDate: u.RegisterDate.Unix()}, nil
}

func (server *Server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := server.service.Delete(ctx, int(in.UserId))

	if err != nil {
		return nil, err
	}

	return &pb.DeleteResponse{Success: true}, nil
}
