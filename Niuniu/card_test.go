package Niuniu

import (
	"testing"

	"github.com/Ajenpan/chinese_poker_go/poker"
)

func TestNiuniu(t *testing.T) {
	cards := NNHandCards{}
	type handcards struct {
		cards string
		typ   NNType
	}
	list := []handcards{
		{cards: "D8 D2 DT C2 H9", typ: BULL_ONE},
		{cards: "DT C2 H9 D8 D2", typ: BULL_ONE},
		{cards: "DT C2 H9 D8 D3", typ: BULL_TWO},
		{cards: "D9 CA HA D8 D3", typ: BULL_TWO},
		{cards: "DT CJ HQ DK DQ", typ: BULL_BULL},
		{cards: "DJ CQ HK DK DQ", typ: BULL_FIVEFLOWER},
		{cards: "DA CA HA D3 D2", typ: BULL_FIVESMALL},
		{"DA CA HA SA D2", BULL_FOURBOMB},
	}

	for _, c := range list {
		cards.Cards = poker.StringToCards(c.cards)
		cards.calculate()
		if cards.typ != c.typ {
			t.Fail()
		}
	}
}
