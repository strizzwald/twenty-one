package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/strizzwald/twentyone/server/models"
)

func AddGame(ctx context.Context, game models.Game) error {
	defaultUuid := uuid.UUID{}

	if game.Id == defaultUuid  {
		fmt.Println("Game guid error")
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

func GetGame(ctx context.Context, gameId uuid.UUID) (*models.Game, error) {
	var game models.Game
	db, err := GetDb(ctx)

	if err != nil {
		return nil, err
	}

	response := db.client.Get(ctx, gameId.String())

	if response.Err() != nil {
		return nil, response.Err()
	}

	 err = json.Unmarshal([]byte(response.Val()), game)

	 if err != nil {
	 	return nil, err
	 }

	 return &game, nil
}

func UpdateGame(ctx context.Context, gameId uuid.UUID, game models.Game) error {

	db, err := GetDb(ctx)

	if err != nil {
		return  err
	}

	if db.client.Exists(ctx, gameId.String()).Val() == 0 {
		return errors.New(fmt.Sprintf("game with id: %v does not exist", gameId))
	}

	response := db.client.Set(ctx, game.Id.String(), game, 0)

	if response.Err() != nil {
		return response.Err()
	}

	return nil
}
