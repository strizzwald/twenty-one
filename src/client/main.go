package main

import (
	"context"
	"fmt"
	"github.com/strizzwald/twentyone/client/api"
	"google.golang.org/grpc"
	"log"
)

const address = "localhost:7300"

var game *api.Game
var player *api.Player

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	ctx := context.Background()

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	fmt.Print(help)
	fmt.Println("Enter command...")

	for {
		var command string
		_, err := fmt.Scan(&command)

		if err != nil {
			fmt.Printf("command error: %s\n", err)
			continue
		}

		switch command {
		case "new-game":
			{
				if game != nil {
					fmt.Printf("game with id %s has already been started", game.GameId)
					break
				}
				g, err := api.NewGame(conn, ctx)

				if err != nil {
					fmt.Println(err)
				} else {
					game = &g
					player = &api.Player{Id: game.PlayerId}

					fmt.Printf("Game created\nGameId: %v\nPlayerId: %v\n", game.GameId, game.PlayerId)
				}
			}
		case "start-game":
			{
				_, err := game.StartGame(ctx, conn)

				if err != nil {
					fmt.Printf("Failed to start game\n%s", err)
				} else {
					player.AwaitTurn(*game, ctx, conn)
				}
			}

		}
	}
}

const help = `Welcome to Twentyone
====================================================
Game commands
####################################################

- new-game: Creates new game
- start-game: Starts current game

Player commands
####################################################

- 
- hit: Blackjack hit
- stand: Blackjack stand
`
