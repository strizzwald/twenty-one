package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/strizzwald/twentyone/server/models"
	"github.com/strizzwald/twentyone/server/persistence"
	game "github.com/strizzwald/twentyone/server/rpc"
)

type GameService struct {
	game.UnimplementedGameServer
}

func (*GameService) NewGame(ctx context.Context, req *game.CreateGameRequest) (*game.CreateGameResponse, error) {
	var playerId uuid.UUID
	var gameId uuid.UUID

	var err error

	if req.PlayerId == "" {
		playerId = uuid.New()
	} else {
		playerId, err = uuid.Parse(req.PlayerId)

		if err != nil {
			fmt.Println(err)
			return new(game.CreateGameResponse), err
		}
	}

	if req.GameId == "" {
		gameId = uuid.New()
	} else {
		gameId, err = uuid.Parse(req.GameId)

		if err != nil {
			return new(game.CreateGameResponse), err
		}
	}

	g, err := models.CreateGame(gameId, playerId)

	if err != nil {
		return new(game.CreateGameResponse), err
	}

	err = persistence.AddGame(ctx, *g)

	if err != nil {
		return new(game.CreateGameResponse), err
	}

	return &game.CreateGameResponse{GameId: g.Id.String()}, nil
}

func (*GameService) StartGame(ctx context.Context, req *game.StartGameRequest) (*game.StartGameResponse, error) {

	gameId, err := uuid.Parse(req.GameId)

	if err != nil {
		return &game.StartGameResponse{GameStarted:false}, err
	}

	g, err := persistence.GetGame(ctx, gameId)

	if err != nil {
		return &game.StartGameResponse{GameStarted:false}, err
	}

	if err = g.StartGame(); err != nil {
		return &game.StartGameResponse{GameStarted:false}, err
	}

	if len(g.Players) == 0 {
		return &game.StartGameResponse{GameStarted:false}, errors.New("waiting for more players to join")
	}

	if err := persistence.UpdateGame(ctx, g.Id, *g); err != nil {
		return  &game.StartGameResponse{GameStarted:false}, err
	}

	return &game.StartGameResponse{GameStarted:true}, nil
}