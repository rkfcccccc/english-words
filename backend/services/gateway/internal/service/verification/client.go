package verification

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/rkfcccccc/english_words/proto/verification"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.VerificationServiceClient
}

func NewClient(addr string) *Client {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewVerificationServiceClient(conn)
	return &Client{conn, client}
}

func (c *Client) SendCode(ctx context.Context, email string, typeId int) (string, error) {
	response, err := c.client.SendCode(ctx, &pb.SendCodeRequest{Email: email, TypeId: int32(typeId)})
	return response.GetRequestId(), err
}

func (c *Client) Verify(ctx context.Context, requestId string, code int) (bool, error) {
	response, err := c.client.Verify(ctx, &pb.VerifyRequest{RequestId: requestId, Code: int32(code)})
	return response.GetSuccess(), err
}

func (c *Client) Close() {
	c.conn.Close()
}
