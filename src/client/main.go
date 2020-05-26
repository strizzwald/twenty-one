package main

import (
	"context"
	deckservice "github.com/strizzwald/twentyone/client/rpc"
	"google.golang.org/grpc"
	"log"
	"time"
)

const address = "localhost:7300"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := deckservice.NewDeckServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.NewDeck(ctx, new(deckservice.NewDeckRequest))

	if err != nil {
		log.Fatalf("could not retrieve deck: %v", err)
	}

	log.Printf("Deck: %v", r.Deck)
}
