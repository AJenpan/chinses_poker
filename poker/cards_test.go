package poker

import (
	"fmt"
	"testing"
)

func TestCardsSub(t *testing.T) {
	deck1 := NewDeck()
	deck2 := NewDeck()
	deck1.Sub(deck2)
	if deck1.Size() != 0 {
		t.FailNow()
	}
}

func BenchmarkCardsSubB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deck1 := NewDeck()
		deck2 := NewDeck()
		deck1.Sub(deck2)
	}
}

func TestSort(t *testing.T) {
	deck := NewDeckWithoutJoker()
	deck.SortByRank()
	cs := deck.String()
	fmt.Println(cs)
	deck.SortBySuit()
	cs = deck.String()
	fmt.Println(cs)
	deck.SortByByte()
	cs = deck.String()
	fmt.Println(cs)
}
