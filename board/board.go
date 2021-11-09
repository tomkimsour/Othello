package board

import (
	"container/list"
	"fmt"
	"strconv"

	"github.com/tomkimsour/Othello/action"
)

const BoardSize int = 8

type Board struct {
	Board     [BoardSize + 2][BoardSize + 2]rune
	maxPlayer bool // True if the player is white
}

var localBoard *Board

// create a new board
func New() *Board {
	return new(Board).Init()
}

// Initialize the structure board
func (b *Board) Init() *Board {
	for i := 0; i < (BoardSize + 2); i++ {
		for j := 0; j < (BoardSize + 2); j++ {
			b.Board[i][j] = 'E'
		}
	}
	b.Board[BoardSize/2][BoardSize/2] = 'W'
	b.Board[BoardSize/2+1][BoardSize/2+1] = 'W'
	b.Board[BoardSize/2][BoardSize/2+1] = 'B'
	b.Board[BoardSize/2+1][BoardSize/2] = 'B'
	b.maxPlayer = true
	return b
}

// Initialize the board accordind to the string input
func (b *Board) OthelloPosition(s string) *Board {
	if len(s) != 65 {
		b.Init()
	} else {
		if s[0] == 'W' {
			b.maxPlayer = true
		} else {
			b.maxPlayer = false
		}
		for i := 1; i <= 64; i++ {
			var c rune
			if s[i] == 'E' {
				c = 'E'
			} else if s[i] == 'O' {
				c = 'W'
			} else {
				c = 'B'
			}
			column := ((i - 1) % 8) + 1
			row := (i-1)/8 + 1
			b.Board[row][column] = c
		}
	}
	return b
}

// return which player turn it is
// true is white and false black
func (b *Board) GetMaxPlayer() bool {
	return b.maxPlayer
}

func (b *Board) ChangeMaxPlayer() {
	b.maxPlayer = !b.maxPlayer
}

// take the x,y position of a cell
// return a boolean corresponding to the presence of neighbours
func hasNeighbor(row int, column int) bool {
	if !isFree(row-1, column) {
		return true
	}
	if !isFree(row-1, column+1) {
		return true
	}
	if !isFree(row, column+1) {
		return true
	}
	if !isFree(row+1, column+1) {
		return true
	}
	if !isFree(row+1, column) {
		return true
	}
	if !isFree(row+1, column-1) {
		return true
	}
	if !isFree(row, column-1) {
		return true
	}
	if !isFree(row-1, column-1) {
		return true
	}
	return false
}

// return if the cell is free
func isFree(row int, col int) bool {
	if localBoard.Board[row][col] == 'E' {
		return true
	}
	return false
}

// return if a piece can be placed on this cell
func isCandidate(row int, col int) bool {
	if isFree(row, col) && hasNeighbor(row, col) {
		return true
	}
	return false
}

// return if the cell is own by the opponent
func isOpponentSquare(row int, column int) bool {
	if localBoard.maxPlayer && (localBoard.Board[row][column] == 'B') {
		return true
	}
	if !localBoard.maxPlayer && (localBoard.Board[row][column] == 'W') {
		return true
	}
	return false
}

//return if the cell is own by the current player
func isOwnSquare(row, column int) bool {
	if !localBoard.maxPlayer && (localBoard.Board[row][column] == 'B') {
		return true
	}
	if localBoard.maxPlayer && (localBoard.Board[row][column] == 'W') {
		return true
	}
	return false
}

func checkNorth(row, column int) bool {
	if !isOpponentSquare(row-1, column) {
		return false
	}
	for i := (row - 2); i > 0; i-- {
		if isFree(i, column) {
			return false
		}
		if isOwnSquare(i, column) {
			return true
		}
	}
	return false
}

func checkEast(row, col int) bool {
	if !isOpponentSquare(row, col+1) {
		return false
	}
	for i := col + 2; i <= BoardSize; i++ {
		if isFree(row, i) {
			return false
		}
		if isOwnSquare(row, i) {
			return true
		}
	}
	return false
}

func checkSouth(row, col int) bool {
	if !isOpponentSquare(row+1, col) {
		return false
	}
	for i := row + 2; i <= BoardSize; i++ {
		if isFree(i, col) {
			return false
		}
		if isOwnSquare(i, col) {
			return true
		}
	}
	return false
}

func checkWest(row, col int) bool {
	if !isOpponentSquare(row, col-1) {
		return false
	}
	for i := col - 2; i > 0; i-- {
		if isFree(row, i) {
			return false
		}
		if isOwnSquare(row, i) {
			return true
		}
	}
	return false
}

func checkNorthEast(row, col int) bool {
	if !isOpponentSquare(row-1, col+1) {
		return false
	}
	for i := 2; (row-i) > 0 && (col+i) <= BoardSize; i++ {
		if isFree(row-i, col+i) {
			return false
		}
		if isOwnSquare(row-i, col+i) {
			return true
		}
	}
	return false
}

func checkSouthEast(row, col int) bool {
	if !isOpponentSquare(row+1, col+1) {
		return false
	}
	for i := 2; row+1 <= BoardSize && col+i <= BoardSize; i++ {
		if isFree(row+i, col+i) {
			return false
		}
		if isOwnSquare(row+i, col+i) {
			return true
		}
	}
	return false
}
func checkSouthWest(row, col int) bool {
	if !isOpponentSquare(row+1, col-1) {
		return false
	}
	for i := 2; row+i <= BoardSize && col-i > 0; i++ {
		if isFree(row+i, col-i) {
			return false
		}
		if isOwnSquare(row+i, col-i) {
			return true
		}
	}
	return false
}
func checkNorthWest(row, col int) bool {
	if !isOpponentSquare(row-1, col-1) {
		return false
	}
	for i := 2; row-i > 0 && col-i > 0; i++ {
		if isFree(row-i, col-i) {
			return false
		}
		if isOwnSquare(row-i, col-i) {
			return true
		}
	}
	return false
}

// check if the move is a valid move
func isMove(row int, col int) bool {
	if checkNorth(row, col) {
		return true
	}
	if checkNorthEast(row, col) {
		return true
	}
	if checkEast(row, col) {
		return true
	}
	if checkSouthEast(row, col) {
		return true
	}
	if checkSouth(row, col) {
		return true
	}
	if checkSouthWest(row, col) {
		return true
	}
	if checkWest(row, col) {
		return true
	}
	if checkNorthWest(row, col) {
		return true
	}

	return false

}

// get all the possible moves starting from a board
// return a linked list of the moves
func (b *Board) GetMoves() *list.List {
	localBoard = b
	moves := list.New()

	var candidates [BoardSize + 2][BoardSize + 2]bool
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			candidates[i][j] = isCandidate(i+1, j+1)
		}
	}
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if candidates[i][j] {
				if isMove(i+1, j+1) {
					moves.PushBack(action.New(i+1, j+1))
				}
			}
		}
	}
	return moves
}

func printHorizontalBorder() {
	fmt.Print("---")
	for i := 1; i <= BoardSize; i++ {
		fmt.Print("|---")
	}
	fmt.Println("|---")
}

// display in the console the state of the game
func (b *Board) Print() {
	fmt.Print("   ")
	for i := 1; i <= BoardSize; i++ {
		fmt.Print("| " + strconv.Itoa(i) + " ")
	}
	fmt.Println("|")
	printHorizontalBorder()
	for i := 1; i <= BoardSize; i++ {
		fmt.Print(" " + strconv.Itoa(i) + " ")
		for j := 1; j <= BoardSize; j++ {
			if b.Board[i][j] == 'W' {
				fmt.Print("| 0 ")
			} else if b.Board[i][j] == 'B' {
				fmt.Print("| X ")
			} else {
				fmt.Print("|   ")
			}
		}
		fmt.Println("| " + strconv.Itoa(i) + " ")
		printHorizontalBorder()
	}
	fmt.Print("   ")
	for i := 1; i <= BoardSize; i++ {
		fmt.Print("| " + strconv.Itoa(i) + " ")
	}
	fmt.Println("|")
	fmt.Println()
}

// put the board in a string format
func (b *Board) ToString() string {
	s := ""
	var c, d rune
	if b.maxPlayer {
		s += "W"
	} else {
		s += "B"
	}
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			d = b.Board[i][j]
			if d == 'W' {
				c = 'O'
			} else if d == 'B' {
				c = 'X'
			} else {
				c = 'E'
			}
			s += string(c)
		}
	}
	return s
}

// copy the board structure
func (b *Board) Clone() *Board {
	var newPosition *Board = New()
	newPosition.maxPlayer = b.maxPlayer
	for i := 0; i < BoardSize+2; i++ {
		for j := 0; j < BoardSize+2; j++ {
			newPosition.Board[i][j] = b.Board[i][j]
		}
	}
	return newPosition
}

// change the color of an own cell
func (b *Board) flipSquare(row, col int) {
	if b.maxPlayer {
		b.Board[row][col] = 'W'
	} else {
		b.Board[row][col] = 'B'
	}
}

// play an action on the board
// return the new board with the played action
func (b *Board) MakeMove(action *action.Action) *Board {
	localBoard = b.Clone()

	rowAction := action.GetRow()
	colAction := action.GetColumn()

	if action.IsPassMove() {
		localBoard.maxPlayer = !localBoard.maxPlayer
		return localBoard
	} else {
		// put the piece on the Board
		if localBoard.maxPlayer {
			localBoard.Board[rowAction][colAction] = 'W'
		} else {
			localBoard.Board[rowAction][colAction] = 'B'
		}

		// flip all the coins between the new coin and the existing one
		if checkNorth(rowAction, colAction) {
			for rowIndex := rowAction - 1; isOpponentSquare(rowIndex, colAction); rowIndex-- {
				localBoard.flipSquare(rowIndex, colAction)
			}
		}
		if checkEast(rowAction, colAction) {
			for colIndex := colAction + 1; isOpponentSquare(rowAction, colIndex); colIndex++ {
				localBoard.flipSquare(rowAction, colIndex)
			}
		}
		if checkSouth(rowAction, colAction) {
			for rowIndex := rowAction + 1; isOpponentSquare(rowIndex, colAction); rowIndex++ {
				localBoard.flipSquare(rowIndex, colAction)
			}
		}
		if checkWest(rowAction, colAction) {
			for colIndex := colAction - 1; isOpponentSquare(rowAction, colIndex); colIndex-- {
				localBoard.flipSquare(rowAction, colIndex)
			}
		}

		if checkNorthEast(rowAction, colAction) {
			rowIndex := rowAction - 1
			for colIndex := colAction + 1; isOpponentSquare(rowIndex, colIndex); colIndex++ {
				localBoard.flipSquare(rowIndex, colIndex)
				rowIndex--
			}
		}
		if checkNorthWest(rowAction, colAction) {
			rowIndex := rowAction - 1
			for colIndex := colAction - 1; isOpponentSquare(rowIndex, colIndex); colIndex-- {
				localBoard.flipSquare(rowIndex, colIndex)
				rowIndex--
			}

		}
		if checkSouthWest(rowAction, colAction) {
			rowIndex := rowAction + 1
			for colIndex := colAction - 1; isOpponentSquare(rowIndex, colIndex); colIndex-- {
				localBoard.flipSquare(rowIndex, colIndex)
				rowIndex++
			}

		}
		if checkSouthEast(rowAction, colAction) {
			rowIndex := rowAction + 1
			for colIndex := colAction + 1; isOpponentSquare(rowIndex, colIndex); colIndex++ {
				localBoard.flipSquare(rowIndex, colIndex)
				rowIndex++
			}

		}
		localBoard.ChangeMaxPlayer()
		return localBoard
	}
}
