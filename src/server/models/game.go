package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type GameState uint32

const (
	GAME_CREATED GameState = iota
	GAME_STARTED GameState = iota
	GAME_ENDED GameState = iota
)

type Game struct {
	Id        uuid.UUID `json:"id"`
	Dealer    Dealer    `json:"dealer"`
	Players   []Player  `json:"players"`
	Deck      Deck      `json:"deck"`
	GameState GameState `json:"game_state"`
}

func CreateGame(gameId uuid.UUID, playerId uuid.UUID, joinAsDealer bool) (*Game, error) {
	emptyUuid := uuid.UUID{}

	if gameId == emptyUuid {
		return nil, errors.New(fmt.Sprintf("invalid gameId: %v", gameId))
	}

	if playerId == emptyUuid {
		return nil, errors.New(fmt.Sprintf("invalid playerId: %v", playerId))
	}

	g := &Game{Id: gameId, GameState: GAME_CREATED}

	deck := NewDeck()

	g.Deck = deck.Shuffle()

	if joinAsDealer {
		g.Dealer = Dealer{}
		g.Dealer.SetId(playerId)

	} else {
		g.Players = []Player{{Id:playerId}}
	}

	return g, nil
}