package main

func main() {
	game := newGame()
	game.makeMove(game.playerOne, 3)
	game.printBoard()
}
