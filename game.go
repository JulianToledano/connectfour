package main

import "fmt"

type game struct {
	board     [6][7]int
	rows      int
	columns   int
	playerOne int
	platerTwo int
}

func newGame() *game {
	g := new(game)
	g.rows = 6
	g.columns = 7
	g.playerOne = 1
	g.platerTwo = 2
	return g
}

func (g *game) completeTurn(player int, column int) {
	// Player makes a move
	g.makeMove(player, column)
	// Check if game ended

}

func (g *game) checkWin(player int) {
	// Check for columns
	// check for rows
	// Check right-diagonal
	// Check left-diagonal
}

func (g *game) checkWinRows(player int) {
	for i := 0; i < g.rows; i++ {
		g.checkSubSecuence(i, 0, 0, 1)
	}
}

func (g *game) checkSubSecuence(row int, column int, rowMove int, columnMove int) bool {

}

func (g *game) makeMove(player int, column int) {
	column--
	if !g.isColumnFull(column) {
		g.insertToken(player, column)
	}
}

func (g *game) isCellEmpty(row int, column int) bool {
	if g.board[row][column] != 0 {
		return false
	}
	return true
}

func (g *game) isColumnFull(column int) bool {
	return !g.isCellEmpty(0, column)
}

func (g *game) insertToken(player int, column int) {
	for i := g.rows - 1; i >= 0; i-- {
		if g.isCellEmpty(i, column) {
			g.board[i][column] = player
			break
		}
	}
}

func (g *game) printBoard() {
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.columns; j++ {
			fmt.Print(g.board[i][j])
			fmt.Print(" | ")
		}
		fmt.Println()
	}
}
