package verification

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	pb "github.com/rkfcccccc/english_words/proto/verification"
	"github.com/rkfcccccc/english_words/services/gateway/internal/service/base"
)

var (
	ErrTooManyRequests = errors.New("user has too many verification requests")
	ErrNotFound        = errors.New("verification request was not found")
	ErrNoAttemptsLeft  = errors.New("no attempts left")
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.VerificationServiceClient
	base.Client
}

func NewClient(addr string) *Client {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewVerificationServiceClient(conn)
	return &Client{conn: conn, client: client}
}

func (c *Client) SendCode(ctx context.Context, email string, typeId Type) (string, error) {
	var trailer metadata.MD
	response, err := c.client.SendCode(ctx, &pb.SendCodeRequest{Email: email, TypeId: int32(typeId)}, grpc.Trailer(&trailer))

	if c.GetErrorName(trailer) == "TOO_MANY_REQUESTS" {
		return "", ErrTooManyRequests
	}

	return response.GetRequestId(), err
}

func (c *Client) Verify(ctx context.Context, requestId string, code int) (bool, error) {
	var trailer metadata.MD
	response, err := c.client.Verify(ctx, &pb.VerifyRequest{RequestId: requestId, Code: int32(code)}, grpc.Trailer(&trailer))

	if c.GetErrorName(trailer) == "REQUEST_NOT_FOUND" {
		return false, ErrNotFound
	}

	if c.GetErrorName(trailer) == "NO_ATTEMPTS_LEFT" {
		return false, ErrNoAttemptsLeft
	}

	return response.GetSuccess(), err
}

func (c *Client) Close() {
	c.conn.Close()
}
