package models

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func NewDeck() (Deck, error) {
	d := Deck{}
	d.Cards = []Card{}

	suits := []CardSuit{HEARTS, SPADES, CLUBS, DIAMONDS}
	cardTypes := []CardType{KING, QUEEN, JACK}

	var newCard Card
	var err error

	for _, s := range suits {
		for v := 1; v <= 10; v++ {
			newCard, err = NewCard(s, uint32(v), PIP)

			if err != nil {
				return Deck{}, err
			}

			d.Cards = append(d.Cards, newCard)
		}

		for _, t := range cardTypes {
			newCard, err = NewCard(s, 10, t)

			if err != nil {
				return Deck{}, err
			}

			d.Cards = append(d.Cards, newCard)
		}
	}

	return d, nil
}

func (d Deck) Shuffle() Deck {
	for i := len(d.Cards) - 1;  i > 1; i-- {
		rand.Seed(time.Now().UTC().UnixNano())

		j := rand.Intn(len(d.Cards) - 1)

		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}

	return d
}
