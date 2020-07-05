package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/strizzwald/twentyone/server/models"
)

const GameCoordinators = "GameCoordinators"

/*
func AddGameCoordinator(ctx context.Context, coordinator models.GameCoordinator) error {
	defaultUuid := uuid.UUID{}
	var coordinatorJSON []byte

	if coordinator.GameId == defaultUuid {
		return errors.New("gameId is required")
	}

	key := coordinator.Id.String()

	db, err := GetDb(ctx)

	if err != nil {
		return err
	}

	if db.client.Exists(ctx, key).Val() == 1 {
		return nil
	}

	if coordinatorJSON, err = json.Marshal(coordinator); err != nil {
		return  err
	}

	db.client.HSet(ctx, GameCoordinators, key, coordinatorJSON)

	return nil
} */

func GetGameCoordinator(ctx context.Context, coordinatorId uuid.UUID) (*models.GameCoordinator, error) {
	var coordinator = new(models.GameCoordinator)

	db, err := GetDb(ctx)

	if err != nil {
		return nil, err
	}

	resp := db.client.HGet(ctx, GameCoordinators, coordinatorId.String())

	if resp.Err() != nil {
		return nil, resp.Err()
	}

	var bytes []byte

	if bytes, err = resp.Bytes(); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bytes, coordinator); err != nil {
		return nil, err
	}

	return coordinator, nil

}

func UpdateGameCoordinator(ctx context.Context, coordinatorId uuid.UUID, coordinator models.GameCoordinator) error {
	var err error
	var db *Db
	var key string = coordinatorId.String()

	if db, err = GetDb(ctx); err != nil {
		return err
	}

	if db.client.Exists(ctx, key).Val() == 0 {
		return errors.New(fmt.Sprintf("coordinator with id: %s does not exist ", key))
	}

	coordinatorJson, err := json.Marshal(coordinator)

	if err != nil {
		return err
	}

	db.client.HSet(ctx, GameCoordinators, key, coordinatorJson)

	return nil
}