package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rkfcccccc/english_words/shared_pkg/services/movie"
)

type movieInput struct {
	ImdbId    string `json:"imdb_id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Year      int    `json:"year" binding:"required"`
	PosterUrl string `json:"poster_url" binding:"required"`
}

type movieCreateInput struct {
	Movie        *movieInput `json:"movie" binding:"required"`
	SubtitlesUrl string      `json:"subtitles_url" binding:"required"`
}

func (h *Handlers) MovieCreate(c *gin.Context) {
	var body movieCreateInput

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	movie := &movie.Movie{
		ImdbId:    body.Movie.ImdbId,
		Title:     body.Movie.Title,
		Year:      body.Movie.Year,
		PosterUrl: body.Movie.PosterUrl,
	}

	movieId, err := h.Services.Movie.CreateByUrl(c, movie, body.SubtitlesUrl)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"movie_id": movieId})
}

func (h *Handlers) MovieSearch(c *gin.Context) {
	query := c.Query("query")

	movies, err := h.Services.Movie.Search(c, query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *Handlers) MovieFavorite(c *gin.Context) {
	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt("user_id")
	if err := h.Services.Movie.AddUser(c, movieId, userId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h *Handlers) MovieUnfavorite(c *gin.Context) {
	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt("user_id")
	if err := h.Services.Movie.RemoveUser(c, movieId, userId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
