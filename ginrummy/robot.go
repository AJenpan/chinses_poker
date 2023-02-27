package ginrummy

import (
	"sort"

	"github.com/ajenpan/poker_algorithm/poker"
)

type HoldpowerRes struct {
	Card       poker.Card
	Sum        float32
	BeMeldRate float32
	ScoreRate  float32
}

func (r *CardRate) HoldpowerRateByOrder(cards *poker.Cards) []*HoldpowerRes {
	ret := []*HoldpowerRes{}
	if cards == nil || cards.Size() == 0 {
		return ret
	}
	scoreSum := float32(CardsPoint(cards))

	for _, c := range cards.Inner {
		point := float32(CardPoint(c))
		pointRate := (100.0 - point) / 100.0 // 分数权重
		proportion := 1.0 - point/scoreSum   // 占比

		temp := &HoldpowerRes{
			Card:       c,
			BeMeldRate: r.BeMeldRate(c) * 10,
			ScoreRate:  pointRate * proportion * 10,
		}

		//TODO: what is the better way ?
		temp.Sum = temp.BeMeldRate * temp.ScoreRate
		ret = append(ret, temp)
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Sum < ret[j].Sum
	})
	return ret
}

type Robot struct {
	DrawCardRate *CardRate
}
