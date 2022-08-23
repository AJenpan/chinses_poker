package niuniu

import (
	"github.com/ajenpan/poker_algorithm/poker"
)

const maxHandCardCout = 5

type NNType uint8

const (
	BULL_ZERO       NNType = 0  //没牛
	BULL_ONE        NNType = 1  //牛一 //牛丁
	BULL_TWO        NNType = 2  //牛二
	BULL_THREE      NNType = 3  //牛三
	BULL_FOUR       NNType = 4  //牛四
	BULL_FIVE       NNType = 5  //牛五
	BULL_SIX        NNType = 6  //牛六
	BULL_SEVEN      NNType = 7  //牛七
	BULL_EIGHT      NNType = 8  //牛八
	BULL_NINE       NNType = 9  //牛九
	BULL_BULL       NNType = 10 //牛牛
	BULL_FOURBOMB   NNType = 11 //炸弹牛 //四炸 //5张牌中有4张一样的牌，此时无需有牛
	BULL_FIVEFLOWER NNType = 12 //五花牛 //指的是JQK，五花牛指的是手上的5张牌全为JQK的特殊牛牛牌型
	BULL_FIVESMALL  NNType = 13 //五小牛 //五张牌都小余5，且牌点总数小余或等于10
	ERROR_TYPE      NNType = 255
)

type NNHandCards struct {
	*poker.Cards
	typ       NNType
	power     poker.Card
	calculate bool
}

type NNDeck struct {
	*poker.Cards
}

func NewNNDeck() *NNDeck {
	return &NNDeck{
		Cards: poker.NewDeckWithoutJoker(),
	}
}

func (nn *NNDeck) DealHandCards() *NNHandCards {
	cards := nn.DealCards(maxHandCardCout)
	ret := &NNHandCards{}
	ret.Cards = cards
	ret.calculate = false
	return ret
}

func IsBull(cards *poker.Cards) bool {
	if cards.Size() != 3 {
		return false
	}
	return calculatePoint(cards)%10 == 0
}

func (c *NNHandCards) Type() NNType {
	return c.typ
}

func (c *NNHandCards) Power() poker.Card {
	return c.power
}

// Compare return c >= other?
func (c *NNHandCards) Compare(other *NNHandCards) bool {
	if c.typ == ERROR_TYPE || other.typ == ERROR_TYPE {
		panic("error type")
	}
	if c.typ > other.typ {
		return true
	} else if c.typ == other.typ {
		return CompareCard(c.power, other.power)
	}
	return false
}

func (c *NNHandCards) isFourBomb() bool {
	count := make(map[int]int)
	for _, v := range c.Cards.Inner {
		if _, ok := count[v.RankInt()]; ok {
			count[v.RankInt()]++
		} else {
			count[v.RankInt()] = 1
		}
	}
	for r, v := range count {
		if v == 4 {
			c.typ = BULL_FOURBOMB
			c.power = poker.NewCard(poker.SPADE, poker.CardRank(r))
			return true
		}
	}
	return false
}

func (c *NNHandCards) isFiveSmall() bool {
	point := calculatePoint(c.Cards)
	if point > 10 {
		return false
	}
	for _, v := range c.Cards.Inner {
		if v.RankInt() > 5 {
			return false
		}
	}
	c.typ = BULL_FIVESMALL
	return true
}

func (c *NNHandCards) isFiveFlower() bool {
	if calculatePoint(c.Cards) != 50 {
		return false
	}
	for _, v := range c.Cards.Inner {
		if v.RankInt() <= 10 {
			return false
		}
	}
	c.typ = BULL_FIVEFLOWER
	return true
}

func (c *NNHandCards) IsCalculate() bool {
	return c.calculate
}

func (c *NNHandCards) Calculate() {
	if c.calculate {
		return
	}
	c.calculate = true
	if c.isFiveSmall() {
		return
	}
	if c.isFourBomb() {
		return
	}
	if c.isFiveFlower() {
		return
	}

	ret := &[]*poker.Cards{}
	combine(c.Cards, poker.NewEmptyCards(), ret, 3)

	bestType := BULL_ZERO

	for _, v := range *(ret) {
		typ := BULL_ZERO

		if calculatePoint(v)%10 == 0 {
			temp := c.Cards.Clone()
			for _, c := range v.Inner {
				temp.RemoveCard(c)
			}
			point := calculatePoint(temp) % 10
			if point == 0 {
				point = 10
			}
			typ = NNType(point)
		}

		if typ > bestType {
			bestType = typ
		}
	}
	if bestType == BULL_ZERO {
		bestPower := poker.NewCard(poker.DIAMOND, 1)
		for _, v := range c.Cards.Inner {
			if CompareCard(v, bestPower) {
				bestPower = v
			}
		}
		c.power = bestPower
	}
	c.typ = bestType
}

func CardPoint(card poker.Card) int {
	rank := card.RankInt()
	if rank > 10 {
		return 10
	}
	return rank
}
func CardsPoint(cards *poker.Cards) int {
	v := 0
	if cards == nil || cards.Size() == 0 {
		return v
	}
	for _, c := range cards.Inner {
		v += CardPoint(c)
	}
	return v
}

func calculatePoint(cards *poker.Cards) int {
	return CardsPoint(cards)
}

// rank compare : k>q>j>10>9>8>7>6>5>4>3>2>a
// suit compare : 黑桃>红桃>梅花>方块
// CompareCard return a > b ?
func CompareCard(a, b poker.Card) bool {
	if a.RankInt() > b.RankInt() {
		return true
	} else if a.RankInt() == b.RankInt() {
		return a.Suit() > b.Suit()
	}
	return false
}

// TODO: re-write with go-style
func combine(raw *poker.Cards, subset *poker.Cards, out *[]*poker.Cards, m int) {
	if raw.Size() < 1 {
		return
	}

	if raw.Size()+subset.Size() < m {
		return
	}

	cards := raw.Clone()
	for i := 0; i < raw.Size(); i++ {
		temp := subset.Clone()
		temp.BrickCard(cards.Inner[0])
		cards.Remove(0)
		if m != temp.Size() {
			combine(cards, temp, out, m)
		} else {
			*out = append(*out, temp)
		}
	}
}
