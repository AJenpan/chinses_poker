package poker

// suit to chinses:
// diamond 方块
// club    梅花
// heart   红桃
// spade   黑桃

// rank to chinses:
// 1-10  : two,three,four,five,six,seven,eight,nine,ten.
// J-A   : Jack,Queen,King,Ace
// joker : 大小王

//CardSuit suit of card
type CardSuit byte

//CardRank Rank value of card
type CardRank byte

//Card a card of poker,memony value 4bits == |--2 bits suit -- 2 bits rank--|
type Card byte

const (
	EmptyCard Card = 0x00

	//Mask
	CARD_SUIT_MASK = byte(0xF0) //1111 0000
	CARD_RANK_MASK = byte(0x0F) //0000 1111

	EMPTY   CardSuit = 0
	DIAMOND CardSuit = 1
	CLUB    CardSuit = 2
	HEART   CardSuit = 3
	SPADE   CardSuit = 4
	JOKER   CardSuit = 5
)

var Suits = []byte("?DCHSJ")
var Ranks = []byte("?A23456789TJQK")

//CreateCard by suit and rank
func CreateCard(suit CardSuit, rank CardRank) Card {
	if !(suit.Valid() && rank.Valid()) {
		return EmptyCard
	}
	raw := byte(suit) << 4
	raw |= byte(rank)

	return Card(raw)
}

func StringToCard(str string) Card {
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
	return CreateCard(suit, rank)
}

//ByteToCard byte to card
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

//Suit return card's suit
func (c Card) Suit() CardSuit {
	return CardSuit(byte(c) >> 4)
}

func (c Card) SuitInt() int {
	return int(c.Suit())
}

//Rank return card's rank
func (c Card) Rank() CardRank {
	return CardRank(byte(c) & CARD_RANK_MASK)
}

func (c Card) RankInt() int {
	return int(c.Rank())
}

func (c Card) Valid() bool {
	return c.Suit().Valid() && c.Rank().Valid()
}

func (c Card) String() string {
	return string([]byte{Ranks[c.Rank()], Suits[c.Suit()]})
}

func (c Card) Byte() byte {
	return byte(c)
}

func (c Card) SetSuit(suit CardSuit) Card {
	return CreateCard(suit, c.Rank())
}

func (c Card) SetRank(rank CardRank) Card {
	return CreateCard(c.Suit(), rank)
}
