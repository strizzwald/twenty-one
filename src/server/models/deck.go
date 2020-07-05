package models

import (
	"errors"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	d := Deck{}
	d.Cards = []Card{}

	suits := []CardSuit{HEARTS, SPADES, CLUBS, DIAMONDS}
	cardTypes := []CardType{KING, QUEEN, JACK}
	var newCard Card

	for _, s := range suits {
		for v := 1; v <= 10; v++ {
			newCard = NewCard(s, uint32(v), PIP)

			d.Cards = append(d.Cards, newCard)
		}

		for _, t := range cardTypes {
			newCard = NewCard(s, 10, t)

			d.Cards = append(d.Cards, newCard)
		}
	}

	return d
}

func (d *Deck) Shuffle() {
	for i := len(d.Cards) - 1; i > 1; i-- {
		rand.Seed(time.Now().UTC().UnixNano())

		j := rand.Intn(len(d.Cards) - 1)

		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func (d *Deck) HasCards() bool {
	return len(d.Cards) != 0
}

func (d *Deck) DrawCard() (Card, error) {
	if !d.HasCards() {
		return Card{}, errors.New("deck is empty")
	}

	if len(d.Cards) > 1 {
		d.Shuffle()
	}

	card := d.Cards[0]
	d.Cards = d.Cards[1:]

	return card, nil
}
