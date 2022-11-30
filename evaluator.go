// Package evaluator gives functions that gives heuristic value of a Board structure
package main

// coinParity return an integer representing the coin parity ranged from -100 to 100
// given the pointer of a board state
func coinParity(position *Board) int {
	boardSize := BoardSize
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

// mobility return an integer representing the number of moves possible ranged from -100 to 100
// given the pointer of a board state
func mobility(position *Board) int {
	var maxCounter, minCounter int

	maxCounter = position.GetMoves().Len()
	position.ChangeMaxPlayer()
	minCounter = position.GetMoves().Len()
	position.ChangeMaxPlayer()

	if maxCounter+minCounter != 0 {
		return 100 * (maxCounter - minCounter) / (maxCounter + minCounter)
	}
	return 0
}

// cornersCaptured return an integer representing the possession of corners between -100 and 100
// given the pointer of a board state
func cornersCaptured(position *Board) int {
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

/* Evaluate returns an integer, representing a heuristic evaluation of the postion. */
func Evaluate(position *Board) int {

	// coin parity
	mobilityIndex := mobility(position)
	coinParityIndex := coinParity(position)
	cornersCapturedIndex := cornersCaptured(position)

	return mobilityIndex*100 + coinParityIndex*20 + cornersCapturedIndex*800
}
