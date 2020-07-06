package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	log.Printf("message from the client: %s", request.Body)
	return &HelloReply{Body: "Hello client"}, nil
}
