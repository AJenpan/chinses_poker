package poker

import (
	"math/rand"
	"strings"
	"time"
)

/*
pack (of cards), deck 整付牌
fan 开扇
spread 摊牌
cut 切牌
shuffle 洗牌
*/

//Cards a list of card
type Cards []Card

func CreateEmpty() Cards {
	return []Card{}
}

//CreateDeck create a pack of cards
func CreateDeck() Cards {
	deck := Cards{}
	for i := CardRank(1); i <= 13; i++ {
		deck = append(deck, CreateCard(DIAMOND, i))
		deck = append(deck, CreateCard(CLUB, i))
		deck = append(deck, CreateCard(HEART, i))
		deck = append(deck, CreateCard(SPADE, i))
	}
	deck = append(deck, CreateCard(JOKER, 1)) //小王
	deck = append(deck, CreateCard(JOKER, 2)) //大王
	return deck
}

//CreateDeckWithoutJoker without 2 jokers
func CreateDeckWithoutJoker() Cards {
	deck := CreateDeck()
	return deck[:len(deck)-2]
}

func StringToCards(str string) Cards {
	cs := strings.Split(str, " ")
	cards := Cards{}
	for _, v := range cs {
		card := StringToCard(v)
		if !card.Valid() {
			return Cards{}
		}
		cards.BrickCard(card)
	}
	return cards
}

func BytesToCards(raw []byte) Cards {
	cards := Cards{}
	for _, v := range raw {
		card := ByteToCard(v)
		if !card.Valid() {
			return Cards{}
		}
		cards.BrickCard(card)
	}
	return cards
}

//DealCard 发1张牌
func (d *Cards) DealCard() (ret Card) {
	if len(*d) < 1 {
		ret = ErrorCard
		return
	}
	ret, *d = (*d)[0], (*d)[1:]
	return
}

//DealCards 发牌
func (d *Cards) DealCards(n int) (ret Cards) {
	if n >= len(*d) {
		ret = *d
		*d = Cards{}
		return
	}
	(ret), *d = (*d)[:n], (*d)[n:]
	return
}

//BrickCard
func (d *Cards) BrickCard(n Card) {
	*d = append(*d, n)
}

func (d *Cards) BrickDeck(n Cards) {
	*d = append(*d, n...)
}

func (d *Cards) Remove(index int) {
	if index >= d.Size() {
		return
	}
	*d = append((*d)[:index], (*d)[index+1:]...)
}

func (d *Cards) RemoveCard(c Card) {
	for i, v := range *d {
		if v.Byte() == c.Byte() {
			d.Remove(i)
			return
		}
	}
}

func (d Cards) Copy() Cards {
	a := Cards{}
	return append(a, d...)
}

//Shuffle 洗牌
func (d *Cards) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})
}

func (d Cards) Bytes() []byte {
	ret := make([]byte, len(d))
	for i, v := range d {
		ret[i] = v.Byte()
	}
	return ret
}

func (d Cards) String() string {
	ret := ""
	if d.Size() < 1 {
		return ret
	}
	for _, card := range d {
		ret += card.String()
		ret += " "
	}
	return ret[0 : len(ret)-1]
}

func (d *Cards) Size() int {
	return len(*d)
}

func (d Cards) Len() int {
	return len(d)
}
