package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/strizzwald/twentyone/server/models"
	"github.com/strizzwald/twentyone/server/persistence"
	player "github.com/strizzwald/twentyone/server/rpc"
)

type PlayerService struct {
	player.UnimplementedPlayerServer
}

func (*PlayerService) JoinGame(ctx context.Context, req *player.JoinGameRequest) (*player.JoinGameResponse, error) {

	gameId, err := uuid.Parse(req.GameId)

	if err != nil {
		return nil, err
	}

	playerId, err := uuid.Parse(req.PlayerId)

	if err != nil {
		return nil, err
	}

	g, err := persistence.GetGame(ctx, gameId)

	if err != nil {
		return nil, err
	}

	if g.GameState == models.GameEnded {
		return nil, errors.New(fmt.Sprintf("game %v has ended", gameId))
	}

	if g.GameState == models.GameStarted {
		return nil,  errors.New(fmt.Sprintf("game %v has already started", gameId))
	}

	for _, p := range g.Players {
		if p.Id == playerId {
			return nil, nil
		}
	}

	// TODO: check if player already joined other currently running games.

	g.Players = append(g.Players, models.Player{Id:playerId})

	err = persistence.UpdateGame(ctx, gameId, *g)

	if err != nil {
		return nil, err
	}

	return nil, nil
}