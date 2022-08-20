package movie

type Movie struct {
	Id        int    `db:"id"`
	ImdbId    string `db:"imdb_id"`
	Title     string `db:"title"`
	Year      int    `db:"year"`
	PosterUrl string `db:"poster_url"`
}
