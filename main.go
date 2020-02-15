package main

import (
	"log"
	"net"
	"net/http"

	"github.com/wakuwaku3/example.grpc.go.api/cat"
	"google.golang.org/grpc"
)

func main() {
	go serveDocument()

	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	catService := &MyCatService{}
	cat.RegisterCatServer(server, catService)
	log.Println("serve gRPC! http://127.0.0.1:8080")
	log.Fatal(server.Serve(listenPort))
}

func serveDocument() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./cat/index.html") })
	log.Println("serve document! http://127.0.0.1:52300")
	log.Fatal(http.ListenAndServe(":52300", nil))
}
