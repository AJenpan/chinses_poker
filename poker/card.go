package poker

/*
To chinses:
diamond 方块
club 梅花
heart 红桃
spade 黑桃

大小王 joker
two,three,four,five,six,seven,eight,nine,ten.
J是Jack Q是Queen K是King A是Ace
*/

//CardSuit suit of card
type CardSuit byte

//CardRank Rank value of card
type CardRank byte

//Card a card of poker,memony value 4bits == |--2 bits suit -- 2 bits rank--|
// type Card byte
type Card struct {
	raw byte
}

var ErrorCard = Card{raw: 0}

const (
	//操作掩码
	CARD_SUIT_MASK = 0xF0 //1111 0000
	CARD_RANK_MASK = 0x0F //0000 1111

	// ErrorCard Card =0x00

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
		return ErrorCard
	}
	raw := byte(suit) << 4
	raw |= byte(rank)
	//(Card(suit) << 4) | Card(rank)
	return Card{raw: raw}
}

func StringToCard(str string) Card {
	if len(str) != 2 {
		return ErrorCard
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
		return ErrorCard
	}
	return CreateCard(suit, rank)
}

//ByteToCard byte to card
func ByteToCard(b byte) Card {
	rank := CardRank(b | CARD_RANK_MASK)
	if !rank.Valid() {
		return ErrorCard
	}
	suit := CardSuit(b | CARD_SUIT_MASK)
	if !suit.Valid() {
		return ErrorCard
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
func (c *Card) Suit() CardSuit {
	return CardSuit(c.raw >> 4)
}

//Rank return card's rank
func (c *Card) Rank() CardRank {
	return CardRank(c.raw & CARD_RANK_MASK)
}

func (c *Card) RankInt() int {
	return int(c.Rank())
}

func (c *Card) Valid() bool {
	return c.Suit().Valid() && c.Rank().Valid()
}

func (c *Card) String() string {
	return string([]byte{suits[c.Suit()], ranks[c.Rank()]})
}

func (c *Card) Byte() byte {
	return c.raw
}
