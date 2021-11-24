package alphabetapruning

import (
	"container/list"
	"time"

	"github.com/tomkimsour/Othello/action"
	"github.com/tomkimsour/Othello/board"
	"github.com/tomkimsour/Othello/evaluator"
)

const PosInfty int = 2147483647
const NegInfty int = -2147483648
const MaxPlayer bool = true
const MinPlayer bool = false

var semaphore bool = false

type AlphaBetaPruning struct {
	searchDepth int
	moves       list.List
	timeout     int
}

// creates a new Structure AlphaBetaPruning with a depth and timeout value
func New(depth, time int) *AlphaBetaPruning {
	return new(AlphaBetaPruning).Init(depth, time)
}

// Initialize a structure AlphaBetaPruning
func (abp *AlphaBetaPruning) Init(depth, time int) *AlphaBetaPruning {
	abp.searchDepth = depth
	abp.timeout = time
	return abp
}

// Evaluate function execute a deepening search from a board and
// returns the best move among the iterations according to heuristic
func (abp *AlphaBetaPruning) Evaluate(position *board.Board) *action.Action {
	var finalPosition, lastPosition *action.Action

	// go routine that stops the program after timeout using a boolean semaphore
	go func(timeout int) {
		time.Sleep(time.Duration(timeout) * time.Second)
		semaphore = true
	}(abp.timeout)

	for i := 7; !semaphore; i++ {
		lastPosition = abp.alphabeta(position, i, NegInfty, PosInfty)
		if !semaphore {
			finalPosition = lastPosition
		}
	}
	return finalPosition
}

// Sets the maximum search depth of the algorithm.
func (abp *AlphaBetaPruning) SetSearchDepth(depth int) {
	abp.searchDepth = depth
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
func (abp *AlphaBetaPruning) alphabeta(position *board.Board, depth, alpha, beta int) *action.Action {
	moveList := position.GetMoves()
	var moveValue int

	lastAction := action.New(0, 0)

	if depth == 0 {
		moveValue = evaluator.Evaluate(position)
		lastAction.SetValue(moveValue)
		return lastAction
	}
	if moveList.Front() == nil {
		moveValue = evaluator.Evaluate(position)
		lastAction.SetValue(moveValue)
		lastAction.SetPassMove(true)
		return lastAction
	}

	var nextBoard *board.Board

	if position.GetMaxPlayer() {
		moveValue = NegInfty
		lastAction.SetValue(moveValue)
		for e := moveList.Front(); e != nil; e = e.Next() {
			if semaphore {
				return action.New(0, 0)
			}
			move := e.Value.(*action.Action)
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
		lastAction := action.New(0, 0)
		lastAction.SetValue(moveValue)
		for e := moveList.Front(); e != nil; e = e.Next() {
			if semaphore {
				return action.New(0, 0)
			}
			move := e.Value.(*action.Action)
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
