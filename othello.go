package main

import (
	// "Othello/golang/OthelloPosition.go"
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
	fmt.Println(sequence)
	// othelloPosition := OthelloPosition.Ini(sequence)
	// fmt.Print(othelloPosition.toString())
}
