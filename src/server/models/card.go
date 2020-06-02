package models

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

func NewCard(suit CardSuit, value uint32, cardType CardType) Card {
	var card Card

	if (value == 1 || value == 11) && cardType == PIP {
		card = Card{suit, []uint32{1, 11}, cardType}
	}

	if cardType == PIP && (value >= 2 && value <= 10) {
		card = Card{suit, []uint32{value}, cardType}
	}

	cardTypeMap := map[CardType]CardType{KING: KING, QUEEN: QUEEN, JACK: JACK}

	if _, ok := cardTypeMap[cardType]; value == 10 && ok {
		card = Card{suit, []uint32{value}, cardType}
	}

	return card
}

func (c Card) IsAce() bool {
	return len(c.Values) == 2 && (c.Values[0] == 1 && c.Values[1] == 11) || (c.Values[0] == 11 && c.Values[1] == 1)
}