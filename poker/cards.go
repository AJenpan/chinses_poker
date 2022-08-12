package poker

import (
	"math/rand"
	"sort"
	"time"
)

// deck    整付牌 (pack of cards)
// fan     开扇
// spread  摊牌
// cut     切牌
// shuffle 洗牌

//Cards a list of card
type Cards struct {
	Inner []Card
}

func NewEmptyCards() *Cards {
	return &Cards{
		Inner: []Card{},
	}
}

//DealCard deal a card from deck. 发一张牌
func (d *Cards) DealCard() Card {
	if len(d.Inner) < 1 {
		return EmptyCard
	}
	d.Inner = d.Inner[1:]
	return d.Inner[0]
}

//DealCards deal n cards from deck. 发多张牌
func (d *Cards) DealCards(n int) *Cards {
	ret := NewEmptyCards()
	if n >= len(d.Inner) {
		return ret
	}
	ret.Inner, d.Inner = d.Inner[:n], d.Inner[n:]
	return ret
}

//BrickCard 加入一张牌
func (d *Cards) BrickCard(n Card) {
	d.Inner = append(d.Inner, n)
}

func (d *Cards) BrickDeck(n *Cards) {
	d.Inner = append(d.Inner, n.Inner...)
}

func (d *Cards) Sort(fn func(i, j int) bool) {
	sort.Slice(d.Inner, fn)
}

func (d *Cards) Range(fn func(i int, c Card)) {
	for i, v := range d.Inner {
		fn(i, v)
	}
}

func (d *Cards) Remove(index int) {
	if index >= d.Size() {
		return
	}
	d.Inner = append(d.Inner[:index], d.Inner[index+1:]...)
}

//  remove card from deck. 删除一张牌
func (d *Cards) RemoveCard(c Card) {
	for i, v := range d.Inner {
		if v.Byte() == c.Byte() {
			d.Remove(i)
			return
		}
	}
}

func (d *Cards) Copy() *Cards {
	new := &Cards{Inner: make([]Card, len(d.Inner))}
	copy(new.Inner, d.Inner)
	return new
}

//Shuffle 洗牌
func (d *Cards) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(d.Size(), func(i, j int) {
		d.Inner[i], d.Inner[j] = d.Inner[j], d.Inner[i]
	})
}

func (d *Cards) Bytes() []byte {
	ret := make([]byte, d.Size())
	for i, v := range d.Inner {
		ret[i] = v.Byte()
	}
	return ret
}

func (d *Cards) String() string {
	ret := ""
	if d.Size() < 1 {
		return ret
	}
	for _, card := range d.Inner {
		ret += card.String()
		ret += " "
	}
	return ret[0 : len(ret)-1]
}

func (d *Cards) Size() int {
	return len(d.Inner)
}
