package guandan

import (
	"github.com/ajenpan/poker_algorithm/poker"
)

type GDCards = poker.Cards

type GDDeck = poker.Cards

func NewDeck() *GDDeck {
	deck := poker.NewDeck()
	deck.Append(poker.NewDeck())
	return deck
}

type DeckType int

const (
	DeckPass          DeckType = 0  // 过
	DeckSingle        DeckType = 1  // 单张
	DeckPair          DeckType = 2  // 对子
	DeckThree         DeckType = 3  // 3张
	DeckThreeWithTwo  DeckType = 4  // 3带2
	DeckStraight      DeckType = 5  // 顺子
	DeckStraightFlush DeckType = 6  // 同花顺
	DeckStraightPair  DeckType = 7  // 连对
	DeckStraightThree DeckType = 8  // 钢板 (2个连续的三张牌)
	DeckBomb4         DeckType = 9  // 4炸
	DeckBomb5         DeckType = 10 // 5炸
	DeckBomb6         DeckType = 11 // 6炸
	DeckBomb7         DeckType = 12 // 7炸
	DeckBomb8         DeckType = 13 // 8炸
	DeckBomb9         DeckType = 14 // 9炸
	DeckBomb10        DeckType = 15 // 10炸
	DeckBombJoker     DeckType = 16 // 王炸
	DeckWindflow      DeckType = 17 // 接风
)

type DeckPower struct {
	DeckType  DeckType
	DeckValue int
	Err       error
}

// result 1: dp > other, 0: dp == other, -1: dp < other, -2: cannot compare, -3: error
func (dp *DeckPower) Compare(other *DeckPower) int {
	if dp.DeckType == DeckWindflow || other.DeckType == DeckWindflow {
		return -2
	}
	if other.DeckType == DeckPass {
		return 1
	}
	if dp.DeckType == DeckPass {
		return -1
	}

	// same type
	if dp.DeckType == other.DeckType {
		if dp.DeckValue > other.DeckValue {
			return 1
		} else if dp.DeckValue < other.DeckValue {
			return -1
		}
		return 0
	}

	dpIsBomb := false
	if dp.DeckType >= DeckBomb4 && dp.DeckType <= DeckBombJoker {
		dpIsBomb = true
	}
	otherIsBomb := false
	if other.DeckType >= DeckBomb4 && other.DeckType <= DeckBombJoker {
		otherIsBomb = true
	}

	if !dpIsBomb && !otherIsBomb {
		// 牌型不同且非炸弹牌型无法比较
		return -2
	}

	// dp是炸弹，other不是炸弹
	if dpIsBomb && !otherIsBomb {
		return 1
	}

	// dp不是炸弹，other是炸弹
	if !dpIsBomb && otherIsBomb {
		return -1
	}

	if dp.DeckType > other.DeckType {
		return 1
	}

	if dp.DeckType < other.DeckType {
		return -1
	}
	return -2
}

func CardSingleCardPower(wildcard poker.CardRank, card poker.Card) int {
	if card.Rank() == wildcard {
		return 20
	}
	if card.Rank() == poker.RANK_A {
		return int(poker.RANK_K) + 1
	}
	if card.Suit() == poker.JOKER {
		return 30 + int(card.Rank())
	}
	return int(card.Rank())
}

func getMinMaxRank(ranks map[poker.CardRank]int) (poker.CardRank, poker.CardRank) {
	var minRank, maxRank poker.CardRank = poker.RANK_K + 10, 0
	for k := range ranks {
		if k < minRank {
			minRank = k
		}
		if k > maxRank {
			maxRank = k
		}
	}
	return minRank, maxRank
}

type info struct {
	wildcard poker.CardRank
	cards    *poker.Cards

	normalCards *poker.Cards
	wildCards   *poker.Cards

	suitCnt map[poker.CardSuit]int
	rankCnt map[poker.CardRank]int
}

func getDeckPowerWithNoWildcard(in *info) *DeckPower {
	errpower := &DeckPower{DeckPass, 0, nil}

	switch in.cards.Size() {
	case 2:
		if len(in.rankCnt) == 1 {
			return &DeckPower{DeckPair, CardSingleCardPower(in.wildcard, in.normalCards.Front()), nil}
		}
	case 3:
		if len(in.rankCnt) == 1 {
			return &DeckPower{DeckThree, CardSingleCardPower(in.wildcard, in.normalCards.Front()), nil}
		}
	case 4: // 4炸, 天王炸
		if in.suitCnt[poker.JOKER] == 4 {
			return &DeckPower{DeckBombJoker, 0, nil}
		}
		if len(in.rankCnt) == 1 {
			return &DeckPower{DeckBomb4, CardSingleCardPower(in.wildcard, in.normalCards.Front()), nil}
		}
	case 5: // 顺子, 同花顺, 5炸, 3带2
		if len(in.rankCnt) == 1 { // 5炸
			return &DeckPower{DeckBomb5, CardSingleCardPower(in.wildcard, in.normalCards.Front()), nil}
		} else if len(in.rankCnt) == 2 { // 3带2
			var rankA, rankB poker.CardRank = 0, 0
			for k, v := range in.rankCnt {
				if v == 3 {
					rankA = k
				} else if v == 2 {
					rankB = k
				}
			}
			if rankA != 0 && rankB != 0 {
				return &DeckPower{DeckThreeWithTwo, CardSingleCardPower(in.wildcard, poker.NewCard(1, poker.CardRank(rankA))), nil}
			}
		} else if len(in.rankCnt) == 5 { // 顺子
			in.cards.SortByRank()
			startRank := in.cards.Inner[0].Rank()
			endRank := in.cards.Inner[4].Rank()
			power := 0
			if endRank == poker.RANK_K && startRank == poker.RANK_A {
				for i := 1; i < 4; i++ {
					if in.cards.Inner[i].Rank()+5-poker.CardRank(i) != endRank {
						return errpower
					}
				}
				power = CardSingleCardPower(in.wildcard, poker.NewCard(1, poker.RANK_A))
			} else {
				for i := 1; i < 5; i++ {
					if startRank+poker.CardRank(i) != in.cards.Inner[i].Rank() {
						return errpower
					}
				}
				power = CardSingleCardPower(in.wildcard, poker.NewCard(1, endRank))
			}
			if len(in.suitCnt) == 1 {
				return &DeckPower{DeckStraightFlush, power, nil}
			} else {
				return &DeckPower{DeckStraight, power, nil}
			}
		}
	case 6: // 连对, 钢板, 6炸
		if len(in.rankCnt) == 1 { // 6炸
			return &DeckPower{DeckBomb6, CardSingleCardPower(in.wildcard, in.cards.Inner[0]), nil}
		}
		var minRank, MaxRank = poker.RANK_K + 1, poker.CardRank(0)
		for k, v := range in.rankCnt {
			if v != 3 {
				return errpower
			}
			if k < minRank {
				minRank = k
			}
			if k > MaxRank {
				MaxRank = k
			}
		}

		if len(in.rankCnt) == 2 { // 钢板
			if minRank == poker.RANK_A {
				if MaxRank == poker.RANK_2 { // AAA222
					return &DeckPower{DeckStraightThree, int(poker.RANK_2), nil}
				} else if MaxRank == poker.RANK_K { //KKKAAA
					return &DeckPower{DeckStraightThree, int(poker.RANK_K) + 1, nil}
				}
			} else {
				if minRank+1 == MaxRank {
					return &DeckPower{DeckStraightThree, int(MaxRank), nil}
				}
			}
		} else if len(in.rankCnt) == 3 { // 连对
			if minRank == poker.RANK_3 {
				if MaxRank == poker.RANK_2 { // AA2233
					return &DeckPower{DeckStraightThree, int(poker.RANK_3), nil}
				} else if MaxRank == poker.RANK_K { //QQKKAA
					return &DeckPower{DeckStraightThree, int(poker.RANK_K) + 1, nil}
				}
			} else {
				if minRank+2 == MaxRank {
					return &DeckPower{DeckStraightPair, int(MaxRank), nil}
				}
			}
		}
	case 7:
		if len(in.rankCnt) == 1 {
			return &DeckPower{DeckBomb7, CardSingleCardPower(in.wildcard, in.cards.Inner[0]), nil}
		}
	case 8:
		if len(in.rankCnt) == 1 {
			return &DeckPower{DeckBomb8, CardSingleCardPower(in.wildcard, in.cards.Inner[0]), nil}
		}
	}
	return errpower
}

func getDeckPowerWith1Wildcard(in *info) *DeckPower {
	errpower := &DeckPower{DeckPass, 0, nil}

	switch in.cards.Size() {
	case 2:
		if len(in.rankCnt) == 1 {
			return &DeckPower{DeckPair, CardSingleCardPower(in.wildcard, in.cards.Inner[0]), nil}
		}
	case 3:
		if len(in.rankCnt) == 1 {
			return &DeckPower{DeckThree, CardSingleCardPower(in.wildcard, in.cards.Inner[0]), nil}
		}
	case 4: // 4炸, 天王炸
		if in.suitCnt[poker.JOKER] == 4 {
			return &DeckPower{DeckBombJoker, 0, nil}
		}
		if len(in.rankCnt) == 1 {
			return &DeckPower{DeckBomb4, CardSingleCardPower(in.wildcard, in.cards.Inner[0]), nil}
		}
	case 5: // 顺子, 同花顺, 5炸, 3带2
		if len(in.rankCnt) == 1 { // 5炸
			return &DeckPower{DeckBomb5, CardSingleCardPower(in.wildcard, in.cards.Inner[0]), nil}
		} else if len(in.rankCnt) == 2 { // 3带2
			var rankA, rankB poker.CardRank = 0, 0
			for k, v := range in.rankCnt {
				if v == 3 {
					rankA = k
				} else if v == 2 {
					rankB = k
				}
			}
			if rankA != 0 && rankB != 0 {
				return &DeckPower{DeckThreeWithTwo, CardSingleCardPower(in.wildcard, poker.NewCard(1, poker.CardRank(rankA))), nil}
			}
		} else if len(in.rankCnt) == 5 { // 顺子
			in.cards.SortByRank()
			startRank := in.cards.Inner[0].Rank()
			endRank := in.cards.Inner[4].Rank()
			power := 0
			if endRank == poker.RANK_K && startRank == poker.RANK_A {
				for i := 1; i < 4; i++ {
					if in.cards.Inner[i].Rank()+5-poker.CardRank(i) != endRank {
						return errpower
					}
				}
				power = CardSingleCardPower(in.wildcard, poker.NewCard(1, poker.RANK_A))
			} else {
				for i := 1; i < 5; i++ {
					if startRank+poker.CardRank(i) != in.cards.Inner[i].Rank() {
						return errpower
					}
				}
				power = CardSingleCardPower(in.wildcard, poker.NewCard(1, endRank))
			}
			if len(in.suitCnt) == 1 {
				return &DeckPower{DeckStraightFlush, power, nil}
			} else {
				return &DeckPower{DeckStraight, power, nil}
			}
		}
	case 6: // 连对, 钢板, 6炸
		if len(in.rankCnt) == 1 { // 6炸
			return &DeckPower{DeckBomb6, CardSingleCardPower(in.wildcard, in.cards.Inner[0]), nil}
		}
		var minRank, MaxRank = poker.RANK_K + 1, poker.CardRank(0)
		for k, v := range in.rankCnt {
			if v != 3 {
				return errpower
			}
			if k < minRank {
				minRank = k
			}
			if k > MaxRank {
				MaxRank = k
			}
		}

		if len(in.rankCnt) == 2 { // 钢板
			if minRank == poker.RANK_A {
				if MaxRank == poker.RANK_2 { // AAA222
					return &DeckPower{DeckStraightThree, int(poker.RANK_2), nil}
				} else if MaxRank == poker.RANK_K { //KKKAAA
					return &DeckPower{DeckStraightThree, int(poker.RANK_K) + 1, nil}
				}
			} else {
				if minRank+1 == MaxRank {
					return &DeckPower{DeckStraightThree, int(MaxRank), nil}
				}
			}
		} else if len(in.rankCnt) == 3 { // 连对
			if minRank == poker.RANK_3 {
				if MaxRank == poker.RANK_2 { // AA2233
					return &DeckPower{DeckStraightThree, int(poker.RANK_3), nil}
				} else if MaxRank == poker.RANK_K { //QQKKAA
					return &DeckPower{DeckStraightThree, int(poker.RANK_K) + 1, nil}
				}
			} else {
				if minRank+2 == MaxRank {
					return &DeckPower{DeckStraightPair, int(MaxRank), nil}
				}
			}
		}
	case 7:
		if len(in.rankCnt) == 1 {
			return &DeckPower{DeckBomb7, CardSingleCardPower(in.wildcard, in.cards.Inner[0]), nil}
		}
	case 8:
		if len(in.rankCnt) == 1 {
			return &DeckPower{DeckBomb8, CardSingleCardPower(in.wildcard, in.cards.Inner[0]), nil}
		}
	}
	return errpower
}

func getDeckPowerWith2Wildcard(i *info) *DeckPower {
	return &DeckPower{DeckPass, 0, nil}
}

func GetDeckPower(wildcard poker.CardRank, cards *poker.Cards) *DeckPower {
	errpower := &DeckPower{DeckPass, 0, nil}
	if cards.Size() == 0 {
		return errpower
	} else if cards.Size() == 1 {
		return &DeckPower{DeckSingle, CardSingleCardPower(wildcard, cards.Inner[0]), nil}
	}

	info := &info{
		wildcard:    wildcard,
		cards:       cards,
		normalCards: poker.NewEmptyCards(),
		wildCards:   poker.NewEmptyCards(),
		suitCnt:     make(map[poker.CardSuit]int),
		rankCnt:     make(map[poker.CardRank]int),
	}

	for _, card := range cards.Inner {
		if card.Rank() == wildcard {
			info.wildCards.Push(card)
		} else {
			info.normalCards.Push(card)
			info.rankCnt[card.Rank()]++
			info.suitCnt[card.Suit()]++
		}
	}

	switch info.wildCards.Size() {
	case 0:
		return getDeckPowerWithNoWildcard(info)
	case 1:
		return getDeckPowerWith1Wildcard(info)
	case 2:
		return getDeckPowerWith2Wildcard(info)
	}
	return errpower
}
