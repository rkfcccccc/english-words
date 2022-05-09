package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rkfcccccc/english_words/gateway/pkg/server"
)

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())

	apiGroup := router.Group("/api")

	// userGroup := apiGroup.Group("/")
	// userGroup.POST("/auth")
	// userGroup.POST("/signup")
	// userGroup.POST("/recovery")

	// authorized := router.Group("/", authorizedHandler)

	// movieGroup := authorized.Group("/movies")
	// movieGroup.POST("/:id/favorite") - make movie :id unfavorite
	// movieGroup.DELETE("/:id/favorite") - add :id favorite
	// movieGroup.GET("/") - search for movie

	// vocabularyGroup := authorized.Group("/vocabulary")
	// vocabularyGroup.GET("/event") - get an event to show
	// vocabularyGroup.PATCH("/event") - submit an event result

	movieGroup := apiGroup.Group("/movie")
	movieGroup.GET("/search")

	server := server.NewServer(router)
	server.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
