package main

import (
	"log"
	"net"

	"github.com/harish-sharma-94/helloWorldgRPC/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9797")
	if err != nil {
		log.Fatalf("Could not listen to port 9797 %v", err)
	}

	grpcServer := grpc.NewServer()
	s := chat.Server{}
	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not create a grpc server %v", err)
	}
}
