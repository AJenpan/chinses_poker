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
		cards.Push(poker.NewCard(poker.CardSuit(i), poker.CardRank(randRank)))
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
	cards.Push(poker.NewCard(poker.DIAMOND, poker.CardRank(1)))
	for i := 1; i <= 3; i++ {
		cards.Push(poker.NewCard(poker.CardSuit(i), poker.CardRank(randRank)))
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

	cards.Push(poker.NewCard(poker.DIAMOND, poker.CardRank(1)))
	cards.Push(poker.NewCard(poker.CLUB, poker.CardRank(1)))

	if len(PickBestSet(cards)) != 0 {
		t.FailNow()
		return
	}
	cards.Push(poker.NewCard(poker.SPADE, poker.CardRank(1)))
	if len(PickBestSet(cards)) != 1 {
		t.FailNow()
		return
	}
}

func TestPickTheRun(t *testing.T) {
	cards := poker.NewDeckWithoutJoker()
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
func TestDetectAllSet(t *testing.T) {
	cards, err := poker.StringToCards("DA SA CA HA")
	if err != nil {
		t.FailNow()
		return
	}
	t.Log(cards.String())
	cs := cards.Chinese()
	t.Log(cs)
	res := DetectAllSet(cards)
	if len(res) != 4+1 {
		t.FailNow()
		return
	}
}

func TestDetectBest(t *testing.T) {
	cards, err := poker.StringToCards("SA S2 S3 C4 S4 D4 S5 S6 S7 S9")
	if err != nil {
		t.FailNow()
		return
	}
	melds, deadwood := DetectBest(cards)
	for _, v := range melds {
		t.Log(v.Chinese())
	}
	if deadwood != nil {
		t.Log(deadwood.Chinese())
	}
	if len(melds) != 3 || deadwood.Size() != 1 {
		t.FailNow()
		return
	}
}

func BenchmarkDetectBestAg(b *testing.B) {
	cards, err := poker.StringToCards("DA SA CA C2 C3 C4")
	if err != nil {
		return
	}
	for i := 0; i < b.N; i++ {
		DetectBest(cards)
	}
}

func BenchmarkDetectBestRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cards := poker.NewDeckWithoutJoker()
		cards.Shuffle()
		cards.Inner = cards.Inner[0:10]
		DetectBest(cards)
	}
}

func TestDoBestAction(t *testing.T) {
	cards := poker.NewDeckWithoutJoker()
	cards.Shuffle()
	cards.Inner = cards.Inner[0:10]
	deck := poker.NewCards(cards.Inner[10:]).Clone()

	for _, v := range deck.Inner {
		deck.RemoveCard(v)
	}
}

func TestDiscardOne(t *testing.T) {
	allcards := poker.NewDeckWithoutJoker()
	cards, err := poker.StringToCards("DA S2 DK SK")
	if err != nil {
		return
	}
	existRate := &CardRate{}
	restCount := allcards.Size() - cards.Size()
	for _, c := range allcards.Inner {
		if cards.Contain(c) {

			existRate.SetCardRate(c, 1.0)
		} else {
			existRate.SetCardRate(c, float32(1.0)/float32(restCount))

		}
	}

}
