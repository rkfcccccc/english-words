package verification

import (
	"context"
	"errors"

	pb "github.com/rkfcccccc/english_words/proto/verification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	service *Service
	pb.UnimplementedVerificationServiceServer
}

func NewServer(service *Service) *Server {
	return &Server{service: service}
}

func (server *Server) Register(s *grpc.Server) {
	pb.RegisterVerificationServiceServer(s, server)
}

func (server *Server) SendCode(ctx context.Context, in *pb.SendCodeRequest) (*pb.SendCodeResponse, error) {
	requestId, err := server.service.SendCode(ctx, in.Email, int(in.TypeId))

	if errors.Is(err, ErrTooManyRequests) {
		return nil, status.Errorf(codes.ResourceExhausted, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	return &pb.SendCodeResponse{RequestId: requestId}, nil
}

func (server *Server) Verify(ctx context.Context, in *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	ok, err := server.service.Verify(ctx, in.RequestId, int(in.Code))

	if errors.Is(err, ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	return &pb.VerifyResponse{Success: ok}, nil
}
