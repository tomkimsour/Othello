package main

import (
	"os"

	alphabetapruning "github.com/tomkimsour/Othello/alpha-beta-pruning"
	"github.com/tomkimsour/Othello/board"
)

func main() {
	var sequence string
	if len(os.Args) > 1 {
		sequence = os.Args[1]
	} else {
		// sequence = "WEEEEEEEEEXEXEXEEEEXXXEEEEXXOXXEEEEXXOEEEEXEXEXEEEEEEEEXEEEEEEEEE"
		sequence = "WEEEEEEEEEEEEEEEEEEEEEEEEEEEOXEEEEEEXOEEEEEEEEEEEEEEEEEEEEEEEEEEE"
	}

	currentBoard := board.New()

	currentBoard.OthelloPosition(sequence)
	abp := alphabetapruning.New(currentBoard, 7)

	move := abp.Evaluate(currentBoard)
	// nextBoard := currentBoard.MakeMove(move)
	move.Print()

	// currentBoard.Print()
	// var moves *list.List
	// moves = currentBoard.GetMoves()
	// var newBoard *board.Board
	// for e := moves.Front(); e != nil; e = e.Next() {
	// 	value := e.Value.(*action.Action)
	// 	// value.Print()
	// 	newBoard = currentBoard.MakeMove(value)
	// 	newBoard.Print()
	// }
	// newBoard.Print()
	// fmt.Println(newBoard.ToString())

	// othelloPosition := OthelloPosition.Ini(sequence)
	// fmt.Print(othelloPosition.toString())
}
