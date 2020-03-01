package Niuniu

import (
	"fmt"

	"github.com/Ajenpan/chinese_poker_go/poker"
)

const maxHandCardCout = 5

type NNType byte

const (
	NO_POINT        NNType = 0  //没牛
	BULL_ONE        NNType = 1  //牛一
	BULL_TWO        NNType = 2  //牛二
	BULL_THREE      NNType = 3  //牛三
	BULL_FOUR       NNType = 4  //牛四
	BULL_FIVE       NNType = 5  //牛五
	BULL_SIX        NNType = 6  //牛六
	BULL_SEVEN      NNType = 7  //牛七
	BULL_EIGHT      NNType = 8  //牛八
	BULL_NINE       NNType = 9  //牛九
	BULL_BULL       NNType = 10 //牛牛
	BULL_FIVESMALL  NNType = 11 //五小牛
	BULL_FIVEFLOWER NNType = 12 //五花牛
	BULL_FOURBOMB   NNType = 13 //炸弹牛
	ERROR_TYPE      NNType = 255
)

// type NNCards poker.Cards

type NNHandCards struct {
	poker.Cards
	typ   NNType
	power poker.Card
}

type NNDeck struct {
	poker.Cards
}

func NewNNDeck() *NNDeck {
	return &NNDeck{
		Cards: poker.CreateDeckWithoutJoker(),
	}
}

func (nn *NNDeck) DealHandCards() *NNHandCards {
	cards := nn.DealCards(maxHandCardCout)
	ret := &NNHandCards{}
	ret.Cards = cards
	ret.calculate()
	return ret
}

//TODO: re-write with go-style
func combine(raw poker.Cards, subset poker.Cards, out *[]poker.Cards, m int) {
	if raw.Size() < 1 {
		return
	}

	if raw.Size()+subset.Size() < m {
		return
	}

	cards := raw.Copy()
	for i := 0; i < raw.Size(); i++ {
		temp := subset.Copy()
		temp.BrickCard(cards[0])
		cards.Remove(0)
		if m != temp.Size() {
			combine(cards, temp, out, m)
		} else {
			*out = append(*out, temp)
		}
	}
	return
}

func IsBull(cards *poker.Cards) bool {
	if cards.Size() != 3 {
		return false
	}
	return calculatePoint(cards)%10 == 0
}

func calculatePoint(cards *poker.Cards) int {
	v := 0
	for _, c := range *cards {
		if c.RankInt() > 10 {
			v += 10
		} else {
			v += c.RankInt()
		}
	}
	return v
}

func (card *NNHandCards) isFourBomb() bool {
	count := make(map[int]int)
	for _, v := range card.Cards {
		if _, ok := count[v.RankInt()]; ok {
			count[v.RankInt()]++
		} else {
			count[v.RankInt()] = 1
		}
	}
	for r, v := range count {
		if v == 4 {
			card.typ = BULL_FOURBOMB
			card.power = poker.CreateCard(poker.SPADE, poker.CardRank(r))
			return true
		}
	}
	return false
}

func (card *NNHandCards) isFiveSmall() bool {
	point := calculatePoint(&card.Cards)
	if point > 10 {
		return false
	}
	for _, v := range card.Cards {
		if v.RankInt() > 5 {
			return false
		}
	}
	card.typ = BULL_FIVESMALL
	return true
}

func (card *NNHandCards) isFiveFlower() bool {
	if calculatePoint(&card.Cards) != 50 {
		return false
	}
	for _, v := range card.Cards {
		if v.RankInt() <= 10 {
			return false
		}
	}
	card.typ = BULL_FIVEFLOWER
	return true
}

func calculate(card *NNHandCards, subCards *poker.Cards) NNType {
	if subCards.Size() != 3 {
		return ERROR_TYPE
	}

	if IsBull(subCards) {

	} else {
		return NO_POINT
	}

	return ERROR_TYPE
}

//数字比较： k>q>j>10>9>8>7>6>5>4>3>2>a。
//花色比较：黑桃>红桃>梅花>方块。
func compareCard(a, b poker.Card) bool {
	if a.RankInt() > b.RankInt() {
		return true
	} else if a.RankInt() == b.RankInt() {
		return a.Suit() > b.Suit()
	}
	return false
}

func (nn *NNHandCards) calculate() {

	if nn.isFourBomb() {
		return
	}
	if nn.isFiveFlower() {
		return
	}
	if nn.isFiveSmall() {
		return
	}

	fmt.Println(nn.Cards)
	ret := &[]poker.Cards{}
	combine(nn.Cards, poker.CreateEmpty(), ret, 3)

	bestType := NO_POINT

	for _, v := range *(ret) {
		typ := NO_POINT

		if calculatePoint(&v)%10 == 0 {
			temp := nn.Cards.Copy()
			for _, c := range v {
				temp.RemoveCard(c)
			}
			point := calculatePoint(&temp) % 10
			if point == 0 {
				point = 10
			}
			typ = NNType(point)
		}

		if typ > bestType {
			bestType = typ
		}
	}
	if bestType == NO_POINT {
		bestPower := poker.CreateCard(poker.DIAMOND, 1)
		for _, v := range nn.Cards {
			if compareCard(v, bestPower) {
				bestPower = v
			}
		}
		nn.power = bestPower
	}
	nn.typ = bestType
}
