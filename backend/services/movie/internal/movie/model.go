package movie

type Movie struct {
	ImdbId    string `db:"imdb_id"`
	Title     string `db:"title"`
	Year      int    `db:"year"`
	PosterUrl string `db:"poster_url"`
}
