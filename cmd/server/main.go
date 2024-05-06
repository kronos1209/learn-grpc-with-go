package main

import (
	"log"
	"net"

	"github.com/kronos1209/learn-grpc-with-go/internal/server"
	"github.com/kronos1209/learn-grpc-with-go/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	if srv == nil {
		log.Fatal("server is nil")
	}

	service := server.NewGrpcServer()
	proto.RegisterUserServiceServer(srv, service)

	if err := srv.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
