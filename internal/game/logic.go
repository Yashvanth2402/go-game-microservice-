package game

import (
	"math/rand"
	"time"
)

func GenerateRandom(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func (g *Game) CheckGuess(input int) string {
	if input < g.target {
		return "Too low ⬇️"
	} else if input > g.target {
		return "Too high ⬆️"
	}
	return "correct"
}
