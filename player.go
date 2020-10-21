package main

import (
	"fmt"
	"math/rand"
	"time"
)

type player interface {
	chooseMove() int
	getColor() int
}

type humanPlayer struct {
	color int
}

type iaPlayer struct {
	color     int
	difficult int
}

func newHumanPlayer(color int) *humanPlayer {
	hp := new(humanPlayer)
	hp.color = color
	return hp
}

func (p humanPlayer) chooseMove() int {
	move := -1
	fmt.Println("Choose a column")
	for move < 0 || move > 7 {
		_, err := fmt.Scan(&move)
		if err != nil {
			fmt.Println("Not valid input...")
		}
	}
	return move
}

func (p humanPlayer) getColor() int {
	return p.color
}

func newIaPlayer(color, difficult int) *iaPlayer {
	ia := new(iaPlayer)
	ia.color = color
	ia.difficult = difficult
	return ia
}

func (ia iaPlayer) chooseMove() int {
	return ia.easyMode()

}

func (ia iaPlayer) getColor() int {
	return ia.color
}

func (ia iaPlayer) easyMode() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(7) + 1
}
