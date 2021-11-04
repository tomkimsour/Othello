package othellofunctions

import (
	"fmt"
	"strconv"
)

const BoardSize int = 8

type Board struct {
	board [BoardSize + 2][BoardSize + 2]string
}

var maxPlayer bool
var boardStruct Board

func initializeBoard() Board {

	for i := 0; i < (BoardSize + 2); i++ {
		for j := 0; j < (BoardSize + 2); j++ {
			boardStruct.board[i][j] = "E"
		}
	}
	boardStruct.board[BoardSize/2][BoardSize/2] = "W"
	boardStruct.board[BoardSize/2+1][BoardSize/2+1] = "W"
	boardStruct.board[BoardSize/2][BoardSize/2+1] = "B"
	boardStruct.board[BoardSize/2+1][BoardSize/2] = "B"
	maxPlayer = true

	return boardStruct
}

func PrintHorizontalBorder() {
	fmt.Print("---")
	for i := 1; i <= BoardSize; i++ {
		fmt.Print("|---")
	}
	fmt.Println("|---")
}

func illustrate() {
	fmt.Print("   ")
	for i := 1; i <= BoardSize; i++ {
		fmt.Print("| " + strconv.Itoa(i) + " ")
	}
	fmt.Println("|")
	PrintHorizontalBorder()
	for i := 1; i <= BoardSize; i++ {
		fmt.Print(" " + strconv.Itoa(i) + " ")
		for j := 1; j <= BoardSize; j++ {
			if boardStruct.board[i][j] == "W" {
				fmt.Print("| 0 ")
			} else if boardStruct.board[i][j] == "B" {
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

func toString() string {
	s := ""
	var c, d string
	if maxPlayer {
		s += "W"
	} else {
		s += "B"
	}
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			d = boardStruct.board[i][j]
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
