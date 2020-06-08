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
				g, err := api.NewGame(ctx, conn)

				if err != nil {
					fmt.Println(err)
				} else {
					game = &g
					fmt.Printf("Game created\nGameId: %v\nPlayerId: %v\n", game.GameId, game.PlayerId)
				}
			}
		case "list-games":
			{
				fmt.Println("Listing games.")
				fmt.Printf("\x0c")
			}
		case "join-game":
			{
				fmt.Println("Join game.")
				fmt.Printf("\x0c")
			}
		case "leave-game":
			{
				fmt.Println("Leaving game.")
				fmt.Printf("\x0c")
			}

		}
	}
}

const help = `Welcome to Twentyone
====================================================
Game commands
####################################################

- new-game: Creates new game
- join-game <game-id>: Joins an existing game

Player commands
####################################################

- hit: Blackjack hit
- stand: Blackjack stand
`
