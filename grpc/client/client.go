package main

import (
	"context"
	"fmt"
	"go_base_util/grpc/client/services"
	"google.golang.org/grpc"
	"log"
)

func main() {
	client, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer client.Close()

	prodService := services.NewProdServiceClient(client)
	res, err := prodService.GetProdStock(context.TODO(), &services.ProdRequest{ProdId: 500})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.ProdStock)
}
