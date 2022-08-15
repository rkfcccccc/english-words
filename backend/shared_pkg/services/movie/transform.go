package movie

import (
	pb "github.com/rkfcccccc/english_words/proto/movie"
)

// TODO: maybe somehow remove these shitty duplicates
func transformMovieToGRPC(movie *Movie) *pb.Movie {
	if movie == nil {
		return nil
	}

	return &pb.Movie{
		ImdbId:    movie.ImdbId,
		Title:     movie.Title,
		Year:      int32(movie.Year),
		PosterUrl: movie.PosterUrl,
	}
}

func transformMovieFromGRPC(movie *pb.Movie) *Movie {
	if movie == nil {
		return nil
	}

	return &Movie{
		ImdbId:    movie.ImdbId,
		Title:     movie.Title,
		Year:      int(movie.Year),
		PosterUrl: movie.PosterUrl,
	}
}
