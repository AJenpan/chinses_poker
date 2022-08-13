package poker

import (
	"bytes"
	"fmt"
	"strings"
)

//CreateDeck create a pack of cards
func CreateDeck() *Cards {
	deck := NewEmptyCards()
	for i := CardRank(1); i <= 13; i++ {
		deck.BrickCard(CreateCard(DIAMOND, i))
		deck.BrickCard(CreateCard(CLUB, i))
		deck.BrickCard(CreateCard(HEART, i))
		deck.BrickCard(CreateCard(SPADE, i))
	}
	deck.BrickCard(CreateCard(JOKER, 1)) //小王
	deck.BrickCard(CreateCard(JOKER, 2)) //大王
	return deck
}

//CreateDeckWithoutJoker without 2 jokers
func CreateDeckWithoutJoker() *Cards {
	deck := CreateDeck()
	deck.Inner = deck.Inner[:deck.Size()-2]
	return deck
}

func StringToCards(str string) (*Cards, error) {
	str = strings.TrimSpace(str)
	cs := strings.Split(str, " ")
	cards := make([]Card, 0, len(cs))
	for _, v := range cs {
		card := CreateCardByString(v)
		if !card.Valid() {
			return nil, fmt.Errorf("invalid card: %v", v)
		}
		cards = append(cards, card)
	}
	return NewCards(cards), nil
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
