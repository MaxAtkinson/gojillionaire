package main

import (
	"fmt"
	"github.com/MaxAtkinson/gojillionaire/pkg/game"
	"github.com/MaxAtkinson/gojillionaire/pkg/questions"
)

func main() {
	var game game.Game
	questions := questions.ReadQuestions()

	for _, question := range questions {
		answer := question.Ask()
		isCorrect := question.IsCorrectAnswer(answer)

		if isCorrect {
			game.IncrementScore()
			game.PrintSummary()
		} else {
			fmt.Println("\nGAME OVER.")
			return
		}
	}

	fmt.Println("\nYou're a Gojillionaire.")
}
