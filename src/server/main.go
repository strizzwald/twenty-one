package main

import (
	"context"
	"github.com/strizzwald/twentyone/server/models"
	deckservice "github.com/strizzwald/twentyone/server/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":7300"
)

type server struct {
	deckservice.UnimplementedDeckServiceServer
}

func (*server) NewDeck(_ context.Context, _ *deckservice.NewDeckRequest) (*deckservice.DeckResponse, error) {
	log.Println("New deck request received")

	newDeck, err := models.NewDeck()

	if err != nil {
		return nil, nil
	}

	var c []*deckservice.Card
	for _, card := range newDeck.Cards {
		c = append(c, &deckservice.Card{
			Suit:     string(card.Suit),
			Values:   card.Values,
			CardType: string(card.CardType),
		})
	}

	return &deckservice.DeckResponse{Deck: &deckservice.Deck{C: c}}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	deckservice.RegisterDeckServiceServer(s, new(server))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
