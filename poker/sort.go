package poker

type CardSort = []Card

// ByRank - container for cards sorted by value.
type ByRank CardSort

// BySuit - container for cards sorted by suit.
type BySuit CardSort

// ByValue implements Sort method
func (hand ByRank) Len() int {
	return len(hand)
}

func (hand ByRank) Less(i, j int) bool {
	iRank, jRank := hand[i].RankInt(), hand[j].RankInt()
	if iRank == jRank {
		return hand[i].SuitInt() < hand[j].SuitInt()
	}
	return iRank < jRank
}

func (hand ByRank) Swap(i, j int) {
	hand[i], hand[j] = hand[j], hand[i]
}

// BySuit implements Sort method
func (hand BySuit) Len() int {
	return len(hand)
}

func (hand BySuit) Less(i, j int) bool {
	iSuit, jSuit := hand[i].SuitInt(), hand[j].SuitInt()
	if iSuit == jSuit {
		return hand[i].RankInt() < hand[j].RankInt()
	}
	return iSuit < jSuit
}

func (hand BySuit) Swap(i, j int) {
	hand[i], hand[j] = hand[j], hand[i]
}
