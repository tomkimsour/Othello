package othellofunctions

import (
	"container/list"
	"fmt"
	"strconv"
)

const BoardSize int = 8

type Board struct {
	board     [BoardSize + 2][BoardSize + 2]string
	maxPlayer bool
}

func New() *Board {
	return new(Board).Init()
}

func (b *Board) Init() *Board {
	for i := 0; i < (BoardSize + 2); i++ {
		for j := 0; j < (BoardSize + 2); j++ {
			b.board[i][j] = "E"
		}
	}
	b.board[BoardSize/2][BoardSize/2] = "W"
	b.board[BoardSize/2+1][BoardSize/2+1] = "W"
	b.board[BoardSize/2][BoardSize/2+1] = "B"
	b.board[BoardSize/2+1][BoardSize/2] = "B"
	b.maxPlayer = true
	return b
}

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
			var c string
			if s[i] == 'E' {
				c = "E"
			} else if s[i] == 'O' {
				c = "W"
			} else {
				c = "B"
			}
			column := ((i - 1) % 8) + 1
			row := (i-1)/8 + 1
			b.board[row][column] = c
		}
	}
	return b
}

func (b *Board) hasNeighbor(row int, column int) bool {
	if !b.isFree(row-1, column) {
		return true
	}
	if !b.isFree(row-1, column+1) {
		return true
	}
	if !b.isFree(row, column+1) {
		return true
	}
	if !b.isFree(row+1, column+1) {
		return true
	}
	if !b.isFree(row+1, column) {
		return true
	}
	if !b.isFree(row+1, column-1) {
		return true
	}
	if !b.isFree(row, column-1) {
		return true
	}
	if !b.isFree(row-1, column-1) {
		return true
	}
	return false
}

func (b *Board) isFree(row int, col int) bool {
	if b.board[row][col] == "E" {
		return true
	}
	return false
}

func (b *Board) isCandidate(row int, col int) bool {
	if b.isFree(row, col) && b.hasNeighbor(row, col) {
		return true
	}
	return false
}

func (b *Board) getMoves() *list.List {
	moves := list.New()

	var candidates [BoardSize + 2][BoardSize + 2]bool
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			candidates[i][j] = b.isCandidate(i+1, j+1)
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
			if b.board[i][j] == "W" {
				fmt.Print("| 0 ")
			} else if b.board[i][j] == "B" {
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

func (b *Board) toString() string {
	s := ""
	var c, d string
	if b.maxPlayer {
		s += "W"
	} else {
		s += "B"
	}
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			d = b.board[i][j]
			if d == "W" {
				c = "O"
			} else if d == "B" {
				c = "X"
			} else {
				c = "E"
			}
			s += c
		}
	}
	return s
}
