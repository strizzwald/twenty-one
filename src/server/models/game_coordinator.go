package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

var gameCoordinator *GameCoordinator
var coordinating bool

type GameCoordinator struct {
	Id                 uuid.UUID
	PlayerQueue        map[int]uuid.UUID
	game               *Game
	currentPlayerIndex int
	playerTurn         chan uuid.UUID
	played 			   chan uuid.UUID
	gameError          chan error
	gameEnded          chan bool
}

func GetGameCoordinator() (*GameCoordinator, error) {

	if gameCoordinator == nil {
		return nil, errors.New("Game coordinator is not set")
	}

	return gameCoordinator, nil
}

func (gc *GameCoordinator) GetPlayerTurn() chan uuid.UUID {
	return gc.playerTurn
}

func (gc *GameCoordinator) GetGameError() chan error {
	return gc.gameError
}

func (gc *GameCoordinator) IsGameEnded() chan bool {
	return gc.gameEnded
}

func NewGameCoordinator(game *Game) GameCoordinator {
	gameCoordinator = &GameCoordinator{
		Id:          uuid.New(),
		PlayerQueue: make(map[int]uuid.UUID, len(game.Players)+1),
		game:        game,
		playerTurn:  make(chan uuid.UUID),
		played:      make(chan uuid.UUID),
		gameError:   make(chan error),
	}

	return *gameCoordinator
}

func (gc *GameCoordinator) AddToQueue(playerId uuid.UUID) {
	for _, p := range gc.PlayerQueue {
		if p == playerId {
			return
		}
	}

	playerQueue := gameCoordinator.PlayerQueue
	playerQueue[len(playerQueue) + 1] = playerId

	fmt.Println("Added to queue")

	if !coordinating && len(playerQueue) == len(gc.game.Players) {
		gameCoordinator.PlayerQueue[len(playerQueue) + 1] = gc.game.Dealer.Id

		coordinating = true

		fmt.Println("Coordinator started")
		go gc.Coordinate()
	}
}

func (gc *GameCoordinator) Coordinate() {
	for gc.game.GameState != GameEnded {

		gc.playerTurn <- gc.PlayerQueue[gc.currentPlayerIndex]
		timeout := time.NewTimer(10 * time.Second)

		select {
		case <- timeout.C:
			{
				gc.gameError <- errors.New(fmt.Sprintf("Player with id: %v did not play in time.", gc.PlayerQueue[gc.currentPlayerIndex]))
			}
		case p := <- gc.played:
			{
				if p == gc.PlayerQueue[gc.currentPlayerIndex] {
					gc.currentPlayerIndex = (gc.currentPlayerIndex + 1) % (len(gc.PlayerQueue) + 1)
				} else {
					gc.gameError <- errors.New(fmt.Sprintf("player with id: %v played out of turn", p))
				}
			}
		}
	}
}