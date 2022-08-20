package movie

type Movie struct {
	Id        int    `json:"id"`
	ImdbId    string `json:"imdb_id"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	PosterUrl string `json:"poster_url"`
}
