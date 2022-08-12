package main

import (
	"fmt"

	"github.com/ajenpan/poker_algorithm/niuniu"
)

func main() {
	const playerCount = 5

	type Player struct {
		cards *niuniu.NNHandCards
	}
	deck := niuniu.NewNNDeck()
	deck.Shuffle()

	players := []Player{}

	for i := 0; i < playerCount; i++ {
		p := Player{
			cards: deck.DealHandCards(),
		}
		p.cards.Calculate()
		players = append(players, p)
	}

	banker := players[0]
	fmt.Printf("banker's cards:%s,%v \n", banker.cards.String(), banker.cards.Type())
	for i := 1; i < playerCount; i++ {
		fmt.Printf("player%d's cards:%s,%v \n", i, players[i].cards.String(), players[i].cards.Type())
		if banker.cards.Compare(players[i].cards) {
			fmt.Println("banker win")
		} else {
			fmt.Println("banker lose")
		}
	}
}
