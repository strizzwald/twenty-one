package models

import (
	"errors"
	"fmt"
)

type CardSuit string

type CardType string

const (
	HEARTS CardSuit = "Hearts"
	SPADES CardSuit = "Spades"
	CLUBS CardSuit = "Clubs"
	DIAMONDS CardSuit = "Diamonds"
)

const (
	PIP CardType = "Pip"
	KING CardType = "King"
	QUEEN CardType = "QUEEN"
	JACK CardType = "JACK"
)

type Card struct {
	Suit     CardSuit
	Values   []uint32
	CardType CardType
}

func NewCard(suit CardSuit, value uint32, cardType CardType) (Card, error) {

	if (value == 1 || value == 11) && cardType == PIP {
		return Card{suit, []uint32{1, 11}, cardType}, nil
	}

	if cardType == PIP && (value >= 2 && value <= 10) {
		return Card{suit, []uint32{value}, cardType}, nil
	}

	cardTypeMap := map[CardType]CardType{KING: KING, QUEEN: QUEEN, JACK: JACK}

	if _, ok := cardTypeMap[cardType]; value == 10 && ok {
		return Card{suit, []uint32{value}, cardType}, nil
	}

	return Card{}, errors.New(fmt.Sprintf("%d of %s with type %s is an invalid card", value, suit, cardType))
}

func (c Card) IsAce() bool {
	return len(c.Values) == 2 && (c.Values[0] == 1 && c.Values[1] == 11) || (c.Values[0] == 11 && c.Values[1] == 1)
}