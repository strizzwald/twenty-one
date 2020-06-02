package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/strizzwald/twentyone/server/models"
)

type Game struct {
}

func AddGame(ctx context.Context, game models.Game) error {
	defaultUuid := uuid.UUID{}

	if game.Id == defaultUuid  {
		return errors.New("gameId is required")
	}

	db, err := GetDb(ctx)

	if err != nil {
		return err
	}

	if db.client.Exists(ctx, game.Id.String()).Val() == 1 {
		return nil
	}

	gameJSON, err := json.Marshal(game)

	if err != nil {
		return err
	}

	status := db.client.Set(ctx, game.Id.String(), gameJSON, 0)

	if status.Err() != nil {
		return  status.Err()
	}

	return nil
}
