package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	twentyoneService "github.com/strizzwald/twentyone/client/rpc"
	"google.golang.org/grpc"
)

type Game struct {
	GameId uuid.UUID
	PlayerId uuid.UUID
}

var gameClient twentyoneService.GameClient
var game *Game

func NewGame(ctx context.Context, conn *grpc.ClientConn) (Game, error) {

	if gameClient == nil {
		gameClient = twentyoneService.NewGameClient(conn)
	}

	if game != nil {
		return Game{}, errors.New(fmt.Sprintf("Cannot create multiple games. Game with id: %s has been previously created.", game.GameId))
	} else {
		game = &Game{
			GameId:   uuid.New(),
			PlayerId: uuid.New(),
		}
	}

	g := &twentyoneService.CreateGameRequest{
		PlayerId:	game.PlayerId.String(),
		GameId:     game.GameId.String(),
	}

	_, err := gameClient.NewGame(ctx, g)

	if err != nil {
		return Game{}, errors.New(fmt.Sprintf("Failed to create game: %s", err))
	}

	return *game, nil
}