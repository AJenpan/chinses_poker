package guandan

import (
	"testing"

	"github.com/ajenpan/poker_algorithm/poker"
)

func TestDeckPower_Compare(t *testing.T) {
	tests := []struct {
		name     string
		dp       DeckPower
		other    DeckPower
		expected int
	}{

		{
			name:     "DeckPass vs DeckSingle",
			dp:       DeckPower{DeckType: DeckPass, DeckValue: 0},
			other:    DeckPower{DeckType: DeckSingle, DeckValue: 1},
			expected: -1,
		},
		{
			name:     "DeckSingle vs DeckPass",
			dp:       DeckPower{DeckType: DeckSingle, DeckValue: 1},
			other:    DeckPower{DeckType: DeckPass, DeckValue: 0},
			expected: 1,
		},
		{
			name:     "Same DeckType, dp > other",
			dp:       DeckPower{DeckType: DeckSingle, DeckValue: 2},
			other:    DeckPower{DeckType: DeckSingle, DeckValue: 1},
			expected: 1,
		},
		{
			name:     "Same DeckType, dp < other",
			dp:       DeckPower{DeckType: DeckSingle, DeckValue: 1},
			other:    DeckPower{DeckType: DeckSingle, DeckValue: 2},
			expected: -1,
		},
		{
			name:     "Same DeckType, dp == other",
			dp:       DeckPower{DeckType: DeckSingle, DeckValue: 1},
			other:    DeckPower{DeckType: DeckSingle, DeckValue: 1},
			expected: 0,
		},
		{
			name:     "Different non-bomb DeckTypes",
			dp:       DeckPower{DeckType: DeckSingle, DeckValue: 1},
			other:    DeckPower{DeckType: DeckPair, DeckValue: 1},
			expected: -2,
		},
		{
			name:     "dp is bomb, other is not",
			dp:       DeckPower{DeckType: DeckBomb4, DeckValue: 1},
			other:    DeckPower{DeckType: DeckSingle, DeckValue: 1},
			expected: 1,
		},
		{
			name:     "dp is not bomb, other is",
			dp:       DeckPower{DeckType: DeckSingle, DeckValue: 1},
			other:    DeckPower{DeckType: DeckBomb4, DeckValue: 1},
			expected: -1,
		},
		{
			name:     "Both are bombs, dp > other",
			dp:       DeckPower{DeckType: DeckBomb5, DeckValue: 1},
			other:    DeckPower{DeckType: DeckBomb4, DeckValue: 1},
			expected: 1,
		},
		{
			name:     "Both are bombs, dp < other",
			dp:       DeckPower{DeckType: DeckBomb4, DeckValue: 1},
			other:    DeckPower{DeckType: DeckBomb5, DeckValue: 1},
			expected: -1,
		},
		{
			name:     "DeckWindflow comparison",
			dp:       DeckPower{DeckType: DeckWindflow, DeckValue: 1},
			other:    DeckPower{DeckType: DeckSingle, DeckValue: 1},
			expected: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.dp.Compare(&tt.other)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestGetDeckPowerWithNoWildcard(t *testing.T) {
	tests := []struct {
		name     string
		cardstr  string
		expected *DeckPower
	}{
		{
			name:     "Pair",
			cardstr:  "H3 H3",
			expected: &DeckPower{DeckType: DeckPair, DeckValue: 3},
		},
		{
			name:     "Three of a kind",
			cardstr:  "S4 H4 C4",
			expected: &DeckPower{DeckType: DeckThree, DeckValue: 4},
		},
		{
			name:     "Four of a kind",
			cardstr:  "S5 H5 C5 D5",
			expected: &DeckPower{DeckType: DeckBomb4, DeckValue: 5},
		},
		{
			name:     "Straight",
			cardstr:  "S3 H4 C5 D6 S7",
			expected: &DeckPower{DeckType: DeckStraight, DeckValue: 7},
		},
		{
			name:     "Straight Flush",
			cardstr:  "S3 S4 S5 S6 S7",
			expected: &DeckPower{DeckType: DeckStraightFlush, DeckValue: 7},
		},
		{
			name:     "Three with Two",
			cardstr:  "S8 H8 C8 D9 S9",
			expected: &DeckPower{DeckType: DeckThreeWithTwo, DeckValue: 8},
		},
		{
			name:     "Three with Two error case",
			cardstr:  "S9 H9 C9 D8 S7",
			expected: &DeckPower{DeckType: DeckPass, DeckValue: 0},
		},
		{
			name:     "Invalid Straight",
			cardstr:  "S3 H4 C5 D6 S8",
			expected: &DeckPower{DeckType: DeckPass, DeckValue: 0},
		},
		{
			name:     "Six of a kind",
			cardstr:  "S6 H6 C6 D6 S6 H6",
			expected: &DeckPower{DeckType: DeckBomb6, DeckValue: 6},
		},
		{
			name:     "Seven of a kind",
			cardstr:  "S7 H7 C7 D7 S7 H7 C7",
			expected: &DeckPower{DeckType: DeckBomb7, DeckValue: 7},
		},
		{
			name:     "Eight of a kind",
			cardstr:  "S8 H8 C8 D8 S8 H8 C8 D8",
			expected: &DeckPower{DeckType: DeckBomb8, DeckValue: 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards, err := poker.StringToCards(tt.cardstr)
			if err != nil {
				t.Fatal(err)
			}
			result := GetDeckPower(2, cards)
			if result.DeckType != tt.expected.DeckType || result.DeckValue != tt.expected.DeckValue {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
