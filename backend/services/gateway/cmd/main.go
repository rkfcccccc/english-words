package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/rkfcccccc/english_words/services/gateway/internal/handler"
	"github.com/rkfcccccc/english_words/services/gateway/pkg/auth"
	"github.com/rkfcccccc/english_words/services/gateway/pkg/server"
	"github.com/rkfcccccc/english_words/shared_pkg/cache/redcache"
	"github.com/rkfcccccc/english_words/shared_pkg/redis"
	"github.com/rkfcccccc/english_words/shared_pkg/services"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}

	redis := redis.NewClient(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	cache := redcache.NewCacheRepository(redis)

	authHelper := auth.NewHelper(os.Getenv("JWT_KEY"), os.Getenv("SERVICE_KEY"), cache)

	services := services.NewService()
	handlers := handler.NewHandlers(services, authHelper)
	router := handlers.GetRouter()

	server := server.NewServer(router)
	server.Run()

	quit := make(chan os.Signal, 1)
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
