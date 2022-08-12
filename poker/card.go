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

	// mask
	CARD_SUIT_MASK = 0xF0 //1111 0000
	CARD_RANK_MASK = 0x0F //0000 1111

	UNKNOW  CardSuit = 0
	DIAMOND CardSuit = 1
	CLUB    CardSuit = 2
	HEART   CardSuit = 3
	SPADE   CardSuit = 4
	JOKER   CardSuit = 5
)

var suits = []byte("?DCHSJ")
var ranks = []byte("?A23456789TJQK")

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
	suit := CardSuit(0)
	for i, s := range suits {
		if s == str[0] {
			suit = CardSuit(i)
			break
		}
	}
	rank := CardRank(0)
	for i, r := range ranks {
		if r == str[1] {
			rank = CardRank(i)
			break
		}
	}
	if !(suit.Valid() && rank.Valid()) {
		return EmptyCard
	}
	return CreateCard(suit, rank)
}

//ByteToCard byte to card
func ByteToCard(b byte) Card {
	rank := CardRank(b | CARD_RANK_MASK)
	if !rank.Valid() {
		return EmptyCard
	}
	suit := CardSuit(b | CARD_SUIT_MASK)
	if !suit.Valid() {
		return EmptyCard
	}
	return CreateCard(suit, rank)
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

//Rank return card's rank
func (c Card) Rank() CardRank {
	return CardRank(c & CARD_RANK_MASK)
}

func (c Card) RankInt() int {
	return int(c.Rank())
}

func (c Card) Valid() bool {
	return c.Suit().Valid() && c.Rank().Valid()
}

func (c Card) String() string {
	return string([]byte{suits[c.Suit()], ranks[c.Rank()]})
}

func (c Card) Byte() byte {
	return byte(c)
}
