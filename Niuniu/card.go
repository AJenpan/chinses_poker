package Niuniu

import (
	"fmt"

	"github.com/Ajenpan/chinese_poker_go/poker"
)

const maxHandCardCout = 5

type NNHandCards struct {
	poker.Cards
}

type NNDeck struct {
	poker.Cards
}

func NewNNDeck() *NNDeck {
	return &NNDeck{
		Cards: poker.CreateDeckWithoutJoker(),
	}
}

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
	BULL_FOURBOMB   NNType = 11 //炸弹牛
)

//func (nn *NNDeck) DealHandCards() NNHandCards {
// return NNHandCards(nn.DealCards(maxHandCardCout))
//}

func (nn *NNHandCards) BestType() {
	b := poker.Card(0)
	a := poker.Card(0)

	fmt.Println(a, b)
	if a > b {

	}
}
