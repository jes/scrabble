package main

import (
	"fmt"
)

func main() {
	dictionary := NewDictionary()
	dictionary.AddFile("dictionary")
	game := NewGame(dictionary)
	p1 := NewRandomAI(1, 50000)
	p2 := NewRandomAI(5, 10000)
	p3 := NewRandomAI(50, 1000)
	game.AddPlayer(p1)
	game.AddPlayer(p2)
	game.AddPlayer(p3)

	for {
		game.OneTurn()
		fmt.Printf("Player 1: %d\n", game.scores[0])
		fmt.Printf("Player 2: %d\n", game.scores[1])
		fmt.Printf("Player 3: %d\n", game.scores[2])
		fmt.Printf("%v\n\n", game.board)
		fmt.Printf("player 1 letters: %v\n", p1.letters)
		fmt.Printf("player 2 letters: %v\n", p2.letters)
		fmt.Printf("player 3 letters: %v\n", p3.letters)
	}
}
