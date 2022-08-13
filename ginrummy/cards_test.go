package ginrummy

import (
	"math/rand"
	"testing"

	"github.com/ajenpan/poker_algorithm/poker"
)

func TestPickTheSet(t *testing.T) {
	randRank := rand.Int31n(13) + 1
	cards := poker.NewEmptyCards()
	for i := 1; i < 5; i++ {
		cards.Push(poker.CreateCard(poker.CardSuit(i), poker.CardRank(randRank)))
	}

	res := PickTheSet(cards)
	if len(res) != 1 {
		t.Errorf("PickTheSet should return 1 set but returned %d", len(res))
		t.FailNow()
	}

}
