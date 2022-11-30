public class DefaultEvaluator implements OthelloEvaluator {
    public int evaluate(OthelloPosition position) {
        int maxPlayer = 0;
        int minPlayer = 0;

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

        return maxPlayer - minPlayer;
    }
}
