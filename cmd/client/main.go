package main

import (
	"context"
	"log"
	"time"

	"github.com/kronos1209/learn-grpc-with-go/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	res, err := client.CreateUser(ctx, &proto.CreateUserRequest{
		UserId:   "1234",
		Name:     "Test",
		Password: "password",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("", res.AccountId)
}
