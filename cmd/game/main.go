package main

import (
	"fmt"
	"github.com/your-username/go-game-microservice/internal/game"
)

func main() {
	fmt.Println("🎮 Welcome to Guess the Number Game")

	g := game.NewGame(1, 100)
	g.Start()
}
