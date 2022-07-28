package movie

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/rkfcccccc/english_words/proto/movie"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.MovieServiceClient
}

func NewClient(addr string) *Client {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewMovieServiceClient(conn)
	return &Client{conn, client}
}

func (c *Client) Create(ctx context.Context, movie *Movie, words []string) error {
	_, err := c.client.Create(ctx, &pb.CreateRequest{Movie: transformMovieToGRPC(movie), Words: words})
	return err
}

func (c *Client) Get(ctx context.Context, imdbId string) (*Movie, error) {
	response, err := c.client.Get(ctx, &pb.GetRequest{ImdbId: imdbId})
	return transformMovieFromGRPC(response), err
}

func (c *Client) GetWords(ctx context.Context, imdbId string) ([]string, error) {
	response, err := c.client.GetWords(ctx, &pb.GetRequest{ImdbId: imdbId})
	return response.GetWords(), err
}

func (c *Client) Delete(ctx context.Context, imdbId string) ([]string, error) {
	response, err := c.client.GetWords(ctx, &pb.GetRequest{ImdbId: imdbId})
	return response.GetWords(), err
}

func (c *Client) GetUserFavorites(ctx context.Context, userId int) ([]*Movie, error) {
	response, err := c.client.GetUserFavorites(ctx, &pb.FavoritesRequest{UserId: int32(userId)})

	grpcMovies := response.GetMovies()
	if grpcMovies == nil {
		return nil, err
	}

	movies := make([]*Movie, len(grpcMovies))
	for i, movie := range grpcMovies {
		movies[i] = transformMovieFromGRPC(movie)
	}

	return movies, err
}

func (c *Client) AddUser(ctx context.Context, imdbId string, userId int) error {
	_, err := c.client.AddUser(ctx, &pb.AddUserRequest{UserId: int32(userId), ImdbId: imdbId})
	return err
}

func (c *Client) RemoveUser(ctx context.Context, imdbId string, userId int) error {
	_, err := c.client.RemoveUser(ctx, &pb.RemoveUserRequest{UserId: int32(userId), ImdbId: imdbId})
	return err
}
