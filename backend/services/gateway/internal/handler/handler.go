package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rkfcccccc/english_words/services/gateway/pkg/auth"
	"github.com/rkfcccccc/english_words/shared_pkg/services"
)

type Handlers struct {
	Services *services.Services
	Auth     *auth.Helper
}

func NewHandlers(services *services.Services, authHelper *auth.Helper) *Handlers {
	return &Handlers{services, authHelper}
}

func (h *Handlers) GetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())

	api := router.Group("/api")
	authorized := api.Group("/", h.AuthRequired)
	service := api.Group("/", h.ServiceKeyRequired)

	user := api.Group("/user")

	user.POST("/signup", h.UserSignup)
	user.POST("/login", h.UserLogin)
	user.POST("/refresh", h.UserRefresh)
	user.POST("/recovery", h.UserRecovery)

	movies := authorized.Group("/movies")
	movies.GET("/", h.MovieSearch)
	// movieGroup.GET("/:id") - get info about :id
	// movieGroup.UPDATE("/:id/favorite") - make movie :id unfavorite
	// movieGroup.DELETE("/:id/favorite") - add :id favorite
	// movieGroup.GET("/") - search for movie

	vocabulary := authorized.Group("/vocabulary")
	vocabulary.GET("/challenge", h.GetChallenge)
	vocabulary.POST("/challenge", h.FinishChallenge)

	service.POST("/movies", h.MovieCreate)

	return router
}

func (h *Handlers) newError(errName string) map[string]any {
	return map[string]any{"error_name": errName}
}
