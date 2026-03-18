package test

import (
	"testing"
	"github.com/your-username/go-game-microservice/internal/game"
)

func TestCheckGuess(t *testing.T) {
	g := game.NewGame(1, 10)
	gTarget := g

	result := gTarget.CheckGuess(gTarget.CheckGuess(5))

	if result == "" {
		t.Errorf("Expected a response, got empty")
	}
}
