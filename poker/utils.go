package poker

import (
	"bytes"
	"fmt"
	"strings"
)

// NewDeck create a pack of cards
func NewDeck() *Cards {
	deck := NewEmptyCards()
	for i := CardRank(1); i <= 13; i++ {
		deck.BrickCard(NewCard(DIAMOND, i))
		deck.BrickCard(NewCard(CLUB, i))
		deck.BrickCard(NewCard(HEART, i))
		deck.BrickCard(NewCard(SPADE, i))
	}
	deck.BrickCard(NewCard(JOKER, 1)) //小王
	deck.BrickCard(NewCard(JOKER, 2)) //大王
	return deck
}

// NewDeckWithoutJoker without 2 jokers
func NewDeckWithoutJoker() *Cards {
	deck := NewDeck()
	deck.Inner = deck.Inner[:deck.Size()-2]
	return deck
}

func StringArrToCards(cs []string) (*Cards, error) {
	cards := make([]Card, 0, len(cs))
	for _, v := range cs {
		card := NewCardByString(v)
		if !card.Valid() {
			return nil, fmt.Errorf("invalid card: %v", v)
		}
		cards = append(cards, card)
	}
	return NewCards(cards), nil
}

func StringToCards(str string) (*Cards, error) {
	str = strings.TrimSpace(str)
	cs := strings.Split(str, " ")
	return StringArrToCards(cs)
}

func BytesToCards(raw []byte) (*Cards, error) {
	raw = bytes.TrimSpace(raw)
	cards := make([]Card, 0, len(raw))
	for _, v := range raw {
		card := ByteToCard(v)
		if !card.Valid() {
			return nil, fmt.Errorf("invalid card: %v", v)
		}
		cards = append(cards, card)
	}
	return NewCards(cards), nil
}
