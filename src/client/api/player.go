package api

import (
	"context"
	"github.com/google/uuid"
	twentyoneService "github.com/strizzwald/twentyone/client/rpc"
	"google.golang.org/grpc"
	"log"
)

var playerClient twentyoneService.PlayerClient

func JoinGame(gameId uuid.UUID, playerId uuid.UUID, ctx context.Context, conn *grpc.ClientConn) {
	if playerClient == nil {
		playerClient = twentyoneService.NewPlayerClient(conn)
	}

	req := &twentyoneService.JoinGameRequest{
		GameId: gameId.String(),
		PlayerId: playerId.String(),
	}

	_, err := playerClient.JoinGame(ctx, req)

	if err != nil {
		log.Fatalf("Error joining game: %s", err)
	}

	log.Println("Game joined")
}