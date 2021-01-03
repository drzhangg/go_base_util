package main

import (
	"go_base_util/grpc/go-grpc-example/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type StreamService struct {
}

const (
	PORT = "9002"
)

func (s *StreamService) List(r *proto.StreamRequest, stream proto.StreamService_ListServer) error {
	return nil
}

func (s *StreamService) Record(stream proto.StreamService_RecordServer) error {
	return nil
}

func (s *StreamService) Route(stream proto.StreamService_RouteServer) error {
	return nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)

}
