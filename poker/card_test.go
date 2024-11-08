package poker

import (
	"math/rand"
	"testing"
	"time"
)

func TestCard(t *testing.T) {
	for s := 1; s <= 4; s++ {
		for r := 1; r <= 13; r++ {
			card := NewCard(CardSuit(s), CardRank(r))
			if int(card.Suit()) != s || int(card.Rank()) != r {
				t.FailNow()
			}
		}
	}

	for s := 1; s <= 4; s++ {
		for r := 14; r <= 26; r++ {
			card := NewCard(CardSuit(s), CardRank(r))
			if int(card.Suit()) == s {
				t.FailNow()
			}
			if card != EmptyCard {
				t.FailNow()
			}
		}
	}
}
func TestPointSize(t *testing.T) {
	card1 := NewCard(CardSuit(1), CardRank(1))
	card2 := NewCard(CardSuit(1), CardRank(2))

	if byte(card1) > byte(card2) {
		t.Fail()
	}
}

func TestDeckIsEmpty(t *testing.T) {
	deck := NewEmptyCards()
	if !deck.IsEmpty() {
		t.Error("Deck should be empty but is not.")
	}
}
func TestStr2Cards(t *testing.T) {
	rand.Seed(time.Hour.Nanoseconds())
	for i := 0; i < 100; i++ {
		randSuit := rand.Int31n(4) + 1
		randRank := rand.Int31n(13) + 1
		card := NewCard(CardSuit(randSuit), CardRank(randRank))
		if card.String() != string([]byte{byte(Ranks[randRank]), byte(Suits[randSuit])}) {
			t.Errorf("%s != %s", card.String(), string([]byte{byte(randSuit), byte(randRank)}))
			t.FailNow()
		}
	}
}
func TestNewCard(t *testing.T) {
	tests := []struct {
		suit CardSuit
		rank CardRank
		want Card
	}{
		{DIAMOND, 1, DIAMOND_A},
		{CLUB, 13, CLUB_K},
		{HEART, 5, HEART_5},
		{SPADE, 10, SPADE_10},
		{JOKER, 1, JOKER_BLACK},
		{JOKER, 2, JOKER_RED},
		{CardSuit(6), 1, EmptyCard}, // Invalid suit
		{DIAMOND, 14, EmptyCard},    // Invalid rank
	}

	for _, tt := range tests {
		got := NewCard(tt.suit, tt.rank)
		if got != tt.want {
			t.Errorf("NewCard(%v, %v) = %v; want %v", tt.suit, tt.rank, got, tt.want)
		}
	}
}

func TestNewCardByString(t *testing.T) {
	tests := []struct {
		str  string
		want Card
	}{
		{"DA", DIAMOND_A},
		{"CK", CLUB_K},
		{"H5", HEART_5},
		{"S10", SPADE_10},
		{"J1", JOKER_BLACK},
		{"J2", JOKER_RED},
		{"X1", EmptyCard},  // Invalid suit
		{"D14", EmptyCard}, // Invalid rank
		{"", EmptyCard},    // Invalid length
	}

	for _, tt := range tests {
		got := NewCardByString(tt.str)
		if got != tt.want {
			t.Errorf("NewCardByString(%v) = %v; want %v", tt.str, got, tt.want)
		}
	}
}

func TestByteToCard(t *testing.T) {
	tests := []struct {
		b    byte
		want Card
	}{
		{0x11, DIAMOND_A},
		{0x2D, CLUB_K},
		{0x35, HEART_5},
		{0x4A, SPADE_10},
		{0x51, JOKER_BLACK},
		{0x52, JOKER_RED},
		{0x00, EmptyCard}, // Invalid card
	}

	for _, tt := range tests {
		got := ByteToCard(tt.b)
		if got != tt.want {
			t.Errorf("ByteToCard(%v) = %v; want %v", tt.b, got, tt.want)
		}
	}
}

func TestCardMethods(t *testing.T) {
	card := NewCard(DIAMOND, 5)

	if card.Suit() != DIAMOND {
		t.Errorf("card.Suit() = %v; want %v", card.Suit(), DIAMOND)
	}

	if card.Rank() != 5 {
		t.Errorf("card.Rank() = %v; want %v", card.Rank(), 5)
	}

	if !card.Valid() {
		t.Errorf("card.Valid() = %v; want %v", card.Valid(), true)
	}

	if card.String() != "5D" {
		t.Errorf("card.String() = %v; want %v", card.String(), "5D")
	}

	if card.Byte() != 0x15 {
		t.Errorf("card.Byte() = %v; want %v", card.Byte(), 0x15)
	}

	newCard := card.SetSuit(CLUB)
	if newCard.Suit() != CLUB || newCard.Rank() != 5 {
		t.Errorf("card.SetSuit(CLUB) = %v; want suit %v and rank %v", newCard, CLUB, 5)
	}

	newCard = card.SetRank(10)
	if newCard.Suit() != DIAMOND || newCard.Rank() != 10 {
		t.Errorf("card.SetRank(10) = %v; want suit %v and rank %v", newCard, DIAMOND, 10)
	}

	if card.Chinese() != "方块5" {
		t.Errorf("card.Chinese() = %v; want %v", card.Chinese(), "方块5")
	}
}
