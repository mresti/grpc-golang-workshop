package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pbBooks "github.com/mresti/grpc-golang-workshop/start/books"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	fmt.Println("Server running at http://0.0.0.0:50051")
	service := service{}
	pbBooks.RegisterBookServiceServer(grpcServer, &service)
	grpcServer.Serve(lis)
}

type service struct {
}

func (s *service) List(context.Context, *pbBooks.Empty) (*pbBooks.Empty, error) {
	return &pbBooks.Empty{}, status.Error(codes.Unimplemented, "The server does not implement this method")
}
