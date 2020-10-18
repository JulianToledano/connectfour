package main

import "fmt"

func main() {
	game := newGame()
	winner := game.play()
	fmt.Println(winner)
}
