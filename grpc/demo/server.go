package main

import (
	"go_base_util/grpc/demo/services"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()
	services.RegisterProdServiceServer(server, new(services.ProdService))

	listen, _ := net.Listen("tcp", ":8081")
	server.Serve(listen)
}
