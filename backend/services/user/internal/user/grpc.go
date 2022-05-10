package user

import (
	"context"

	pb "github.com/rkfcccccc/english_words/proto/user"
	"google.golang.org/grpc"
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
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{UserId: int32(userId)}, nil
}

func (server *Server) GetById(ctx context.Context, in *pb.GetByIdRequest) (*pb.User, error) {
	u, err := server.service.GetById(ctx, int(in.UserId))
	if err != nil {
		return nil, err
	}

	return &pb.User{Id: int32(u.Id), Email: u.Email, Password: u.Password, RegisterDate: u.RegisterDate.Unix()}, nil
}

func (server *Server) GetByEmail(ctx context.Context, in *pb.GetByEmailRequest) (*pb.User, error) {
	u, err := server.service.GetByEmail(ctx, in.Email)
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
