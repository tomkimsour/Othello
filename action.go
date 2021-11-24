// package action gives tools to represent actions with a board position and an heuristic value
package main

import (
	"fmt"
)

type Action struct {
	/* The row where the marker is placed. */
	row int

	/* The column where the marker is placed. */
	col int

	/* The estimated value of the move. */
	value int

	/* True if the player has to pass, i.e., if there is no legal move. */
	pass bool
}

// create a new structure Action
func NewAction(r, c int) *Action {
	return new(Action).Init(r, c)
}

// Initialise the Action structure
func (a *Action) Init(r, c int) *Action {
	a.row = r
	a.col = c
	a.pass = false
	return a
}

/* Sets the estimated value of the move. */
func (a *Action) SetValue(v int) {
	a.value = v
}

/* Returns the estimated value of the move. */
func (a *Action) GetValue() int {
	return a.value
}

/* Sets the column where the marker is to be placed. */
func (a *Action) SetColumn(c int) {
	a.col = c
}

/* Returns the column where the marker is to be placed. */
func (a *Action) GetColumn() int {
	return a.col
}

/* Sets the row where the marker is to be placed. */
func (a *Action) SetRow(r int) {
	a.row = r
}

/* Returns the row where the marker is to be placed. */
func (a *Action) GetRow() int {
	return a.row
}

// Sets the boolean that indicates whether this is a pass move. This should only
// be true if there are no legal moves.
func (a *Action) SetPassMove(b bool) {
	a.pass = b
}

// Returns true if this is a pass move, indicating that the player has no legal
// moves. Otherwise returns false.
func (a *Action) IsPassMove() bool {
	return a.pass
}

// Print the action
func (a *Action) Print() {
	if a.pass {
		fmt.Println("pass")
	} else {
		fmt.Printf("(%d,%d)\n", a.row, a.col)
	}
}
