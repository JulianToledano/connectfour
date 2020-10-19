package main

import "fmt"

func main() {
	game := newGame(newHumanPlayer(1), newIaPlayer(2, 1))
	winner := game.play()
	fmt.Println(winner)
}
