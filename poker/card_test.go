package poker

import (
	"math/rand"
	"testing"
	"time"
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
			if card != EmptyCard {
				t.FailNow()
			}
		}
	}
}
func TestPointSize(t *testing.T) {
	card1 := CreateCard(CardSuit(1), CardRank(1))
	card2 := CreateCard(CardSuit(1), CardRank(2))

	if byte(card1) > byte(card2) {
		t.Fail()
	}
}

func TestDeckIsEmpty(t *testing.T) {
	deck := NewEmptyCards()
	if !deck.IsEmpty() {
		t.Error("Deck should be empty but is not.")
	}
}
func TestStr2Cards(t *testing.T) {
	rand.Seed(time.Hour.Nanoseconds())
	for i := 0; i < 100; i++ {
		randSuit := rand.Int31n(4) + 1
		randRank := rand.Int31n(13) + 1
		card := CreateCard(CardSuit(randSuit), CardRank(randRank))
		if card.String() != string([]byte{byte(Ranks[randRank]), byte(Suits[randSuit])}) {
			t.Errorf("%s != %s", card.String(), string([]byte{byte(randSuit), byte(randRank)}))
			t.FailNow()
		}
	}
}
