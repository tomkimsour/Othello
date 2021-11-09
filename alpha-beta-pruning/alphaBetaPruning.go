package alphabetapruning

import (
	"container/list"
	"fmt"

	"github.com/tomkimsour/Othello/action"
	"github.com/tomkimsour/Othello/board"
	"github.com/tomkimsour/Othello/evaluator"
)

const PosInfty int = 2147483647
const NegInfty int = -2147483648
const MaxPlayer bool = true
const MinPlayer bool = false
const DefaultDepth int = 7

type AlphaBetaPruning struct {
	boardEvaluator evaluator.Evaluator
	searchDepth    int
	moves          list.List
	position       board.Board
}

func New(currentBoard *board.Board, depth int) *AlphaBetaPruning {
	return new(AlphaBetaPruning).Init(*currentBoard, depth)
}

func (abp *AlphaBetaPruning) Init(currentBoard board.Board, depth int) *AlphaBetaPruning {
	abp.position = currentBoard
	abp.searchDepth = depth
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
	return abp.alphabeta(position, abp.searchDepth, NegInfty, PosInfty)
}

/**
* Returns the <code>OthelloAction</code> the algorithm considers to be the
* best move.
 */
func (abp *AlphaBetaPruning) EvaluatePrintGraph(position *board.Board) *action.Action {
	return abp.alphabetaPrintGraph(position, abp.searchDepth, NegInfty, PosInfty)
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

	if depth == 0 {
		lastAction := action.New(0, 0)
		moveValue = abp.boardEvaluator.Evaluate(position)
		lastAction.SetValue(moveValue)
		return lastAction
	}
	if moveList.Front() == nil {
		lastAction := action.New(0, 0)
		moveValue = abp.boardEvaluator.Evaluate(position)
		lastAction.SetValue(moveValue)
		lastAction.SetPassMove(true)
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
			moveValue = max(moveValue, abp.alphabeta(nextBoard, depth-1, alpha, beta).GetValue())
			if moveValue > beta {
				move.SetValue(moveValue)
				return move
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
			move := e.Value.(*action.Action)
			nextBoard = position.MakeMove(move)
			evaluationValue := abp.alphabeta(nextBoard, depth-1, alpha, beta).GetValue()
			moveValue = min(moveValue, evaluationValue)
			if moveValue <= alpha {
				move.SetValue(moveValue)
				return move
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

//     if depth == 0 || node is a terminal node then
//         return the heuristic value of node
//     if maximizingPlayer then
//         value := −∞
//         for each child of node do
//             value := max(value, alphabeta(child, depth − 1, α, β, FALSE))
//             if value ≥ β then
//                 break (* β cutoff *)
//             α := max(α, value)
//         return value
//     else
//         value := +∞
//         for each child of node do
//             value := min(value, alphabeta(child, depth − 1, α, β, TRUE))
//             if value ≤ α then
//                 break (* α cutoff *)
//             β := min(β, value)
//         return value
// }

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
	if depth == 0 || moveList.Front() == nil {
		lastAction := action.New(0, 0)
		moveValue = abp.boardEvaluator.Evaluate(position)
		lastAction.SetValue(moveValue)
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
			evaluationValue := abp.alphabetaPrintGraph(nextBoard, depth-1, alpha, beta).GetValue()
			moveValue = min(moveValue, evaluationValue)
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
