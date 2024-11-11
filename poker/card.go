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
// 1-10  : one,two,three,four,five,six,seven,eight,nine,ten.
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
	RANK_A CardRank = 0x01 + iota
	RANK_2
	RANK_3
	RANK_4
	RANK_5
	RANK_6
	RANK_7
	RANK_8
	RANK_9
	RANK_10
	RANK_J
	RANK_Q
	RANK_K
)

const (
	DIAMOND_A Card = 0x11 + iota
	DIAMOND_2
	DIAMOND_3
	DIAMOND_4
	DIAMOND_5
	DIAMOND_6
	DIAMOND_7
	DIAMOND_8
	DIAMOND_9
	DIAMOND_10
	DIAMOND_J
	DIAMOND_Q
	DIAMOND_K

	CLUB_A Card = 0x21 + iota
	CLUB_2
	CLUB_3
	CLUB_4
	CLUB_5
	CLUB_6
	CLUB_7
	CLUB_8
	CLUB_9
	CLUB_10
	CLUB_J
	CLUB_Q
	CLUB_K

	HEART_A Card = 0x31 + iota
	HEART_2
	HEART_3
	HEART_4
	HEART_5
	HEART_6
	HEART_7
	HEART_8
	HEART_9
	HEART_10
	HEART_J
	HEART_Q
	HEART_K

	SPADE_A Card = 0x41 + iota
	SPADE_2
	SPADE_3
	SPADE_4
	SPADE_5
	SPADE_6
	SPADE_7
	SPADE_8
	SPADE_9
	SPADE_10
	SPADE_J
	SPADE_Q
	SPADE_K

	JOKER_BLACK Card = 0x51 + iota
	JOKER_RED
)

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
var Ranks = []byte("?A234567890JQK")

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
	suit := CardSuit(0)
	for i, s := range Suits {
		if s == str[0] {
			suit = CardSuit(i)
			break
		}
	}
	rank := CardRank(0)
	for i, r := range Ranks {
		if r == str[1] {
			rank = CardRank(i)
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

	suitStr := s.Chinese()
	rankStr := string(Ranks[c.Rank()])
	return suitStr + string(rankStr)
}
