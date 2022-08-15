package movie

import (
	"context"

	pb "github.com/rkfcccccc/english_words/proto/movie"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	service *Service
	pb.UnimplementedMovieServiceServer
}

func NewServer(service *Service) *Server {
	return &Server{service: service}
}

func (server *Server) Register(s *grpc.Server) {
	pb.RegisterMovieServiceServer(s, server)
}

func (server *Server) CreateByUrl(ctx context.Context, in *pb.CreateByUrlRequest) (*pb.CreateByUrlResponse, error) {
	err := server.service.CreateByUrl(ctx, TransformMovieFromGRPC(in.Movie), in.SubtitlesUrl)

	if err != nil {
		return nil, err
	}

	return &pb.CreateByUrlResponse{}, nil
}

func (server *Server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	err := server.service.Create(ctx, TransformMovieFromGRPC(in.Movie), in.Words)

	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{}, nil
}

func (server *Server) Get(ctx context.Context, in *pb.GetRequest) (*pb.Movie, error) {
	movie, err := server.service.Get(ctx, in.ImdbId)

	if err != nil {
		return nil, err
	}

	if movie == nil {
		return nil, status.Errorf(codes.NotFound, "movie was not found")
	}

	return TransformMovieToGRPC(movie), nil
}

func (server *Server) GetWords(ctx context.Context, in *pb.GetRequest) (*pb.MovieWords, error) {
	words, err := server.service.GetWords(ctx, in.ImdbId)

	if err != nil {
		return nil, err
	}

	return &pb.MovieWords{Words: words}, nil
}

func (server *Server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := server.service.Delete(ctx, in.ImdbId)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteResponse{}, nil
}

func (server *Server) GetUserFavorites(ctx context.Context, in *pb.FavoritesRequest) (*pb.FavoritesResponse, error) {
	movies, err := server.service.GetUserFavorites(ctx, int(in.UserId))

	if err != nil {
		return nil, err
	}

	grpcMovies := make([]*pb.Movie, len(movies))
	for i, movie := range movies {
		grpcMovies[i] = TransformMovieToGRPC(&movie)
	}

	return &pb.FavoritesResponse{Movies: grpcMovies}, nil
}

func (server *Server) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	err := server.service.AddUser(ctx, in.ImdbId, int(in.UserId))

	if err != nil {
		return nil, err
	}

	return &pb.AddUserResponse{}, nil
}

func (server *Server) RemoveUser(ctx context.Context, in *pb.RemoveUserRequest) (*pb.RemoveUserResponse, error) {
	err := server.service.RemoveUser(ctx, in.ImdbId, int(in.UserId))

	if err != nil {
		return nil, err
	}

	return &pb.RemoveUserResponse{}, nil
}
