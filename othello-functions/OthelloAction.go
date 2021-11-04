package othellofunctions

type Action struct {
	/** The row where the marker is placed. */
	row int

	/** The column where the marker is placed. */
	column int

	/** The estimated value of the move. */
	value int

	/** True if the player has to pass, i.e., if there is no legal move. */
	pass bool
}

/**
 * Creates a new <code>OthelloAction</code> with row <code>r</code>, column
 * <code>c</code>, and value 0.
 */
// func OthelloAction(int r, int c) {

//     row = r;
//     column = c;
//     value = 0;
// }

// func OthelloAction(int r, int c, p bool) {
//     row = r;
//     column = c;
//     value = 0;
//     pass = p
// 		row = -1
// 		column = -1
// 		value = 0
// 		pass = false
// 	}

// func OthelloAction(s string) {
//     if (s.equals("pass")) {
//         row = 0;
//         column = 0;
//         value = 0;
//         pass = true;
//     } else {
//         row = Character.getNumericValue(s.charAt(1));
//         column = Character.getNumericValue(s.charAt(3));
//         value = 0;
//     }
// }

// /** Sets the estimated value of the move. */
// public void setValue(int v) {
//     value = v;
// }

// /** Returns the estimated value of the move. */
// public int getValue() {
//     return value;
// }

// /** Sets the column where the marker is to be placed. */
// public void setColumn(int c) {
//     column = c;
// }

// /** Returns the column where the marker is to be placed. */
// public int getColumn() {
//     return column;
// }

// /** Sets the row where the marker is to be placed. */
// public void setRow(int r) {
//     row = r;
// }

// /** Returns the row where the marker is to be placed. */
// public int getRow() {
//     return row;
// }

// /**
//  * Sets the boolean that indicates whether this is a pass move. This should only
//  * be true if there are no legal moves.
//  */
// public void setPassMove(boolean b) {
//     pass = b;
// }

// /**
//  * Returns true if this is a pass move, indicating that the player has no legal
//  * moves. Otherwise returns false.
//  */
// public boolean isPassMove() {
//     return pass;
// }

// func print() {
// 	if pass {
// 		fmt.Println("pass")
// 	} else {
// 		fmt.Println("(" + row + "," + column + ")")
// 	}
// }
