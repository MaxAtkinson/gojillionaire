package game

import (
	"fmt"
)

var PRIZE_MONEY_INCREMENTS = []int{0, 500, 1000, 2000, 3000, 5000, 7000, 10000, 20000, 30000, 50000, 100000, 250000, 500000, 1000000}

type Game struct {
	PrizeMoneyIndex int
}

func (g *Game) IncrementScore() {
	g.PrizeMoneyIndex += 1
}

func (g *Game) PrintSummary() {
	fmt.Println(fmt.Sprintf("\nYou're on £%d",
		PRIZE_MONEY_INCREMENTS[g.PrizeMoneyIndex]))
}
