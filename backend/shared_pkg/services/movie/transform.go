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
		Id:        int32(movie.Id),
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
		Id:        int(movie.Id),
		ImdbId:    movie.ImdbId,
		Title:     movie.Title,
		Year:      int(movie.Year),
		PosterUrl: movie.PosterUrl,
	}
}

func transformSearchResultFromGRPC(movie *pb.SearchResult) *SearchResult {
	if movie == nil {
		return nil
	}

	return &SearchResult{
		Id:        int(movie.Id),
		ImdbId:    movie.ImdbId,
		Title:     movie.Title,
		Year:      int(movie.Year),
		PosterUrl: movie.PosterUrl,

		VocabularyPercent: float32(movie.VocabularyPercent),
	}
}
