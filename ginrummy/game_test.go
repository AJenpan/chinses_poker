package ginrummy

import "testing"

func TestGameStart(t *testing.T) {

	g := NewGame()
	g.Start()

	g.PrintStatus()
}
