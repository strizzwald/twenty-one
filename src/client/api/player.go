package api

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	twentyoneService "github.com/strizzwald/twentyone/client/rpc"
	"google.golang.org/grpc"
	"io"
)

type Player struct {
	Id uuid.UUID
}

var playerClient twentyoneService.PlayerServiceClient

func (p *Player) AwaitTurn(game Game, ctx context.Context, conn *grpc.ClientConn) error {
	client := *getPlayerCientConn(conn)

	fmt.Println("await")

	c, err := client.AwaitTurn(ctx, &twentyoneService.GetTurnRequest{
		PlayerId: p.Id.String(),
		GameId: game.GameId.String(),
	})

	fmt.Println(c)

	if err != nil {
		return err
	}

	for {

		resp, err := c.Recv()

		if err == io.EOF {
			fmt.Println("Stream closed")
		}

		if err != nil {
			return err
		}

		playerId, err := uuid.Parse(resp.PlayerId)

		if err != nil {
			return err
		}

		if playerId == p.Id {
			// TODO: Play
			fmt.Println("Your turn")
		}
	}

}

func (p Player) Hit() {}

func (p Player) Stand() {}

func getPlayerCientConn(conn *grpc.ClientConn) *twentyoneService.PlayerServiceClient {
	if playerClient == nil {
		playerClient = twentyoneService.NewPlayerServiceClient(conn)
	}

	return &playerClient
}