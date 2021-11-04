package main

import (
	"fmt"
	"os"
)

func main() {
	var sequence string
	if len(os.Args) > 1 {
		sequence = os.Args[1]
	} else {
		sequence = "WEEEEEEEEEEEEEEEEEEEEEEEEEEEOXEEEEEEXOEEEEEEEEEEEEEEEEEEEEEEEEEEE"
	}
	board := Board.New()
	fmt.Println(sequence)
	// othelloPosition := OthelloPosition.Ini(sequence)
	// fmt.Print(othelloPosition.toString())
}
