package main

import (
	"log"
	"net"

	cat "github.com/wakuwaku3/example.grpc.proto"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	catService := &MyCatService{}
	cat.RegisterCatServer(server, catService)
	server.Serve(listenPort)
}
