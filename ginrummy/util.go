package ginrummy

import (
	"github.com/ajenpan/poker_algorithm/poker"
	"gonum.org/v1/gonum/stat/combin"
)

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

// sum of all cards point
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
				ret = append(ret, poker.NewCards(cards.Inner[pos:pos+currDistance]).Clone())
			}
		}
		pos += 1
	}
	return ret
}

var c43 = combin.Combinations(4, 3)

func suitC43[T any](m []T, fn func([]T)) {
	res := make([]T, 3)
	for _, c := range c43 {
		res = res[0:0]
		for _, v := range c {
			res = append(res, m[v])
		}
		fn(res)
	}
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
		if currDistance == 3 {
			ret = append(ret, poker.NewCards(cards.Inner[pos:pos+currDistance]).Clone())
		} else if currDistance == 4 {
			target := poker.NewCards(cards.Inner[pos : pos+currDistance]).Clone()
			suitC43(target.Inner, func(res []poker.Card) {
				t := poker.NewCards(res).Clone()
				ret = append(ret, t)
			})
			ret = append(ret, target)
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

	temp, retDeadwood := DetectBest(cards.Clone().Sub(alls[0]))
	retMelds := append([]*poker.Cards{alls[0]}, temp...)
	bestDeadpoint := CardsPoint(retDeadwood)

	for i := 1; i < len(alls); i++ {
		meld := alls[i]

		newTemp, newDeadwood := DetectBest(cards.Clone().Sub(meld))
		newDeadpoint := CardsPoint(newDeadwood)

		if newDeadpoint < bestDeadpoint {
			retMelds = append([]*poker.Cards{meld}, newTemp...)
			retDeadwood = newDeadwood
			bestDeadpoint = newDeadpoint
		}
	}
	return retMelds, retDeadwood
}

func DetectAllRunV2(cards [13][4]poker.Card) []*poker.Cards {
	const minDistance = 3
	const maxDistance = 10

	ret := []*poker.Cards{}

	for iSuit := 0; iSuit < 4; iSuit++ {
		for pos := 0; pos <= 13-minDistance; {
			if cards[pos][iSuit].Valid() {
				pos++
				continue
			}
			rank := cards[pos][iSuit].RankInt()
			currDistance := 1
			for i := pos + 1; i < 13; i++ {
				target := cards[i][iSuit]
				if target.Valid() {
					break
				}
				if target.RankInt()+(pos-i) != rank {
					break
				}

				currDistance++
				if currDistance >= minDistance && currDistance <= maxDistance {
					cc := []poker.Card{}
					for slipt := pos; slipt <= pos+currDistance; slipt++ {
						cc = append(cc, cards[slipt][iSuit])
					}
					ret = append(ret, poker.NewCards(cc))
				}
			}
			pos++
		}
	}
	return ret
}

func DetectAllSet2(Cards [13][4]poker.Card) []*poker.Cards {
	ret := []*poker.Cards{}
	cards := make([]poker.Card, 0, 4)
	for iRank := 0; iRank < 13; iRank++ {
		cards = cards[0:0]
		for iSuit := 0; iSuit <= 4; iSuit++ {
			if Cards[iRank][iSuit].Valid() {
				continue
			}
			cards = append(cards, Cards[iRank][iSuit])
		}
		if len(cards) == 4 {
			suitC43(cards, func(t []poker.Card) {
				ret = append(ret, poker.NewCards(t).Clone())
			})
			ret = append(ret, poker.NewCards(cards).Clone())
		} else if len(cards) == 3 {
			ret = append(ret, poker.NewCards(cards).Clone())
		}
	}
	return ret
}
