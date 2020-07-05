package models

import "errors"

type Dealer struct {
	Player `json:"player"`
}

func (d *Dealer) Deal(deck *Deck, players []*Player) (err error) {
	if (len(players)*2)+1 > len(deck.Cards) {
		return errors.New("not enough cards in deck")
	}

	for _, p := range players {
		assignCard(p, deck)
		assignCard(p, deck)
	}

	assignCard(&d.Player, deck)

	return nil
}

func assignCard(player *Player, deck *Deck) {
	card, _ := deck.DrawCard()
	player.Hand = append(player.Hand, card)
}
