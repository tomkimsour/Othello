package main

import (
	"fmt"
	"os"

	Board "github.com/tomkimsour/Othello/othello-functions"
)

func main() {
	var sequence string
	if len(os.Args) > 1 {
		sequence = os.Args[1]
	} else {
		sequence = "WEEEEEEEEEEEEEEEEEEEEEEEEEEEOXEEEEEEXOEEEEEEEEEEEEEEEEEEEEEEEEEEE"
	}
	fmt.Println(sequence)
	// var board othellofunctions.Board
	board := Board.New()
	board.OthelloPosition(sequence)
	board.Print()
	// othelloPosition := OthelloPosition.Ini(sequence)
	// fmt.Print(othelloPosition.toString())
}
