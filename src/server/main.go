package main

import (
	"fmt"
	twentyoneservice "github.com/strizzwald/twentyone/server/rpc"
	"github.com/strizzwald/twentyone/server/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":7300"
)

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	twentyoneservice.RegisterGameServer(s, new(services.GameService))

	fmt.Printf("Server running on port: %s\n", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
