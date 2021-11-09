package evaluator

import (
	"github.com/tomkimsour/Othello/board"
)

type Evaluator struct {
	value int
}

var color rune

func coinParity(position *board.Board) (int, int) {
	boardSize := board.BoardSize
	myCounter := 0
	opponentCounter := 0
	for i := 1; i < boardSize-1; i++ {
		for j := 1; j < boardSize-1; j++ {
			if position.Board[i][j] != 'E' {
				if position.Board[i][j] == color {
					myCounter++
				} else {
					opponentCounter++
				}
			}
		}
	}
	return myCounter, opponentCounter
}

func mobility(position *board.Board) (int, int) {
	myCounter := position.GetMoves().Len()
	position.ChangeMaxPlayer()
	opponentCounter := position.GetMoves().Len()
	return myCounter, opponentCounter
}

/** Returns an integer, representing a heuristic evaluation of the postion. */
func (e *Evaluator) Evaluate(position *board.Board) int {
	// Black and white are reversed
	if position.GetMaxPlayer() {
		color = 'W'
	} else {
		color = 'B'
	}
	var myCounter, opponentCounter int

	// coin parity
	myCounter, opponentCounter = mobility(position)
	return myCounter - opponentCounter
}
