package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/strizzwald/twentyone/server/models"
	"github.com/strizzwald/twentyone/server/persistence"
	player "github.com/strizzwald/twentyone/server/rpc"
)

type PlayerService struct {
	player.UnimplementedPlayerServiceServer
}

func (*PlayerService) JoinGame(context.Context, *player.JoinGameRequest) (*player.JoinGameResponse, error) {
	fmt.Println("Game joined")
	return nil, nil
}

func (*PlayerService) AwaitTurn(req *player.GetTurnRequest, stream player.PlayerService_AwaitTurnServer) error {
	var playerId uuid.UUID
	var gameId uuid.UUID
	var gc *models.GameCoordinator
	var err error
	var gameEnded bool

	if playerId, err = uuid.Parse(req.PlayerId); err != nil {
		return err
	}

	if gameId, err = uuid.Parse(req.GameId); err != nil {
		return err
	}

	c := context.Background()

	if _, err = persistence.GetGame(c, gameId); err != nil {
		return err
	}

	if gc, err = models.GetGameCoordinator(); err != nil {
		return err
	}

	fmt.Printf("Game coordinator: %v", gc)

	gc.AddToQueue(playerId)

	for !gameEnded {
		select {
		case playerId := <-gc.GetPlayerTurn():
			{
				game, _ := persistence.GetGame(context.Background(), gameId)

				stream.Send(&player.GetTurnResponse{
					PlayerId:             playerId.String(),
					Game:                 fromModel(*game),
				})
			}
		case gameError := <-gc.GetGameError():
			{
				fmt.Println(gameError)
				// TODO: Send error on stream.
				// stream.Send(nil)
			}
		case gameEnded = <-gc.IsGameEnded():
			{
				// TODO: Close stream
				fmt.Println(gameEnded)
			}
		}
	}

	return nil
}

func fromModel(game models.Game) *player.Game {
	players := make([]*player.Player, len(game.Players))
	var hand []*player.Card

	for _, p := range game.Players {
		hand = make([]*player.Card, len(p.Hand))

		for _, c := range p.Hand  {
			hand = append(hand, &player.Card{
				Suit: string(c.Suit),
				Values: c.Values,
				CardType: string(c.CardType),
			})
		}

		players = append(players, &player.Player{
			Id:	  p.Id.String(),
			Hand: hand,
		})
	}

	hand = make([]*player.Card, len(game.Dealer.Hand))

	for _, c := range game.Dealer.Hand {
		hand = append(hand, &player.Card{
			Suit:                 string(c.Suit),
			Values:               c.Values,
			CardType:             string(c.CardType),
		})
	}

	return &player.Game{
		Id:      game.Id.String(),
		Dealer:  &player.Dealer{
			Hand: hand,
		},
		Players: players,
	}

}