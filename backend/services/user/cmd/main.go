package main

import (
	"log"
	"net"

	pb "github.com/rkfcccccc/english_words/proto/users"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	log.Printf("listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
