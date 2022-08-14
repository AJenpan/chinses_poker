package ginrummy

import (
	"github.com/ajenpan/poker_algorithm/poker"
)

type DeckRate [13][4]float32

type HandCards struct {
	Cards        *poker.Cards
	DiscardStack []poker.Card

	ExistRate DeckRate
}

func NewHandCards(cards *poker.Cards) *HandCards {
	ret := &HandCards{Cards: cards}
	return ret
}

// when other discard one, need to update the rate and push the card to discardstack
func (d *HandCards) OnDiscard(c poker.Card) {
	d.DiscardStack = append(d.DiscardStack, c)
	d.ExistRate[c.RankInt()-1][c.Suit()-1] = 0
}

func (d *HandCards) UpdateDeckRate() {
	// the full deck, sub discard deck, melds card
}

func (d *HandCards) DiscardScore(c poker.Card) float32 {
	const DiscardScoreFix = float32(1000)
	return (float32(1)-d.HoldpowerRate(c))*10 + (float32(CardPoint(c)) / DiscardScoreFix)
}

// the rate = set rate + run rate, the higher the rate, the better the hand
func (d *HandCards) HoldpowerRate(c poker.Card) float32 {
	return d.TheRunRate(c) + d.TheSetRate(c)
}

func (d *HandCards) CardGotRate(c poker.Card) float32 {
	if !c.Valid() {
		return 0
	}
	return d.ExistRate[c.RankInt()-1][c.Suit()-1]
}

func (d *HandCards) TheSetRate(c poker.Card) float32 {
	others := make([]poker.Card, 0, 3)
	for suit := poker.CardSuit(1); suit <= 4; suit++ {
		if suit == c.Suit() {
			continue
		}
		others = append(others, poker.NewCard(suit, c.Rank()))
	}
	// opts := combin.Combinations(3, 2)
	// 0,1 0,2 1,2
	rate1 := d.CardGotRate(others[0]) * d.CardGotRate(others[1])
	rate2 := d.CardGotRate(others[0]) * d.CardGotRate(others[2])
	rate3 := d.CardGotRate(others[1]) * d.CardGotRate(others[2])
	return rate1 + rate2 + rate3
}

func (d *HandCards) TheRunRate(c poker.Card) float32 {
	// to A, the rate is 2's rate * 3's rate
	// to 2, the rate is 3's rate * 4's rate + 1's rate * 3's rate
	// to 3, the rate is 4's rate * 5's rate + 2's rate * 4's rate + 1's rate * 2's rate
	ret := float32(0.0)
	for _, v := range [][]int8{{-2, -1}, {-1, 1}, {1, 2}} {
		card1 := poker.NewCard(c.Suit(), poker.CardRank(int8(c.Rank())+v[0]))
		card2 := poker.NewCard(c.Suit(), poker.CardRank(int8(c.Rank())+v[1]))
		ret += d.CardGotRate(card1) * d.CardGotRate(card2)
	}
	return ret
}

func (d *HandCards) DiscardOne(cards *poker.Cards) poker.Card {
	if cards.IsEmpty() {
		return poker.EmptyCard
	}
	badest := cards.Inner[0]
	badestRate := d.HoldpowerRate(badest)
	for i := 1; i < cards.Size(); i++ {
		c := cards.Get(i)
		rate := d.HoldpowerRate(c)
		if rate < badestRate {
			badest = c
			badestRate = rate
		}
	}
	return badest
}

func (d *HandCards) CheckNeed(c poker.Card) poker.Card {
	beforeMelds, beforeDeadwood := DetectBest(d.Cards)
	beforeScore := CardsPoint(beforeDeadwood)

	newCards := d.Cards.Clone()
	newCards.Push(c)

	newMelds, newDeadwood := DetectBest(newCards)
	discard := d.DiscardOne(newDeadwood)
	newDeadwood.RemoveCard(discard)

	if len(newMelds) > len(beforeMelds) {
		return discard
	}
	newScore := CardsPoint(newDeadwood)
	if newScore < beforeScore {
		return discard
	}
	return c
}
