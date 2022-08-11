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
var ErrInvalidPassword = errors.New("invalid password")
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

func (c *Client) CanCreate(ctx context.Context, email, password string) (bool, error) {
	var trailer metadata.MD
	response, err := c.client.CanCreate(ctx, &pb.CanCreateRequest{Email: email, Password: password}, grpc.Trailer(&trailer))

	switch c.GetErrorName(trailer) {
	case "ALREADY_EXISTS":
		return false, ErrAlreadyExists
	case "INVALID_EMAIL":
		return false, ErrInvalidEmail
	case "INVALID_PASSWORD":
		return false, ErrInvalidPassword
	}

	return response.GetOk(), err
}

func (c *Client) UpdatePassword(ctx context.Context, userId int, password string) error {
	var trailer metadata.MD
	_, err := c.client.UpdatePassword(ctx, &pb.UpdatePasswordRequest{UserId: int32(userId), Password: password}, grpc.Trailer(&trailer))

	if c.GetErrorName(trailer) == "INVALID_PASSWORD" {
		return ErrInvalidPassword
	}

	return err
}

func (c *Client) Create(ctx context.Context, email, password string) (int, error) {
	var trailer metadata.MD
	response, err := c.client.Create(ctx, &pb.CreateRequest{Email: email, Password: password}, grpc.Trailer(&trailer))

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

func (c *Client) GetByEmailAndPassword(ctx context.Context, email, password string) (*User, error) {
	in := &pb.GetByEmailAndPasswordRequest{Email: email, Password: password}
	response, err := c.client.GetByEmailAndPassword(ctx, in)

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
