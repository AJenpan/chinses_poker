package ginrummy

import (
	"fmt"
	"math/rand"

	"github.com/ajenpan/poker_algorithm/poker"
)

type Robot struct {
	// HandCards *poker.Cards
	Hand *HandCards
}

func NewRobot() *Robot {
	return &Robot{}
}
func (r *Robot) OnTurn() {

}

func (r *Robot) OnFristTurn(card poker.Card) (pass bool) {

	return false
}

func (r *Robot) StartWithCards(cards *poker.Cards) {
	r.Hand = &HandCards{
		Cards: cards,
	}
}

func (r *Robot) Status() string {
	return fmt.Sprintf("handcards: %v", r.Hand.Cards.Chinese())
}

func (r *Robot) DrawAndDiscard(card poker.Card) poker.Card {
	return card
}

type Game struct {
	DiscardStack []poker.Card

	p1 *Robot
	p2 *Robot

	deck *poker.Cards
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Start() {
	deck := poker.NewDeckWithoutJoker()
	deck.Shuffle()

	g.p1 = NewRobot()
	g.p2 = NewRobot()

	card := deck.DealCard()
	g.DiscardStack = append(g.DiscardStack, card)

	g.p1.StartWithCards(deck.DealCards(10))
	g.p2.StartWithCards(deck.DealCards(10))

	g.deck = deck

	bankSet := rand.Int31n(1) + 1

	if bankSet == 1 {

	} else {

	}
}

func (g *Game) PrintStatus() {
	fmt.Println("deck:", g.deck.Chinese())
	fmt.Println("discard:", g.DiscardStack[len(g.DiscardStack)-1].Chinese())

	fmt.Println("p1:", g.p1.Status())
	fmt.Println("p2:", g.p2.Status())
}
