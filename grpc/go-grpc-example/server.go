package main

import (
	"context"
	pb "go_base_util/grpc/go-grpc-example/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const PORT = "9001"

type SearchService struct {
}

func (c *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})

	listener, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(listener)
}
