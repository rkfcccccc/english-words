package vocabulary

import (
	"context"

	pb "github.com/rkfcccccc/english_words/proto/vocabulary"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	service *Service
	pb.UnimplementedVocabularyServiceServer
}

func NewServer(service *Service) *Server {
	return &Server{service: service}
}

func (server *Server) Register(s *grpc.Server) {
	pb.RegisterVocabularyServiceServer(s, server)
}

func (server *Server) GetChallenge(ctx context.Context, in *pb.GetChallengeRequest) (*pb.GetChallengeResponse, error) {
	data, err := server.service.GetChallenge(ctx, int(in.UserId))
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, status.Errorf(codes.NotFound, "No available challenges")
	}

	return &pb.GetChallengeResponse{WordId: data.WordId, LearningStep: int32(data.LearningStep)}, nil
}

func (server *Server) PromoteWord(ctx context.Context, in *pb.PromoteWordRequest) (*pb.PromoteWordResponse, error) {
	err := server.service.PromoteWord(ctx, int(in.UserId), in.WordId)
	if err != nil {
		return nil, err
	}

	return &pb.PromoteWordResponse{}, nil
}

func (server *Server) ResistWord(ctx context.Context, in *pb.ResistWordRequest) (*pb.ResistWordResponse, error) {
	err := server.service.ResistWord(ctx, int(in.UserId), in.WordId)
	if err != nil {
		return nil, err
	}

	return &pb.ResistWordResponse{}, nil
}

func (server *Server) SetAlreadyLearned(ctx context.Context, in *pb.SetAlreadyLearnedRequest) (*pb.SetAlreadyLearnedResponse, error) {
	err := server.service.SetAlreadyLearned(ctx, int(in.UserId), in.WordId, in.IsAlreadyLearned)
	if err != nil {
		return nil, err
	}

	return &pb.SetAlreadyLearnedResponse{}, nil
}
