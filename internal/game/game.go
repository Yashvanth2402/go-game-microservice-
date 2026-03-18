package game

import (
	"fmt"
)

type Game struct {
	min     int
	max     int
	target  int
	attempts int
}

func NewGame(min, max int) *Game {
	return &Game{
		min: min,
		max: max,
		target: GenerateRandom(min, max),
	}
}

func (g *Game) Start() {
	fmt.Println("Guess a number between", g.min, "and", g.max)

	for {
		var input int
		fmt.Print("Enter your guess: ")
		fmt.Scan(&input)

		g.attempts++

		result := g.CheckGuess(input)

		if result == "correct" {
			fmt.Println("🎉 Correct! Attempts:", g.attempts)
			break
		} else {
			fmt.Println(result)
		}
	}
}
