package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type GameState uint32

const (
	GameCreated GameState = iota
	GameStarted GameState = iota
	GameEnded   GameState = iota
)

type Game struct {
	Id        uuid.UUID `json:"id"`
	Dealer    Dealer    `json:"dealer"`
	Players   []Player  `json:"players"`
	Deck      Deck      `json:"deck"`
	GameState GameState `json:"game_state"`
}

func CreateGame(gameId uuid.UUID, playerId uuid.UUID) (*Game, error) {
	emptyUuid := uuid.UUID{}

	if gameId == emptyUuid {
		return nil, errors.New(fmt.Sprintf("invalid gameId: %v", gameId))
	}

	if playerId == emptyUuid {
		return nil, errors.New(fmt.Sprintf("invalid playerId: %v", playerId))
	}

	g := &Game{Id: gameId, GameState: GameCreated}

	deck := NewDeck()
	deck.Shuffle()

	g.Deck = deck

	g.Dealer = Dealer{Player{Id: uuid.New()}}

	g.Players = []Player{{Id:playerId}}

	return g, nil
}

func (g *Game) StartGame() error {
	if g.GameState == GameStarted {
		return errors.New("game already started")
	}

	if g.GameState == GameEnded {
		return errors.New("game has ended")
	}

	g.GameState = GameStarted

	players := make([]*Player, len(g.Players))

	for i, _ := range players {
		players[i] = &g.Players[i]
	}

	if err := g.Dealer.Deal(&g.Deck, players); err != nil {
		return err
	}

	return nil
}