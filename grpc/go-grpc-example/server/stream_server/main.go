package main

import (
	"go_base_util/grpc/go-grpc-example/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type StreamService struct {
}

const (
	PORT = "9002"
)

func main() {
	server := grpc.NewServer()
	proto.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)

}

func (s *StreamService) List(r *proto.StreamRequest, stream proto.StreamService_ListServer) error {
	for i := 0; i <= 6; i++ {
		err := stream.Send(&proto.StreamResponse{Pt: &proto.StreamPoint{
			Name:  r.Pt.Name,
			Value: r.Pt.Value + int32(i),
		}})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StreamService) Record(stream proto.StreamService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.StreamResponse{Pt: &proto.StreamPoint{
				Name:  "grpc stream server: Record",
				Value: 1,
			}})
		}

		if err != nil {
			return err
		}
		log.Printf("stream.Recv pt.name:%v, pt.value:%v", r.Pt.Name, r.Pt.Value)
	}
	return nil
}

func (s *StreamService) Route(stream proto.StreamService_RouteServer) error {
	return nil
}
