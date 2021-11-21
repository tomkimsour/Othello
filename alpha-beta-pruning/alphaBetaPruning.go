package alphabetapruning

import (
	"container/list"
	"fmt"
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
	boardEvaluator evaluator.Evaluator
	searchDepth    int
	moves          list.List
	position       board.Board
	timeout        int
}

func New(currentBoard *board.Board, depth, time int) *AlphaBetaPruning {
	return new(AlphaBetaPruning).Init(*currentBoard, depth, time)
}

func (abp *AlphaBetaPruning) Init(currentBoard board.Board, depth, time int) *AlphaBetaPruning {
	abp.position = currentBoard
	abp.searchDepth = depth
	abp.timeout = time
	return abp
}

/**
* Sets the <code>OthelloEvaluator</code> the algorithm is to use for
* heuristic evaluation.
 */
func (abp *AlphaBetaPruning) SetEvaluator(evaluator evaluator.Evaluator) {
	abp.boardEvaluator = evaluator
}

/**
* Returns the <code>OthelloAction</code> the algorithm considers to be the
* best move.
 */
func (abp *AlphaBetaPruning) Evaluate(position *board.Board) *action.Action {
	var finalPosition, lastPosition *action.Action

	go func(timeout int) {
		time.Sleep(time.Duration(timeout) * time.Second)
		semaphore = true
	}(abp.timeout)

	for i := 7; !semaphore; i++ {
		lastPosition = abp.alphabeta(position, i, NegInfty, PosInfty)
		if !semaphore {
			finalPosition = lastPosition
		}
		// fmt.Printf("Search ended for depth %d\n", i)
		// finalPosition.Print()
	}
	return finalPosition
}

/** Sets the maximum search depth of the algorithm. */
func (abp *AlphaBetaPruning) SetSearchDepth(depth int) {
	abp.searchDepth = depth
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (abp *AlphaBetaPruning) alphabeta(position *board.Board, depth, alpha, beta int) *action.Action {
	moveList := position.GetMoves()
	var moveValue int

	lastAction := action.New(0, 0)

	if depth == 0 {
		moveValue = abp.boardEvaluator.Evaluate(position)
		lastAction.SetValue(moveValue)
		return lastAction
	}
	if moveList.Front() == nil {
		moveValue = abp.boardEvaluator.Evaluate(position)
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

/**
* Returns the <code>OthelloAction</code> the algorithm considers to be the
* best move.
 */
func (abp *AlphaBetaPruning) EvaluatePrintGraph(position *board.Board) *action.Action {
	var finalPosition, lastPosition *action.Action
	var indexDepthMax int

	go func(timeout int) {
		time.Sleep(time.Duration(timeout) * time.Second)
		semaphore = true
		fmt.Println("done")
	}(abp.timeout)

	for i := 2; i <= 3; i++ {
		lastPosition = abp.alphabetaPrintGraph(position, i, NegInfty, PosInfty)
		if !semaphore && indexDepthMax < i {
			finalPosition = lastPosition
		}
		fmt.Printf("Search ended for depth %d\n", i)
		// finalPosition.Print()
	}
	return finalPosition
}

// rouge max \033[31m
// blue min \033[34m
// white leaf \033[37m
// yellow alpha \033[33m
// purple beta \033[35m
// green value \033[32m
// normal \033[0m
func (abp *AlphaBetaPruning) alphabetaPrintGraph(position *board.Board, depth, alpha, beta int) *action.Action {
	moveList := position.GetMoves()
	var moveValue int
	lastAction := action.New(0, 0)

	if semaphore {
		return lastAction
	}
	if depth == 0 {
		moveValue = abp.boardEvaluator.Evaluate(position)
		lastAction.SetValue(moveValue)
		fmt.Printf("\033[37m(\033[33mα=%d\033[0m,\033[35mβ= %d\033[0m,\033[32mV=%d\033[37m)\033[0m\n", alpha, beta, moveValue)
		return lastAction
	}
	if moveList.Front() == nil {
		moveValue = abp.boardEvaluator.Evaluate(position)
		lastAction.SetValue(moveValue)
		lastAction.SetPassMove(true)
		fmt.Printf("\033[37m(\033[33mα=%d\033[0m,\033[35mβ= %d\033[0m,\033[32mV=%d\033[37m)\033[0m\n", alpha, beta, moveValue)
		return lastAction
	}

	var nextBoard *board.Board

	if position.GetMaxPlayer() {
		moveValue = NegInfty
		lastAction := action.New(0, 0)
		lastAction.SetValue(moveValue)
		for e := moveList.Front(); e != nil; e = e.Next() {
			move := e.Value.(*action.Action)
			nextBoard = position.MakeMove(move)
			if semaphore {
				return action.New(0, 0)
			}
			moveValue = max(moveValue, abp.alphabetaPrintGraph(nextBoard, depth-1, alpha, beta).GetValue())
			if moveValue > beta {
				move.SetValue(moveValue)
				for i := 0; i < depth; i++ {
					fmt.Printf("   ")
				}
				fmt.Printf("\033[31mβ cut (\033[33mα=%d\033[0m,\033[35mβ= %d\033[0m,\033[32mV=%d\033[31m)\033[0m\n", alpha, beta, moveValue)
				return move
			}
			if alpha < moveValue {
				lastAction = move
				alpha = moveValue
				for i := 0; i < depth; i++ {
					fmt.Printf("   ")
				}
				fmt.Printf("\033[31m(\033[33mα=%d\033[0m,\033[35mβ= %d\033[0m,\033[32mV=%d\033[31m)\033[0m\n", alpha, beta, moveValue)
			}
		}
		lastAction.SetValue(moveValue)
		return lastAction
	} else {
		moveValue = PosInfty
		lastAction := action.New(0, 0)
		lastAction.SetValue(moveValue)
		for e := moveList.Front(); e != nil; e = e.Next() {
			move := e.Value.(*action.Action)
			nextBoard = position.MakeMove(move)
			if semaphore {
				return action.New(0, 0)
			}
			moveValue = min(moveValue, abp.alphabetaPrintGraph(nextBoard, depth-1, alpha, beta).GetValue())
			if moveValue <= alpha {
				move.SetValue(moveValue)
				for i := 0; i < depth; i++ {
					fmt.Printf("   ")
				}
				fmt.Printf("\033[34mα cut (\033[33mα=%d\033[0m,\033[35mβ= %d\033[0m,\033[32mV=%d\033[34m)\033[0m\n", alpha, beta, moveValue)
				return move
			}
			if beta > moveValue {
				beta = moveValue
				lastAction = move
				for i := 0; i < depth; i++ {
					fmt.Printf("   ")
				}
				fmt.Printf("\033[34m(\033[33mα=%d\033[0m,\033[35mβ= %d\033[0m,\033[32mV=%d\033[34m)\033[0m\n", alpha, beta, moveValue)
			}
		}
		lastAction.SetValue(moveValue)
		return lastAction
	}
}
