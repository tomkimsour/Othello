// package alpha-beta-pruning gives functions to perform
// iterative alpha beta pruning search on a Board structure
package main

import (
	"container/list"
	"time"
)

const PosInfty int = 2147483647
const NegInfty int = -2147483648
const MaxPlayer bool = true
const MinPlayer bool = false

var semaphore bool = false

type AlphaBetaPruning struct {
	moves   list.List
	timeout int
}

// creates a new Structure AlphaBetaPruning with a depth and timeout value
func NewAlphaBetaPruning(time int) *AlphaBetaPruning {
	return new(AlphaBetaPruning).Init(time)
}

// Initialize a structure AlphaBetaPruning
func (abp *AlphaBetaPruning) Init(time int) *AlphaBetaPruning {
	abp.timeout = time
	return abp
}

// Evaluate function execute a deepening search from a board and
// returns the best move among the iterations according to heuristic
func (abp *AlphaBetaPruning) Evaluate(position *Board) *Action {
	var finalPosition, lastPosition *Action

	// go routine that stops the program after timeout using a boolean semaphore
	go func(timeout int) {
		time.Sleep(time.Duration(timeout) * time.Second)
		semaphore = true
	}(abp.timeout)

	for i := 4; !semaphore; i++ {
		lastPosition = abp.alphabeta(position, i, NegInfty, PosInfty)
		if !semaphore {
			finalPosition = lastPosition
		}
	}
	return finalPosition
}

// takes the minimum value between a and b
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// takes the maximum value between a and b
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// alphabeta is a recursive function that takes a current board, depth, alpha and beta values and
// returns the best Action according to heuristic
func (abp *AlphaBetaPruning) alphabeta(position *Board, depth, alpha, beta int) *Action {
	moveList := position.GetMoves()
	var moveValue int

	lastAction := NewAction(0, 0)

	if depth == 0 {
		moveValue = Evaluate(position)
		lastAction.SetValue(moveValue)
		return lastAction
	}
	if moveList.Front() == nil {
		moveValue = Evaluate(position)
		lastAction.SetValue(moveValue)
		lastAction.SetPassMove(true)
		return lastAction
	}

	var nextBoard *Board

	if position.GetMaxPlayer() {
		moveValue = NegInfty
		lastAction.SetValue(moveValue)
		for e := moveList.Front(); e != nil; e = e.Next() {
			if semaphore {
				return NewAction(0, 0)
			}
			move := e.Value.(*Action)
			nextBoard = position.MakeMove(move)
			moveValue = max(moveValue, abp.alphabeta(nextBoard, depth-1, alpha, beta).GetValue())
			if alpha >= beta {
				break
			}
			if alpha < moveValue {
				lastAction = move
				alpha = moveValue
			}
		}
		lastAction.SetValue(moveValue)
		return lastAction
	} else {
		moveValue = PosInfty
		lastAction := NewAction(0, 0)
		lastAction.SetValue(moveValue)
		for e := moveList.Front(); e != nil; e = e.Next() {
			if semaphore {
				return NewAction(0, 0)
			}
			move := e.Value.(*Action)
			nextBoard = position.MakeMove(move)
			moveValue = min(moveValue, abp.alphabeta(nextBoard, depth-1, alpha, beta).GetValue())
			if beta <= alpha {
				break
			}
			if beta > moveValue {
				beta = moveValue
				lastAction = move
			}
		}
		lastAction.SetValue(moveValue)
		return lastAction
	}
}
