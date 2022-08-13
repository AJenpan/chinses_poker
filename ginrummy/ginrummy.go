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

// Run: 3 or more same suited cards in sequence.
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
func CardPoint(card poker.Card) int {
	v := card.RankInt()
	if v >= 10 {
		return 10
	}
	return v
}

func CardsPoint(cards *poker.Cards) int {
	ret := 0
	for i := 0; i < cards.Size(); i++ {
		ret += CardPoint(cards.Get(i))
	}
	return ret
}

func PickBestRun(cards *poker.Cards) []*poker.Cards {
	const minDistance = 3
	const maxDistance = 10

	ret := []*poker.Cards{}
	if cards.Size() < minDistance {
		return ret
	}

	cards.SortByRank()

	for pos := 0; pos <= cards.Size()-minDistance; {
		rank := cards.Get(pos).RankInt()
		suit := cards.Get(pos).Suit()

		searchEndPos := pos + maxDistance
		if searchEndPos > cards.Size() {
			searchEndPos = cards.Size()
		}

		currDistance := 1
		for i := pos + 1; i < searchEndPos; i++ {
			if suit != cards.Get(i).Suit() {
				break
			}
			if rank+i-pos != cards.Get(i).RankInt() {
				break
			}
			currDistance++
		}

		if currDistance >= minDistance && currDistance <= maxDistance {
			ret = append(ret, poker.NewCards(cards.Inner[pos:pos+currDistance]))
		}
		pos += currDistance
	}
	return ret
}

func PickBestSet(cards *poker.Cards) []*poker.Cards {
	const minDistance = 3
	const maxDistance = 4

	ret := []*poker.Cards{}
	if cards.Size() < minDistance {
		return ret
	}

	cards.SortBySuit()

	for pos := 0; pos <= cards.Size()-minDistance; {
		rank := cards.Get(pos).Rank()

		searchEndPos := pos + maxDistance
		if searchEndPos > cards.Size() {
			searchEndPos = cards.Size()
		}

		currDistance := 1
		for i := pos + 1; i < searchEndPos; i++ {
			if cards.Get(i).Rank() != rank {
				break
			}
			currDistance++
		}

		if currDistance >= minDistance && currDistance <= maxDistance {
			ret = append(ret, poker.NewCards(cards.Inner[pos:pos+currDistance]))
		}
		pos += currDistance
	}
	return ret
}

func DetectAllRun(cards *poker.Cards) []*poker.Cards {
	const minDistance = 3
	const maxDistance = 10

	ret := []*poker.Cards{}
	if cards.Size() < minDistance {
		return ret
	}

	cards.SortBySuit()
	for pos := 0; pos <= cards.Size()-minDistance; {
		rank := cards.Get(pos).RankInt()
		suit := cards.Get(pos).Suit()

		searchEndPos := pos + maxDistance
		if searchEndPos > cards.Size() {
			searchEndPos = cards.Size()
		}

		currDistance := 1
		for i := pos + 1; i < searchEndPos; i++ {
			if suit != cards.Get(i).Suit() {
				break
			}
			if rank+i-pos != cards.Get(i).RankInt() {
				break
			}
			currDistance++
			if currDistance >= minDistance && currDistance <= maxDistance {
				ret = append(ret, poker.NewCards(cards.Inner[pos:pos+currDistance]).Copy())
			}
		}
		pos += 1
	}
	return ret
}

func DetectAllSet(cards *poker.Cards) []*poker.Cards {

	const minDistance = 3
	const maxDistance = 4

	ret := []*poker.Cards{}
	if cards.Size() < minDistance {
		return ret
	}

	cards.SortByRank()

	for pos := 0; pos <= cards.Size()-minDistance; {
		rank := cards.Get(pos).Rank()

		searchEndPos := pos + maxDistance
		if searchEndPos > cards.Size() {
			searchEndPos = cards.Size()
		}

		currDistance := 1
		for i := pos + 1; i < searchEndPos; i++ {
			if cards.Get(i).Rank() != rank {
				break
			}
			currDistance++
		}

		if currDistance >= minDistance && currDistance <= maxDistance {
			ret = append(ret, poker.NewCards(cards.Inner[pos:pos+currDistance]).Copy())
		}

		pos += currDistance
	}
	return ret
}

// return melds and deadwood
func DetectBest(cards *poker.Cards) ([]*poker.Cards, *poker.Cards) {
	if cards.Size() < 3 {
		return nil, cards
	}

	runs := DetectAllRun(cards)
	sets := DetectAllSet(cards)
	alls := append(runs, sets...)
	if len(alls) == 0 {
		return nil, cards
	}

	temp, retDeadwood := DetectBest(cards.Copy().Sub(alls[0]))
	retMelds := append([]*poker.Cards{alls[0]}, temp...)
	bestDeadpoint := CardsPoint(retDeadwood)

	for i := 1; i < len(alls); i++ {
		meld := alls[i]

		newTemp, newDeadwood := DetectBest(cards.Copy().Sub(meld))
		newDeadpoint := CardsPoint(newDeadwood)

		if newDeadpoint < bestDeadpoint {
			retMelds = append([]*poker.Cards{meld}, newTemp...)
			retDeadwood = newDeadwood
			bestDeadpoint = newDeadpoint
		}
	}
	return retMelds, retDeadwood
}

// func PickDeadwood(cards *poker.Cards) poker.Card {
// 	if cards.Size() < 1 {
// 		return poker.EmptyCard
// 	}
// 	melds, deadwood := DetectBest(cards)
// 	return deadwood
// }
