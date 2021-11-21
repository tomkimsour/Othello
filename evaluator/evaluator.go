package evaluator

import (
	"github.com/tomkimsour/Othello/board"
)

type Evaluator struct {
	value int
}

func coinParity(position *board.Board) int {
	boardSize := board.BoardSize
	maxCounter := 0
	minCounter := 0
	for i := 1; i < boardSize-1; i++ {
		for j := 1; j < boardSize-1; j++ {
			if position.Board[i][j] == 'B' {
				minCounter++
			} else if position.Board[i][j] == 'W' {
				maxCounter++
			}
		}
	}
	if maxCounter+minCounter != 0 {
		return 100 * (maxCounter - minCounter) / (maxCounter + minCounter)
	}
	return 0
}

func mobility(position *board.Board) int {
	var maxCounter, minCounter int

	if position.GetMaxPlayer() == false {
		maxCounter = position.GetMoves().Len()
		position.ChangeMaxPlayer()
		minCounter = position.GetMoves().Len()
		position.ChangeMaxPlayer()
	} else {
		minCounter = position.GetMoves().Len()
		position.ChangeMaxPlayer()
		maxCounter = position.GetMoves().Len()
		position.ChangeMaxPlayer()
	}

	if maxCounter+minCounter != 0 {
		return 100 * (maxCounter - minCounter) / (maxCounter + minCounter)
	}
	return 0
}

func cornersCaptured(position *board.Board) int {
	var maxCounter, minCounter int

	if position.Board[1][1] == 'B' {
		minCounter++
	} else if position.Board[1][1] == 'W' {
		maxCounter++
	}

	if position.Board[8][8] == 'B' {
		minCounter++
	} else if position.Board[8][8] == 'W' {
		maxCounter++
	}

	if position.Board[8][1] == 'B' {
		minCounter++
	} else if position.Board[8][1] == 'W' {
		maxCounter++
	}

	if position.Board[1][8] == 'B' {
		minCounter++
	} else if position.Board[1][8] == 'W' {
		maxCounter++
	}

	if maxCounter+minCounter != 0 {
		return 100 * (maxCounter - minCounter) / (maxCounter + minCounter)
	}
	return 0
}

/** Returns an integer, representing a heuristic evaluation of the postion. */
func (e *Evaluator) Evaluate(position *board.Board) int {

	// coin parity
	mobilityIndex := mobility(position)
	coinParityIndex := coinParity(position)
	cornersCapturedIndex := cornersCaptured(position)

	return mobilityIndex*100 + coinParityIndex*20 + cornersCapturedIndex*800
}
