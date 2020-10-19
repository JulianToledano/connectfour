package main

import (
	"fmt"
	"math/rand"
	"time"
)

type game struct {
	board     [6][7]int
	rows      int
	columns   int
	playerOne player
	playerTwo player
}

func newGame(playerOne, playerTwo player) *game {
	g := new(game)
	g.rows = 6
	g.columns = 7
	g.playerOne = playerOne
	g.playerTwo = playerTwo
	return g
}

func (g *game) play() int {
	rand.Seed(time.Now().UnixNano())
	player := g.choosePlayer(rand.Intn(2) + 1)
	winner := 0
	for true {
		moveMade := false
		for !moveMade {
			choosenMove := player.chooseMove()
			moveMade = g.makeMove(player.getColor(), choosenMove)
		}
		g.printBoard()
		if g.checkWin(player.getColor()) {
			winner = player.getColor()
			break
		}
		if g.isBoardFull() {
			break
		}
		player = g.choosePlayer(player.getColor())
	}
	return winner
}

func (g *game) choosePlayer(playerNumber int) player {
	if playerNumber == g.playerOne.getColor() {
		return g.playerTwo
	}
	return g.playerOne

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

func (g *game) isBoardFull() bool {
	for i := 0; i < g.columns; i++ {
		if !g.isColumnFull(i) {
			return false
		}
	}
	return true
}

func (g *game) checkWin(player int) bool {
	if g.checkWinRows(player) {
		return true
	}
	if g.checkWinColumns(player) {
		return true
	}
	if g.checkWinRightDiagonal(player) {
		return true
	}
	if g.checkWinLeftDiagonal(player) {
		return true
	}
	return false
}

func (g *game) checkWinRows(player int) bool {
	for i := 0; i < g.rows; i++ {
		if g.checkSubSecuence(player, i, 0, 0, 1) {
			return true
		}
	}
	return false
}

func (g *game) checkWinColumns(player int) bool {
	for i := 0; i < g.columns; i++ {
		if g.checkSubSecuence(player, 0, i, 1, 0) {
			return true
		}
	}
	return false
}

func (g *game) checkWinRightDiagonal(player int) bool {
	for row := 0; row < g.rows; row++ {
		for column := 0; column < g.columns; column++ {
			if g.checkSubSecuence(player, row, column, 1, 1) {
				return true
			}
		}
	}
	return false
}

func (g *game) checkWinLeftDiagonal(player int) bool {
	for row := 0; row < g.rows; row++ {
		for column := g.columns - 1; column >= 0; column-- {
			if g.checkSubSecuence(player, row, column, 1, -1) {
				return true
			}
		}
	}
	return false
}

func (g *game) checkSubSecuence(player int, row int, column int, rowMove int, columnMove int) bool {
	lastCell := 0
	contiguousCells := 0
	for row < g.rows && column < g.columns && row >= 0 && column >= 0 {
		lastCell = g.board[row][column]
		if lastCell == player {
			contiguousCells++
		} else {
			contiguousCells = 0
		}
		if contiguousCells == 4 {
			return true
		}
		row += rowMove
		column += columnMove
	}
	return false
}

func (g *game) makeMove(player int, column int) bool {
	column--
	if !g.isColumnFull(column) {
		g.insertToken(player, column)
		return true
	}
	return false
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
	fmt.Println("----------------------------")
}
