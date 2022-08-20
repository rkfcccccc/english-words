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

	apiGroup := router.Group("/api")
	authorized := apiGroup.Group("/", h.AuthRequired)

	userGroup := apiGroup.Group("/user")

	userGroup.POST("/signup", h.UserSignup)
	userGroup.POST("/login", h.UserLogin)
	userGroup.POST("/refresh", h.UserRefresh)
	userGroup.POST("/recovery", h.UserRecovery)

	movieGroup := apiGroup.Group("/movies")
	movieGroup.POST("/", h.MovieCreate)
	movieGroup.GET("/", h.MovieSearch)
	// movieGroup.GET("/:id") - get info about :id
	// movieGroup.UPDATE("/:id/favorite") - make movie :id unfavorite
	// movieGroup.DELETE("/:id/favorite") - add :id favorite
	// movieGroup.GET("/") - search for movie

	vocabularyGroup := authorized.Group("/vocabulary")
	vocabularyGroup.GET("/challenge", h.GetChallenge)
	// vocabularyGroup.PATCH("/challenge") - submit the challenge result

	return router
}

func (h *Handlers) newError(errName string) map[string]any {
	return map[string]any{"error_name": errName}
}
