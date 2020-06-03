package main

import (
	"context"
	"github.com/google/uuid"
	twentyoneservice "github.com/strizzwald/twentyone/client/rpc"
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
	gameClient := twentyoneservice.NewGameClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	g := &twentyoneservice.CreateGameRequest{
		PlayerId:     uuid.New().String(),
		GameId:       uuid.New().String(),
		JoinAsDealer: true,
	}

	r, err := gameClient.NewGame(ctx, g)

	if err != nil {
		log.Fatalf("could not create game: %v", err)
	}

	log.Printf("Deck: %v", r)

	playerClient := twentyoneservice.NewPlayerClient(conn)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)

	jr := &twentyoneservice.JoinGameRequest{GameId: r.GameId, PlayerId: uuid.New().String()}

	_, err = playerClient.JoinGame(ctx, jr)

	if err != nil {
		log.Fatalf("could not join game: %v", r.GameId)
	}

}
