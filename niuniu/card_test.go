package niuniu

import (
	"testing"

	"github.com/ajenpan/poker_algorithm/poker"
)

func TestNiuniu(t *testing.T) {
	type handcards struct {
		cards string
		typ   NNType
	}
	list := []handcards{
		{cards: "8D 2D TD 2C 9H", typ: BULL_ONE},
		{cards: "TD 2C 9H 8D 2D", typ: BULL_ONE},
		{cards: "TD 2C 9H 8D 3D", typ: BULL_TWO},
		{cards: "9D AC AH 8D 3D", typ: BULL_TWO},
		{cards: "TD JC QH KD QD", typ: BULL_BULL},
		{cards: "JD QC KH KD QD", typ: BULL_FIVEFLOWER},
		{cards: "AD AC AH 3D 2D", typ: BULL_FIVESMALL},
		{"AD AC AH AS 2D", BULL_FOURBOMB},
	}

	for _, c := range list {
		cards := NNHandCards{}
		cards.Cards, _ = poker.StringToCards(c.cards)
		cards.Calculate()
		if cards.typ != c.typ {
			t.Logf("card %s type:%v", cards.String(), cards.typ)
			t.Fail()
		}
	}
}
