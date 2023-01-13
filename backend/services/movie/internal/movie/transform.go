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

func TransformSearchResultToGRPC(movie *SearchResult) *pb.SearchResult {
	return &pb.SearchResult{
		Id:        int32(movie.Id),
		ImdbId:    movie.ImdbId,
		Title:     movie.Title,
		Year:      int32(movie.Year),
		PosterUrl: movie.PosterUrl,

		VocabularyPercent: movie.VocabularyPercent,
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
