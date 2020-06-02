package persistence

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Db struct {
	client *redis.Client
}

var db *Db

func GetDb(ctx context.Context) (*Db, error)  {
	if db == nil {
		db = &Db{}
		err := db.initDatabase(ctx)

		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func (db *Db) initDatabase(ctx context.Context) error {
	if db.client != nil {
		return  nil
	}

	db.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := db.client.Ping(ctx).Result()

	if err != nil {
		return nil
	}

	return nil
}


