package main

import "fmt"

func main() {
	game := newGame()
	game.makeMove(game.playerOne, 4)
	game.makeMove(game.playerOne, 4)
	game.makeMove(game.playerOne, 4)
	game.makeMove(game.playerTwo, 4)

	game.makeMove(game.playerOne, 3)
	game.makeMove(game.playerOne, 3)
	game.makeMove(game.playerTwo, 3)

	game.makeMove(game.playerOne, 2)
	game.makeMove(game.playerTwo, 2)

	game.makeMove(game.playerTwo, 1)

	game.printBoard()
	if game.checkWin(game.playerTwo) {
		fmt.Printf("hello")
	}
}
