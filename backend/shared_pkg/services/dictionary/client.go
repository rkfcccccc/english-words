package dictionary

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/rkfcccccc/english_words/proto/dictionary"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.DictionaryServiceClient
}

func NewClient(addr string) *Client {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewDictionaryServiceClient(conn)
	return &Client{conn, client}
}

func (c *Client) Create(ctx context.Context, word string) (string, error) {
	response, err := c.client.Create(ctx, &pb.Word{Word: word})
	return response.GetWordId(), err
}

func (c *Client) GetById(ctx context.Context, wordId string) (*WordEntry, error) {
	response, err := c.client.GetById(ctx, &pb.WordId{WordId: wordId})
	return transformFromGRPC(response), err
}

func (c *Client) GetByName(ctx context.Context, word string) (*WordEntry, error) {
	response, err := c.client.GetByName(ctx, &pb.Word{Word: word})
	return transformFromGRPC(response), err
}
