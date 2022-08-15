package dictionary

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/rkfcccccc/english_words/proto/dictionary"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	service *Service
	pb.UnimplementedDictionaryServiceServer
}

func NewServer(service *Service) *Server {
	return &Server{service: service}
}

func (server *Server) Register(s *grpc.Server) {
	pb.RegisterDictionaryServiceServer(s, server)
}

func (server *Server) Create(ctx context.Context, in *pb.Word) (*pb.WordId, error) {
	wordId, err := server.service.Create(ctx, in.Word)
	if errors.Is(err, ErrNoDefinitionsFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if err != nil {
		return nil, fmt.Errorf("service.Create: %v", err)
	}

	return &pb.WordId{WordId: wordId}, nil
}

func (server *Server) GetById(ctx context.Context, in *pb.WordId) (*pb.WordEntry, error) {
	entry, err := server.service.GetById(ctx, in.WordId)
	if err != nil {
		return nil, err
	}

	return TransformToGRPC(entry), nil
}

func (server *Server) GetByName(ctx context.Context, in *pb.Word) (*pb.WordEntry, error) {
	entry, err := server.service.GetByWord(ctx, in.Word)
	if err != nil {
		return nil, err
	}

	if entry == nil {
		return nil, status.Errorf(codes.NotFound, "word was not found")
	}

	return TransformToGRPC(entry), nil
}
