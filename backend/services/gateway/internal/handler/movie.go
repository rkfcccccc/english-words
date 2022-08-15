package handler

import (
	"net/http"

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

	if err := h.Services.Movie.CreateByUrl(c, movie, body.SubtitlesUrl); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
