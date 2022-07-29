package base

import (
	"google.golang.org/grpc/metadata"
)

type Client struct{}

func (client *Client) GetErrorName(md metadata.MD) string {
	errorNames := md.Get("ERROR_NAME")
	if len(errorNames) == 1 {
		return errorNames[0]
	}

	return ""
}
