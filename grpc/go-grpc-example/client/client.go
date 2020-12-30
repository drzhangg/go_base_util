package main

import (
	"context"
	pb "go_base_util/grpc/go-grpc-example/proto"
	"google.golang.org/grpc"
	"log"
)

const PORT = "9001"

func main() {
	client, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer client.Close()

	c := pb.NewSearchServiceClient(client)
	resp, err := c.Search(context.TODO(), &pb.SearchRequest{Request: "hahhaha"})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	log.Printf("resp: %s", resp.GetResponse())
}
