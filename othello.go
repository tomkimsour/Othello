package main

import (
	"log"
	"os"
	"strconv"

	alphabetapruning "github.com/tomkimsour/Othello/alpha-beta-pruning"
	"github.com/tomkimsour/Othello/board"
)

func main() {
	var sequence string
	var time int
	if len(os.Args) > 1 {
		sequence = os.Args[1]
	} else {
		log.Fatal("No string was given")
	}
	if len(os.Args) > 2 {
		time, _ = strconv.Atoi(os.Args[2])
	} else {
		log.Fatal("No timeout given")
	}
	if len(sequence) != 65 {
		log.Fatal("The string has to be of size 65")
	}

	currentBoard := board.New()

	currentBoard.OthelloPosition(sequence)
	abp := alphabetapruning.New(currentBoard, 9, time)

	move := abp.Evaluate(currentBoard)
	move.Print()
}
