package poker

import (
	"testing"
)

func TestCard(t *testing.T) {

	for s := 1; s <= 4; s++ {
		for r := 1; r <= 13; r++ {
			card := CreateCard(CardSuit(s), CardRank(r))
			if int(card.Suit()) != s || int(card.Rank()) != r {
				t.FailNow()
			}
		}
	}

	for s := 1; s <= 4; s++ {
		for r := 14; r <= 26; r++ {
			card := CreateCard(CardSuit(s), CardRank(r))
			if int(card.Suit()) == s {
				t.FailNow()
			}
			if card != ErrorCard {
				t.FailNow()
			}
		}
	}
}

func TestString2Card(t *testing.T) {

}
