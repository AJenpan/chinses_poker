package poker

import (
	"strings"
)

// suit to chinses:
// diamond 方块
// club    梅花
// heart   红桃
// spade   黑桃

// rank to chinses:
// 1-10  : two,three,four,five,six,seven,eight,nine,ten.
// J-A   : Jack,Queen,King,Ace
// joker : 大小王

// CardSuit: suit of card
type CardSuit uint8

// CardRank: Rank value of card
type CardRank uint8

// Card a card of poker, memony value 8bits.
// |--4 bits suit -- 4 bits rank--|
type Card uint8

const (
	EmptyCard Card = 0x00

	//Mask
	CARD_SUIT_MASK = byte(0xF0) //1111 0000
	CARD_RANK_MASK = byte(0x0F) //0000 1111

	//Suit
	EMPTY   CardSuit = 0
	DIAMOND CardSuit = 1
	CLUB    CardSuit = 2
	HEART   CardSuit = 3
	SPADE   CardSuit = 4
	JOKER   CardSuit = 5
)

func (cs CardSuit) Chinese() string {
	switch cs {
	case DIAMOND:
		return "方块"
	case CLUB:
		return "梅花"
	case HEART:
		return "红桃"
	case SPADE:
		return "黑桃"
	case JOKER:
		return "王"
	default:
		return "无效"
	}
}

// ♣♦♥♠
var Suits = []byte("?DCHSJ")
var Ranks = []byte("?A23456789XJQK")

// NewCard by suit and rank
func NewCard(suit CardSuit, rank CardRank) Card {
	if !(suit.Valid() && rank.Valid()) {
		return EmptyCard
	}
	raw := byte(suit) << 4
	raw |= byte(rank)
	return Card(raw)
}

func NewCardByString(str string) Card {
	str = strings.TrimSpace(str)
	if len(str) != 2 {
		return EmptyCard
	}
	rank := CardRank(0)
	for i, r := range Ranks {
		if r == str[0] {
			rank = CardRank(i)
			break
		}
	}
	suit := CardSuit(0)
	for i, s := range Suits {
		if s == str[1] {
			suit = CardSuit(i)
			break
		}
	}
	return NewCard(suit, rank)
}

// ByteToCard byte to card
func ByteToCard(b byte) Card {
	r := Card(b)
	if !r.Valid() {
		return EmptyCard
	}
	return r
}

func (r CardRank) Valid() bool {
	return 0 < r && r < 14
}

func (s CardSuit) Valid() bool {
	return 0 < s && s < 6
}

// Suit return card's suit
func (c Card) Suit() CardSuit {
	return CardSuit(byte(c) >> 4)
}

func (c Card) SuitInt() int {
	return int(c.Suit())
}

// Rank return card's rank
func (c Card) Rank() CardRank {
	return CardRank(byte(c) & CARD_RANK_MASK)
}

func (c Card) RankInt() int {
	return int(c.Rank())
}

func (c Card) Valid() bool {
	if c.Suit() == JOKER {
		return c.Rank() == 1 || c.Rank() == 2
	}
	return c.Suit().Valid() && c.Rank().Valid()
}

func (c Card) String() string {
	return string([]byte{Ranks[c.Rank()], Suits[c.Suit()]})
}

func (c Card) Byte() byte {
	return byte(c)
}

func (c Card) SetSuit(suit CardSuit) Card {
	return NewCard(suit, c.Rank())
}

func (c Card) SetRank(rank CardRank) Card {
	return NewCard(c.Suit(), rank)
}

func (c Card) Chinese() string {
	if !c.Valid() {
		return "空牌"
	}
	s := c.Suit()
	suitStr := s.Chinese()
	rankStr := string(Ranks[c.Rank()])
	if rankStr == "T" {
		rankStr = "10"
	}
	if s == JOKER {
		switch c.Rank() {
		case 1:
			return "小王"
		case 2:
			return "大王"
		default:
			return "空牌"
		}
	}
	return suitStr + string(rankStr)
}
