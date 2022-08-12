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
		cards := NNHandCards{}
		cards.Cards = poker.StringToCards(c.cards)
		cards.Calculate()
		if cards.typ != c.typ {
			t.Logf("card %s type:%v", cards.String(), cards.typ)
			t.Fail()
		}
	}
}
