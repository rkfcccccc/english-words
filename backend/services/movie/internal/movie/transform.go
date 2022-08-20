package movie

import (
	pb "github.com/rkfcccccc/english_words/proto/movie"
)

func TransformMovieToGRPC(movie *Movie) *pb.Movie {
	return &pb.Movie{
		Id:        int32(movie.Id),
		ImdbId:    movie.ImdbId,
		Title:     movie.Title,
		Year:      int32(movie.Year),
		PosterUrl: movie.PosterUrl,
	}
}

func TransformMovieFromGRPC(movie *pb.Movie) *Movie {
	return &Movie{
		Id:        int(movie.Id),
		ImdbId:    movie.ImdbId,
		Title:     movie.Title,
		Year:      int(movie.Year),
		PosterUrl: movie.PosterUrl,
	}
}
