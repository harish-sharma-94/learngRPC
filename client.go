package main

import (
	"context"
	"log"

	"github.com/harish-sharma-94/helloWorldgRPC/chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9797", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not dial connection to port 9797 %v", err)
	}

	c := chat.NewChatServiceClient(conn)

	message := chat.HelloRequest{Body: "Hello server!"}
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Could not call say hello %v", err)
	}
	log.Printf("response from server: %v", response.Body)
}
