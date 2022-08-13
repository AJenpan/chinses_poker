package ginrummy

import (
	"math/rand"
	"testing"
	"time"

	"github.com/ajenpan/poker_algorithm/poker"
)

func TestPickTheSet(t *testing.T) {
	randRank := rand.Int31n(12) + 2
	cards := poker.NewEmptyCards()
	for i := 1; i < 5; i++ {
		cards.Push(poker.CreateCard(poker.CardSuit(i), poker.CardRank(randRank)))
	}

	res := PickBestSet(cards)
	if len(res) != 1 {
		t.Errorf("PickTheSet should return 1 set but returned %d", len(res))
		t.FailNow()
		return
	}

	if res[0].Size() != 4 {
		t.Error("PickTheSet should return 4 cards but returned ", res[0].Size())
		t.FailNow()
		return
	}

	cards.Clear()
	cards.Push(poker.CreateCard(poker.DIAMOND, poker.CardRank(1)))
	for i := 1; i <= 3; i++ {
		cards.Push(poker.CreateCard(poker.CardSuit(i), poker.CardRank(randRank)))
	}
	if len(PickBestSet(cards)) != 1 {
		t.Errorf("PickTheSet should return 1 set but returned %d", len(res))
		t.FailNow()
		return
	}

	cards.Clear()
	if len(PickBestSet(cards)) != 0 {
		t.FailNow()
		return
	}

	cards.Push(poker.CreateCard(poker.DIAMOND, poker.CardRank(1)))
	cards.Push(poker.CreateCard(poker.CLUB, poker.CardRank(1)))

	if len(PickBestSet(cards)) != 0 {
		t.FailNow()
		return
	}
	cards.Push(poker.CreateCard(poker.SPADE, poker.CardRank(1)))
	if len(PickBestSet(cards)) != 1 {
		t.FailNow()
		return
	}
}

func TestPickTheRun(t *testing.T) {

	cards := poker.CreateDeckWithoutJoker()
	cards.SortBySuit()

	rand.Seed(time.Now().UnixNano())

	pos := rand.Int31n(50) + 1
	expLen := int32(3)

	var startPos int32
	var endPos int32

	if pos > 25 {
		startPos = pos - expLen
		endPos = pos
	} else {
		startPos = pos
		endPos = pos + expLen
	}
	cards.Inner = cards.Inner[startPos:endPos]

	cc := cards.String()
	t.Log(cc)

	expectHas := (cards.Get(0).Suit() == cards.Get(cards.Size()-1).Suit()) && cards.Size() >= 3
	res := PickBestRun(cards)

	if expectHas != (len(res) == 1) {
		t.Errorf("PickTheRun should return %v, but returned %v", expectHas, (len(res) == 1))
		t.FailNow()
		return
	}
}

func TestDetectBest(t *testing.T) {
	cards, err := poker.StringToCards("AD AS AC 2C 3C 4C")
	if err != nil {
		t.FailNow()
		return
	}
	t.Log(cards.String())
	cs := cards.Chinese()
	t.Log(cs)
	melds, deadwood := DetectBest(cards)
	if len(melds) != 2 || deadwood.Size() != 0 {
		t.FailNow()
		return
	}
}

func BenchmarkDetectBestAg(b *testing.B) {
	cards, err := poker.StringToCards("AD AS AC 2C 3C 4C")
	if err != nil {
		return
	}
	for i := 0; i < b.N; i++ {
		DetectBest(cards)
	}
}

func BenchmarkDetectBestRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cards := poker.CreateDeckWithoutJoker()
		cards.Shuffle()
		cards.Inner = cards.Inner[0:10]
		DetectBest(cards)
	}
}

func TestDoBestAction(t *testing.T) {
	cards := poker.CreateDeckWithoutJoker()
	cards.Shuffle()
	cards.Inner = cards.Inner[0:10]
	deck := poker.NewCards(cards.Inner[10:]).Copy()

	for _, v := range deck.Inner {
		deck.RemoveCard(v)
	}
}
