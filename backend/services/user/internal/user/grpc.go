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

func (server *Server) CanCreate(ctx context.Context, in *pb.CanCreateRequest) (*pb.CanCreateResponse, error) {
	ok, err := server.service.CanCreate(ctx, in.Email, in.Password)

	if errors.Is(err, ErrAlreadyExists) {
		grpc.SetTrailer(ctx, metadata.Pairs("ERROR_NAME", "ALREADY_EXISTS"))
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	if errors.Is(err, ErrInvalidEmail) {
		grpc.SetTrailer(ctx, metadata.Pairs("ERROR_NAME", "INVALID_EMAIL"))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if errors.Is(err, ErrInvalidPassword) {
		grpc.SetTrailer(ctx, metadata.Pairs("ERROR_NAME", "INVALID_PASSWORD"))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err != nil {
		return nil, fmt.Errorf("service.Create: %v", err)
	}

	return &pb.CanCreateResponse{Ok: ok}, nil
}

func (server *Server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	userId, err := server.service.Create(ctx, in.Email, in.Password)

	if err != nil {
		return nil, fmt.Errorf("service.Create: %v", err)
	}

	return &pb.CreateResponse{UserId: int32(userId)}, nil
}

func (server *Server) UpdatePassword(ctx context.Context, in *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	err := server.service.UpdatePassword(ctx, int(in.UserId), in.Password)

	if err != nil {
		return nil, fmt.Errorf("service.UpdatePassword: %v", err)
	}

	return &pb.UpdatePasswordResponse{}, nil
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

func (server *Server) GetByEmailAndPassword(ctx context.Context, in *pb.GetByEmailAndPasswordRequest) (*pb.User, error) {
	u, err := server.service.GetByEmailAndPassword(ctx, in.Email, in.Password)

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

	return &pb.DeleteResponse{}, nil
}
