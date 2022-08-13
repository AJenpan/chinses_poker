package poker

import (
	"fmt"
	"testing"
)

func TestCardsSub(t *testing.T) {
	deck1 := CreateDeck()
	deck2 := CreateDeck()
	deck1.Sub(deck2)
	if deck1.Size() != 0 {
		t.FailNow()
	}
}

func BenchmarkCardsSubB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deck1 := CreateDeck()
		deck2 := CreateDeck()
		deck1.Sub(deck2)
	}
}

func TestSort(t *testing.T) {
	deck := CreateDeckWithoutJoker()
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
