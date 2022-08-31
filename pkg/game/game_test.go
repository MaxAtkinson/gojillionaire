package game

import (
	"testing"
)

func TestGame(t *testing.T) {
	var game Game

	if game.PrizeMoneyIndex != 0 {
		t.Errorf("got %q, wanted %q", game.PrizeMoneyIndex, 0)
	}

	game.IncrementScore()

	if game.PrizeMoneyIndex != 1 {
		t.Errorf("got %q, wanted %q", game.PrizeMoneyIndex, 1)
	}
}
