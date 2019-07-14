package main

import (
	"fmt"

	// "../../game2048"
	"github.com/cmp102/game2048"
)

func main() {
	g2048 := game2048.NewGame2048(4)
	g2048.AddRandom()
	for {
		g2048.Print()
		fmt.Printf("Possible moves: ")
		mapMoves := make(map[string]func())
		if g2048.ValidateUP() {
			mapMoves["u"] = g2048.MoveUP
			fmt.Printf("u(UP) ")
		}
		if g2048.ValidateDOWN() {
			mapMoves["d"] = g2048.MoveDOWN
			fmt.Printf("d(DOWN) ")
		}
		if g2048.ValidateLEFT() {
			mapMoves["l"] = g2048.MoveLEFT
			fmt.Printf("l(LEFT) ")
		}
		if g2048.ValidateRIGHT() {
			mapMoves["r"] = g2048.MoveRIGHT
			fmt.Printf("r(RIGHT) ")
		}
		fmt.Printf("\n>> ")
		move := ""
		fmt.Scanln(&move)
		fmt.Printf("Move: %s\n", move)
		if mv, ok := mapMoves[move]; ok {
			mv()
			g2048.AddRandom()
		}
	}
}
