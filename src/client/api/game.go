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

func NewGame(ctx context.Context, conn *grpc.ClientConn) (Game, error) {
	game := &Game{
		GameId:   uuid.New(),
		PlayerId: uuid.New(),
	}

	g := &twentyoneService.CreateGameRequest{
		PlayerId:	game.PlayerId.String(),
		GameId:     game.GameId.String(),
	}

	client := *getGameClientConn(conn)
	_, err := client.NewGame(ctx, g)

	if err != nil {
		return Game{}, errors.New(fmt.Sprintf("Failed to create game: %s", err))
	}

	return *game, nil
}

func (g *Game) StartGame(ctx context.Context, conn *grpc.ClientConn) (bool, error) {

	client := *getGameClientConn(conn)
	res, err := client.StartGame(ctx, &twentyoneService.StartGameRequest{GameId: g.GameId.String()})

	if err != nil {
		return false, err
	}

	return res.GameStarted, nil
}

func getGameClientConn(conn *grpc.ClientConn) *twentyoneService.GameClient {
	if gameClient == nil {
		gameClient = twentyoneService.NewGameClient(conn)
	}

	return &gameClient
}