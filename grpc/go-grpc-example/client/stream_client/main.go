package main

import (
	"go_base_util/grpc/go-grpc-example/proto"
	"google.golang.org/grpc"
	"log"
)

const (
	PORT = "9002"
)

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err:%v", err)
	}
	defer conn.Close()

	client := proto.NewStreamServiceClient(conn)

	err = printLists(client, &proto.StreamRequest{Pt: &proto.StreamPoint{
		Name:  "grpc stream list client:list",
		Value: 2021,
	}})
	if err != nil {
		log.Fatalf("printLists.err: %v", err)
	}

	err = printRecord(client,&proto.StreamRequest{Pt: &proto.StreamPoint{
		Name:  "grpc stream record client: record",
		Value: 2021,
	}})
	if err != nil {
		log.Fatalf("printRecord.err: %v", err)
	}

	err = printRoute(client,&proto.StreamRequest{Pt: &proto.StreamPoint{
		Name:  "grpc stream route client: route",
		Value: 2021,
	}})
	if err != nil {
		log.Fatalf("printRoute.err: %v", err)
	}
}

func printLists(client proto.StreamServiceClient, r *proto.StreamRequest) error {
	return nil
}

func printRecord(client proto.StreamServiceClient, r *proto.StreamRequest) error {
	return nil
}

func printRoute(client proto.StreamServiceClient, r *proto.StreamRequest) error {
	return nil
}
