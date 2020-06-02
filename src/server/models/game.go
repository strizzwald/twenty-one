package models

import (
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

func CreateGame(playerId uuid.UUID, joinAsDealer bool) (Game, error) {
	g := Game{Id: uuid.New(), GameState: GAME_CREATED}

	deck, err := NewDeck()

	if err != nil {
		return Game{}, err
	}

	g.Deck = deck.Shuffle()

	if joinAsDealer {
		g.Dealer = Dealer{Id: playerId}
	} else {
		g.Players = []Player{{Id:playerId}}
	}

	return g, nil
}