package user

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/rkfcccccc/english_words/proto/user"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.UserServiceClient
}

func NewClient(addr string) *Client {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewUserServiceClient(conn)
	return &Client{conn, client}
}

func (c *Client) Create(ctx context.Context, email, password string) (int, error) {
	response, err := c.client.Create(ctx, &pb.CreateRequest{Email: email, Password: password})
	return int(response.GetUserId()), err
}

func (c *Client) GetById(ctx context.Context, userId int) (*User, error) {
	response, err := c.client.GetById(ctx, &pb.GetByIdRequest{UserId: int32(userId)})
	return transformFromGRPC(response), err
}

func (c *Client) GetByEmail(ctx context.Context, email string) (*User, error) {
	response, err := c.client.GetByEmail(ctx, &pb.GetByEmailRequest{Email: email})
	return transformFromGRPC(response), err
}

func (c *Client) Delete(ctx context.Context, userId int) error {
	_, err := c.client.Delete(ctx, &pb.DeleteRequest{UserId: int32(userId)})
	return err
}

func (c *Client) Close() {
	c.conn.Close()
}
