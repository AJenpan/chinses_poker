package poker

import (
	"math/rand"
	"sort"
	"strings"
	"time"
)

// deck    整付牌 (pack of cards)
// fan     开扇
// spread  摊牌
// cut     切牌
// shuffle 洗牌

func NewEmptyCards() *Cards {
	return &Cards{Inner: []Card{}}
}

func NewCards(cards []Card) *Cards {
	return &Cards{Inner: cards}
}

// Cards a list of card
type Cards struct {
	Inner []Card
}

// DealCard deal a card from deck. 发一张牌
func (d *Cards) DealCard() Card {
	return d.PopFront()
}

func (d *Cards) PopFront() Card {
	if d.Size() < 1 {
		return EmptyCard
	}
	ret := d.Inner[0]
	d.Inner = d.Inner[1:]
	return ret
}

func (d *Cards) PopBack() Card {
	if d.Size() < 1 {
		return EmptyCard
	}
	ret := d.Inner[d.Size()-1]
	d.Inner = d.Inner[:d.Size()-1]
	return ret
}

// DealCards deal n cards from deck. 发多张牌
func (d *Cards) DealCards(n int) *Cards {
	ret := NewEmptyCards()
	if n > len(d.Inner) {
		return ret
	}
	ret.Inner, d.Inner = d.Inner[:n], d.Inner[n:]
	return ret
}

func (d *Cards) Front() Card {
	if d.Size() < 1 {
		return EmptyCard
	}
	return d.Inner[0]
}

func (d *Cards) Back() Card {
	if d.Size() < 1 {
		return EmptyCard
	}
	return d.Inner[d.Size()-1]
}

func (d *Cards) Push(n Card) {
	d.Inner = append(d.Inner, n)
}

func (d *Cards) Append(n *Cards) {
	d.Inner = append(d.Inner, n.Inner...)
}

func (d *Cards) BrickCards(n *Cards) {
	d.Inner = append(d.Inner, n.Inner...)
}

func (d *Cards) BrickCard(n Card) {
	d.Inner = append(d.Inner, n)
}

// order is DA CA HA SA D2
func (d *Cards) SortByRank() {
	sort.Sort(ByRank(d.Inner))
}

// order is DA D2 D3 D4 D5
func (d *Cards) SortBySuit() {
	sort.Sort(BySuit(d.Inner))
}

// order is DA D2 D3 D4 D5
func (d *Cards) SortByByte() {
	sort.Slice(d.Inner, func(i, j int) bool {
		return byte(d.Inner[i]) < byte(d.Inner[j])
	})
}

func (d *Cards) Sort(fn func(Card, Card) bool) {
	sort.Slice(d.Inner, func(i, j int) bool {
		return fn(d.Inner[i], d.Inner[j])
	})

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

func (d *Cards) Clear() {
	d.Inner = d.Inner[:0]
}

func (d *Cards) Get(index int) Card {
	if index >= d.Size() {
		return EmptyCard
	}
	return d.Inner[index]
}

// Remove card from deck. 删除一张牌
func (d *Cards) RemoveCard(c Card) bool {
	for i, v := range d.Inner {
		if v.Byte() == c.Byte() {
			d.Remove(i)
			return true
		}
	}
	return false
}

// RemoveCards remove cards from deck
// will change the cards order
func (d *Cards) RemoveCards(cs *Cards) bool {
	curr := 0
	for _, c := range cs.Inner {
		if !d.RemoveCard(c) {
			break
		}
		curr += 1
	}

	ok := curr == cs.Size()
	if !ok {
		for i := 0; i < curr; i++ {
			d.Push(cs.Inner[i])
		}
	}
	return ok
}

// 尽可能删除
func (d *Cards) Sub(s *Cards) *Cards {
	dm := d.ToMap()
	for _, v := range s.Inner {
		c, has := dm[v]
		if !has {
			continue
		}
		if c > 0 {
			c = c - 1
			dm[v] = c
		}
		if c == 0 {
			delete(dm, v)
		}
	}

	new := make([]Card, 0, len(dm))
	for k, v := range dm {
		for i := 0; i < v; i++ {
			new = append(new, k)
		}
	}
	d.Inner = new
	return d
}

func (d *Cards) Clone() *Cards {
	new := &Cards{Inner: make([]Card, len(d.Inner))}
	copy(new.Inner, d.Inner)
	return new
}

func (d *Cards) Contain(c Card) bool {
	for _, v := range d.Inner {
		if v.Byte() == c.Byte() {
			return true
		}
	}
	return false
}

// Shuffle 洗牌
func (d *Cards) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(d.Size(), func(i, j int) {
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
	return ret[:len(ret)-1]
}

func (d *Cards) Chinese() string {
	if d.IsEmpty() {
		return ""
	}
	sz := make([]string, d.Size())
	for i, card := range d.Inner {
		sz[i] = card.Chinese()
	}
	return strings.Join(sz, " ")
}

func (d *Cards) Size() int {
	return len(d.Inner)
}

func (d *Cards) IsEmpty() bool {
	return len(d.Inner) == 0
}

// map with card as key and count as value
func (d *Cards) ToMap() map[Card]int {
	ret := make(map[Card]int, d.Size())
	for _, v := range d.Inner {
		ret[v]++
	}
	return ret
}
