package user

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/rkfcccccc/english_words/proto/user"
	"github.com/rkfcccccc/english_words/services/gateway/internal/service/base"
)

var ErrAlreadyExists = errors.New("email already in use")
var ErrInvalidEmail = errors.New("invalid email")
var ErrTooLongPassword = errors.New("too long password")
var ErrNotFound = errors.New("user was not found")

type Client struct {
	conn   *grpc.ClientConn
	client pb.UserServiceClient
	base.Client
}

func NewClient(addr string) *Client {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewUserServiceClient(conn)
	return &Client{conn: conn, client: client}
}

func (c *Client) Create(ctx context.Context, email, password string) (int, error) {
	var trailer metadata.MD
	response, err := c.client.Create(ctx, &pb.CreateRequest{Email: email, Password: password}, grpc.Trailer(&trailer))

	switch c.GetErrorName(trailer) {
	case "ALREADY_EXISTS":
		return -1, ErrAlreadyExists
	case "INVALID_EMAIL":
		return -1, ErrInvalidEmail
	case "TOO_LONG_PASSWORD":
		return -1, ErrTooLongPassword
	}

	return int(response.GetUserId()), err
}

func (c *Client) GetById(ctx context.Context, userId int) (*User, error) {
	response, err := c.client.GetById(ctx, &pb.GetByIdRequest{UserId: int32(userId)})

	if status.Code(err) == codes.NotFound {
		return nil, ErrNotFound
	}

	return transformFromGRPC(response), err
}

func (c *Client) GetByEmail(ctx context.Context, email string) (*User, error) {
	response, err := c.client.GetByEmail(ctx, &pb.GetByEmailRequest{Email: email})

	if status.Code(err) == codes.NotFound {
		return nil, ErrNotFound
	}

	return transformFromGRPC(response), err
}

func (c *Client) Delete(ctx context.Context, userId int) error {
	_, err := c.client.Delete(ctx, &pb.DeleteRequest{UserId: int32(userId)})
	return err
}

func (c *Client) Close() {
	c.conn.Close()
}
