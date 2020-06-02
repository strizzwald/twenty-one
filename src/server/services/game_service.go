package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/strizzwald/twentyone/server/models"
	"github.com/strizzwald/twentyone/server/persistence"
	"github.com/strizzwald/twentyone/server/rpc"
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

	g, err := models.CreateGame(gameId, playerId, req.JoinAsDealer)

	if err != nil {
		return new(game.CreateGameResponse), err
	}

	err = persistence.AddGame(ctx, *g)

	if err != nil {
		return new(game.CreateGameResponse), err
	}

	return &game.CreateGameResponse{GameId: g.Id.String()}, nil
}

func StartGame(gameId uuid.UUID) error {
	// TODO: Check that game exists
	// TODO: Check that game has not ended
	return errors.New("not implemented")
}

func JoinGame(playerId uuid.UUID, gameId uuid.UUID) error {
	// TODO: Check that game exists
	// TODO: Check that game is still running
	// TODO: Check that player is not in another running game

	return errors.New("not implemented")
}