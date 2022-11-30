public class ComponentWiseHeuristic implements OthelloEvaluator {

    public int evaluate(OthelloPosition position) {
        int c = coinParity(position);
        int m = mobility(position);
        int cc = cornersCaptured(position);

        return 25 * c + 5 * m + 30 * cc;
    }

    /**
     * Calculates the coin parity for the current board state
     *
     * @param position the current board state
     * @return an evaluation of the coin parity scaled between -100 and 100
     */
    private int coinParity(OthelloPosition position) {
        int maxPlayer = 0;
        int minPlayer = 0;
        int cp = 0;

        for (int i = 0; i < position.BOARD_SIZE + 2; i++) {
            for (int j = 0; j < position.BOARD_SIZE + 2; j++) {
                if (position.board[i][j] == 'W') {
                    maxPlayer++;
                }

                if (position.board[i][j] == 'B') {
                    minPlayer++;
                }
            }
        }

        if (maxPlayer + minPlayer != 0) {
            cp = 100 * (maxPlayer - minPlayer) / (maxPlayer + minPlayer);
        }

        return cp;
    }

    /**
     * Calculates the mobility (amount of moves available) for the current
     * board state
     *
     * @param position the current board state
     * @return an evaluation of the mobility scaled between -100 and 100
     */
    private int mobility(OthelloPosition position) {
        int maxPlayer = position.getMoves().size();
        position.maxPlayer = !position.toMove(); // Check move list of opponent
        int minPlayer = position.getMoves().size();
        position.maxPlayer = !position.toMove(); // Flip back

        int mobility = 0;

        if (maxPlayer + minPlayer != 0) {
            mobility = 100 * (maxPlayer - minPlayer) / (maxPlayer + minPlayer);
        }

        return mobility;
    }

    /**
     * Calculates the amount of corners captured between the two players for
     * the current board state
     *
     * @param position the current board state
     * @return an evaluation of the corners captured scaled between -100 and 100
     */
    private int cornersCaptured(OthelloPosition position) {
        int maxPlayer = 0;
        int minPlayer = 0;
        int corners = 0;

        if (position.board[1][1] == 'W')
            maxPlayer++;
        if (position.board[1][8] == 'W')
            maxPlayer++;
        if (position.board[8][1] == 'W')
            maxPlayer++;
        if (position.board[8][8] == 'W')
            maxPlayer++;

        if (position.board[1][1] == 'B')
            minPlayer++;
        if (position.board[1][8] == 'B')
            minPlayer++;
        if (position.board[8][1] == 'B')
            minPlayer++;
        if (position.board[8][8] == 'B')
            minPlayer++;

        if (maxPlayer + minPlayer != 0) {
            corners = 100 * (maxPlayer - minPlayer) / (maxPlayer + minPlayer);
        }

        return corners;
    }
}