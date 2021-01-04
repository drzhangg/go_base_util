package main

import (
	"context"
	"go_base_util/grpc/go-grpc-example/proto"
	"google.golang.org/grpc"
	"io"
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

	err = printRecord(client, &proto.StreamRequest{Pt: &proto.StreamPoint{
		Name:  "grpc stream record client: record",
		Value: 2021,
	}})
	if err != nil {
		log.Fatalf("printRecord.err: %v", err)
	}

	err = printRoute(client, &proto.StreamRequest{Pt: &proto.StreamPoint{
		Name:  "grpc stream route client: route",
		Value: 2021,
	}})
	if err != nil {
		log.Fatalf("printRoute.err: %v", err)
	}
}

func printLists(client proto.StreamServiceClient, r *proto.StreamRequest) error {
	stream, err := client.List(context.TODO(), r)
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp: pt.name: %s, pt.value: %v", resp.Pt.Name, resp.Pt.Value)
	}
	return nil
}

func printRecord(client proto.StreamServiceClient, r *proto.StreamRequest) error {
	stream, err := client.Record(context.TODO())
	if err != nil {
		return err
	}

	for n := 0; n < 6; n++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("resp: pt.name:%v, pt.value : %v", resp.Pt.Name, resp.Pt.Value)
	return nil
}

func printRoute(client proto.StreamServiceClient, r *proto.StreamRequest) error {
	stream, err := client.Route(context.TODO())
	if err != nil {
		return err
	}

	for n := 0; n <= 7; n++ {
		err = stream.Send(r)
		if err != nil {
			return err
		}

		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp: pt.name: %v,pt.value: %v", resp.Pt.Name, resp.Pt.Value)
	}

	stream.CloseSend()
	return nil
}
