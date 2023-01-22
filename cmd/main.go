package main

import (
	"fmt"
	"log"
	"net"
	"os"

	counterpb "github.com/MerseniBilel/streamrpc/grpc/counter"
	counter "github.com/MerseniBilel/streamrpc/routes"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("can not listen on tcp port 50051")
		os.Exit(0)
	}
	grpcServer := grpc.NewServer()
	counterpb.RegisterCounterServiceServer(grpcServer, counter.NewCounterService())

	fmt.Println("Starting grpc server at port 50051")
	// start the grpc server
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("error while serving the grpc server")
	}
}
