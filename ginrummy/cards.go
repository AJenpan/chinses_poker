package ginrummy

import (
	"github.com/ajenpan/poker_algorithm/poker"
)

type Deck struct {
	*poker.Cards
}

func NewDeck() *Deck {
	ret := &Deck{
		Cards: poker.CreateDeckWithoutJoker(),
	}
	return ret
}

// Set：3 or 4 of the same rank
func IsTheSet(cards *poker.Cards) bool {
	if cards.Size() < 3 || cards.Size() > 4 {
		return false
	}
	for i := 1; i < cards.Size(); i++ {
		if cards.Get(i).Rank() != cards.Get(0).Rank() {
			return false
		}
	}
	return true
}

// Run: Three or more suited cards in sequence.
func IsTheRun(cards *poker.Cards) bool {
	if cards.Size() < 3 {
		return false
	}
	for i := 1; i < cards.Size(); i++ {
		if cards.Get(i).Suit() != cards.Get(0).Suit() {
			return false
		}
		if cards.Get(i).Rank() != cards.Get(i-1).Rank()+1 {
			return false
		}
	}
	return true
}

// Meld: Set or Run
func IsMeld(cards *poker.Cards) bool {
	return IsTheSet(cards) || IsTheRun(cards)
}

// Point: A=1 point，J/Q/K=10 point, and others according to their numerical values.
func Point(card poker.Card) int {
	v := card.RankInt()
	if v >= 10 {
		return 10
	}
	return v
}

func PickTheRun(cards *poker.Cards) []*poker.Card {
	cards.SortByByte()

	const distance = 3
	if cards.Size() < distance {
		return nil
	}

	for i := 0; i < cards.Size()-distance; i++ {
		if cards.Get(i).Rank() != cards.Get(i+distance-1).Rank()+1 {
			continue
		}
		for j := i; j < i+distance; j++ {
			if cards.Get(j).Rank() != cards.Get(i).Rank() {
				return nil
			}
		}

	}

	return nil
}

func PickTheSet(cards *poker.Cards) []*poker.Cards {
	cards.SortByByte()

	const distance1 = 3
	const distance2 = 4

	ret := []*poker.Cards{}

	if cards.Size() < distance1 {
		return ret
	}

	startPos := 0
	for startPos < cards.Size()-distance1 {
		rank := cards.Get(startPos).Rank()
		currDistance := 1

		for i := startPos + 1; i < startPos+distance2; i++ {
			if cards.Get(i).Rank() != rank {
				break
			}
			currDistance++
		}
		if currDistance == distance1 || currDistance == distance2 {
			ret = append(ret, poker.NewCards(cards.Inner[startPos:startPos+currDistance]))
		}
		startPos++
	}
	return ret
}

// SubsetOfMeld - checks if a meld being made is a subset of a previous meld
// made. Ex. 2C 3C 4C is a subset of 2C 3C 4C 5C 6C. This is a fixed length
// linear search.
func SubOfMeld(melds, m *poker.Cards) bool {
	return false
}
