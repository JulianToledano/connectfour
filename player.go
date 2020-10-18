package main

import "fmt"

type humanPlayer struct {
	color int
}

func newHumanPlayer(color int) *humanPlayer {
	hp := new(humanPlayer)
	hp.color = color
	return hp
}

func (p *humanPlayer) chooseMove() int {
	var move int
	fmt.Println("Choose a column")
	_, err := fmt.Scan(&move)
	if err != nil {
		fmt.Println("Somethig went wrong!!")

	}
	return move
}
