package guandan

import "github.com/ajenpan/poker_algorithm/poker"

type PromptResult struct {
	DeckPower
	Cards *GDCards
}

type Prompt struct {
	// Target      GDCards
	TargetPower DeckPower
	HandCards   GDCards
	Wildcard    GDCard

	Results []*PromptResult

	wildCards   GDCards
	normalCards GDCards

	suitCnt       map[poker.CardSuit]*GDCards
	rankCnt       map[poker.CardRank]*GDCards
	logicValueCnt map[uint8]*GDCards
}

func (p *Prompt) Search() {
	p.suitCnt = make(map[poker.CardSuit]*GDCards)
	p.rankCnt = make(map[poker.CardRank]*GDCards)
	p.logicValueCnt = make(map[uint8]*GDCards)

	p.HandCards.SortByByte()

	for _, card := range p.HandCards.Inner {
		if card == p.Wildcard {
			p.wildCards.Push(card)
		} else {
			p.normalCards.Push(card)

			rankCnt := p.rankCnt[card.Rank()]
			if rankCnt == nil {
				rankCnt = &GDCards{}
				p.rankCnt[card.Rank()] = rankCnt
			}
			suitCnt := p.suitCnt[card.Suit()]
			if suitCnt == nil {
				suitCnt = &GDCards{}
				p.suitCnt[card.Suit()] = suitCnt
			}
			logicValue := CardLogicValue(p.Wildcard, card)
			logicValueCnt := p.logicValueCnt[logicValue]
			if logicValueCnt == nil {
				logicValueCnt = &GDCards{}
				p.logicValueCnt[logicValue] = logicValueCnt
			}

			rankCnt.Push(card)
			suitCnt.Push(card)
			logicValueCnt.Push(card)
		}
	}
}

func (p *Prompt) searchSingle() {
	if p.TargetPower.DeckType != DeckSingle {
		return
	}

	for v, cards := range p.logicValueCnt {
		if v > p.TargetPower.DeckValue {
			p.Results = append(p.Results, &PromptResult{
				DeckPower: DeckPower{
					DeckType:  DeckSingle,
					DeckValue: v,
				},
				Cards: poker.NewCards(cards.Inner[:1]),
			})
		}
	}
}

func (p *Prompt) searchPair() {
	if p.TargetPower.DeckType != DeckPair {
		return
	}

	for v, cards := range p.logicValueCnt {
		if v >= p.TargetPower.DeckValue {
			continue
		}
		// 不允许拆炸弹
		if cards.Size() >= 4 {
			continue
		}
		if cards.Size()+p.wildCards.Size() < 2 {
			continue
		}

		needWildCnt := 0
		needCardCnt := 2

		if cards.Size() < 2 {
			needWildCnt = 2 - cards.Size()
			needCardCnt = cards.Size()
		}

		result := &PromptResult{
			DeckPower: DeckPower{
				DeckType:  DeckPair,
				DeckValue: v,
			},
			Cards: poker.NewCards(cards.Inner[:needCardCnt]),
		}
		if needWildCnt > 0 {
			result.Cards.Append(poker.NewCards(p.wildCards.Inner[:needWildCnt]))
		}
		p.Results = append(p.Results, result)
	}
}

func (p *Prompt) searchThree() {
	if p.TargetPower.DeckType != DeckThree {
		return
	}
}

func (p *Prompt) searchThreeWithTwo() {

}

func (p *Prompt) searchStraight() {

}

func (p *Prompt) searchStraightFlush() {

}

func (p *Prompt) searchStraightPair() {

}

func (p *Prompt) searchStraightThree() {

}

func (p *Prompt) searchBomb4() {

}

func (p *Prompt) searchBomb5() {

}

func (p *Prompt) searchBomb6() {

}

func (p *Prompt) searchBomb7() {

}

func (p *Prompt) searchBomb8() {

}

func (p *Prompt) searchBomb9() {

}

func (p *Prompt) searchBomb10() {

}

func (p *Prompt) searchBombJoker() {

}

func (p *Prompt) searchWindflow() {

}
