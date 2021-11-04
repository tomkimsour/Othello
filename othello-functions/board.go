package othellofunctions

import (
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

// func getMoves(board *Board) List{
// 	moves := list.New()

// }

func PrintHorizontalBorder() {
	fmt.Print("---")
	for i := 1; i <= BoardSize; i++ {
		fmt.Print("|---")
	}
	fmt.Println("|---")
}

func (b *Board) illustrate() {
	fmt.Print("   ")
	for i := 1; i <= BoardSize; i++ {
		fmt.Print("| " + strconv.Itoa(i) + " ")
	}
	fmt.Println("|")
	PrintHorizontalBorder()
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
		PrintHorizontalBorder()
	}
	fmt.Print("   ")
	for i := 1; i <= BoardSize; i++ {
		fmt.Print("| " + strconv.Itoa(i) + " ")
	}
	fmt.Println("|\n")
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
