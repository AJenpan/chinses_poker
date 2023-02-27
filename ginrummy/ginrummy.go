package ginrummy

import (
	"github.com/ajenpan/poker_algorithm/poker"
	"gonum.org/v1/gonum/stat/combin"
)

// set
const minSetDistance int = 3
const maxSetDistance int = 4

// run
const minRunDistance int = 3
const maxRunDistance int = 10

// count
const RankCount = 13
const SuitCount = 4

// Set：3 or 4 of the same rank
func IsTheSet(cards *poker.Cards) bool {
	if cards.Size() < minSetDistance || cards.Size() > maxSetDistance {
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
	if cards.Size() < minRunDistance {
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
	ret := []*poker.Cards{}
	if cards.Size() < minRunDistance {
		return ret
	}
	cards.SortByRank()

	for pos := 0; pos <= cards.Size()-minRunDistance; {
		rank := cards.Get(pos).RankInt()
		suit := cards.Get(pos).Suit()

		searchEndPos := pos + maxRunDistance
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

		if currDistance >= minRunDistance && currDistance <= maxRunDistance {
			ret = append(ret, poker.NewCards(cards.Inner[pos:pos+currDistance]))
		}
		pos += currDistance
	}
	return ret
}

func PickBestSet(cards *poker.Cards) []*poker.Cards {
	ret := []*poker.Cards{}
	if cards.Size() < minSetDistance {
		return ret
	}

	cards.SortBySuit()

	for pos := 0; pos <= cards.Size()-minSetDistance; {
		rank := cards.Get(pos).Rank()

		searchEndPos := pos + maxSetDistance
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

		if currDistance >= minSetDistance && currDistance <= maxSetDistance {
			ret = append(ret, poker.NewCards(cards.Inner[pos:pos+currDistance]))
		}
		pos += currDistance
	}
	return ret
}

func DetectAllRun(cards *poker.Cards) []*poker.Cards {
	ret := []*poker.Cards{}
	if cards.Size() < minRunDistance {
		return ret
	}

	cards.SortBySuit()
	for pos := 0; pos <= cards.Size()-minRunDistance; {
		rank := cards.Get(pos).RankInt()
		suit := cards.Get(pos).Suit()

		searchEndPos := pos + maxRunDistance
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
			if currDistance >= minRunDistance && currDistance <= maxRunDistance {
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
	ret := []*poker.Cards{}
	if cards.Size() < minSetDistance {
		return ret
	}

	cards.SortByRank()

	for pos := 0; pos <= cards.Size()-minSetDistance; {
		rank := cards.Get(pos).Rank()

		searchEndPos := pos + maxSetDistance
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

		newMelds, newDeadwood := DetectBest(cards.Clone().Sub(meld))
		newDeadpoint := CardsPoint(newDeadwood)

		isBetter := (newDeadpoint < bestDeadpoint) || ((newDeadpoint == bestDeadpoint) && (len(newMelds) < len(retMelds)))

		if isBetter {
			retMelds = append(newMelds, meld)
			retDeadwood = newDeadwood
			bestDeadpoint = newDeadpoint
		}
	}
	return retMelds, retDeadwood
}

type CardRate struct {
	Inner [RankCount][SuitCount]float32
}

func (r *CardRate) Init() {
	for i := range r.Inner {
		for j := range r.Inner[i] {
			r.Inner[i][j] = 0
		}
	}
}

func (r *CardRate) Clone() *CardRate {
	ret := &CardRate{}
	ret.Inner = r.Inner
	return ret
}

func (r *CardRate) SetCardRate(c poker.Card, rate float32) {
	if !c.Valid() {
		return
	}
	r.Inner[c.RankInt()-1][c.SuitInt()-1] = rate
}

func (r *CardRate) SetCardsRate(cards *poker.Cards, rate float32) {
	cards.Range(func(i int, c poker.Card) {
		r.SetCardRate(c, rate)
	})
}

func (r *CardRate) GetRate(c poker.Card) float32 {
	if !c.Valid() {
		return 0
	}
	return r.Inner[c.RankInt()-1][c.SuitInt()-1]
}

func (r *CardRate) TheSetRate(c poker.Card) float32 {
	others := make([]poker.Card, 0, 3)
	for suit := poker.CardSuit(1); suit <= 4; suit++ {
		if suit == c.Suit() {
			continue
		}
		others = append(others, poker.NewCard(suit, c.Rank()))
	}
	// hard code is the fast: 0,1 0,2 1,2
	rate1 := r.GetRate(others[0]) * r.GetRate(others[1])
	rate2 := r.GetRate(others[0]) * r.GetRate(others[2])
	rate3 := r.GetRate(others[1]) * r.GetRate(others[2])
	return rate1 + rate2 + rate3
}

func (r *CardRate) TheRunRate(c poker.Card) float32 {
	// to A, the rate is 2's rate * 3's rate
	// to 2, the rate is 3's rate * 4's rate + 1's rate * 3's rate
	// to 3, the rate is 4's rate * 5's rate + 2's rate * 4's rate + 1's rate * 2's rate
	ret := float32(0.0)
	for _, v := range [][]int8{{-2, -1}, {-1, 1}, {1, 2}} {
		card1 := poker.NewCard(c.Suit(), poker.CardRank(int8(c.Rank())+v[0]))
		card2 := poker.NewCard(c.Suit(), poker.CardRank(int8(c.Rank())+v[1]))
		ret += r.GetRate(card1) * r.GetRate(card2)
	}

	return ret
}

func (r *CardRate) BeMeldRate(c poker.Card) float32 {
	return r.TheRunRate(c) + r.TheSetRate(c)
}
