package poker

import (
	"fmt"
	"math/rand"
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

func TestDealCards(t *testing.T) {
	deck := NewDeck()
	decksize := deck.Size()
	n := 5
	dealtCards := deck.DealCards(n)
	if dealtCards.Size() != n {
		t.Errorf("expected %d cards, got %d", n, dealtCards.Size())
	}
	if deck.Size() != decksize-n {
		t.Errorf("expected %d cards left in deck, got %d", decksize-n, deck.Size())
	}
}

func TestRemoveCards2(t *testing.T) {
	deck := NewDeckWithoutJoker()
	decksize := deck.Size()
	dealcnt := rand.Int31n(int32(decksize))

	cardsToRemove := deck.Clone().DealCards(int(dealcnt))
	cardsToRemove.Push(NewCard(JOKER, 1))

	success := deck.RemoveCards(cardsToRemove)
	if success {
		t.Errorf("expected removal to be successful")
	}
	if deck.Size() != decksize {
		t.Errorf("expected %d cards left in deck, got %d", 52-5, deck.Size())
	}
}

func TestRemoveCards(t *testing.T) {
	deck := NewDeckWithoutJoker()
	decksize := deck.Size()
	dealcnt := rand.Int31n(int32(decksize))

	cardsToRemove := deck.Clone().DealCards(int(dealcnt))
	success := deck.RemoveCards(cardsToRemove)
	if !success {
		t.Errorf("expected removal to be successful")
	}
	if deck.Size() != decksize-int(dealcnt) {
		t.Errorf("expected %d cards left in deck, got %d", 52-5, deck.Size())
	}

	// Test removing cards that are not in the deck
	nonExistentCards := NewEmptyCards()
	nonExistentCards.Push(NewCard(JOKER, 1))
	success = deck.RemoveCards(nonExistentCards)
	if success {
		t.Errorf("expected removal to be unsuccessful")
	}
	if deck.Size() != decksize-int(dealcnt) {
		t.Errorf("expected %d cards left in deck, got %d", decksize-int(dealcnt), deck.Size())
	}
}

func BenchmarkRemoveCards(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deck := NewDeck()
		cardsToRemove := deck.DealCards(5)
		deck.RemoveCards(cardsToRemove)
	}
}

func BenchmarkDealCards(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deck := NewDeck()
		deck.DealCards(5)
	}
}
