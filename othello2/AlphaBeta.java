public class AlphaBeta implements OthelloAlgorithm {
    private OthelloEvaluator evaluator;
    private static final int maxInf = Integer.MAX_VALUE;
    private static final int minInf = Integer.MIN_VALUE;
    private int searchDepth;
    private int time = (int)System.currentTimeMillis();
    private int timeLimit = 2000;
    private boolean finished = false;

    /**
     * The evaluate method implements iterative deepening for the miniMax
     * method. It keeps track of the last finished evaluated action and
     * returns the final evaluated action.
     *
     * @param position the current board state
     * @return the suggested action
     */
    public OthelloAction evaluate(OthelloPosition position) {
        int alpha = Integer.MIN_VALUE;
        int beta = Integer.MAX_VALUE;
        OthelloAction currentAction;
        OthelloAction finalAction = new OthelloAction(0,0);
        int counter = 0;

        while(!finished && counter < searchDepth) {
            counter++;
            currentAction = miniMax(position, counter, alpha, beta);

            if (!finished) {
                finalAction = currentAction;
            }
        }

        return finalAction;
    }

    /**
     * Checks if the current position is at depth 0 or a pass move
     * and returns the static evaluation of that position if that is the case.
     * If not, call the min or max method depending on which player has the move.
     *
     * @param position the board state
     * @param depth the current depth in the minimax tree expansion
     * @param alpha alpha value for alpha-beta pruning
     * @param beta beta value for alpha-beta pruning
     * @return the static evaluation of a board state
     */
    private OthelloAction miniMax(OthelloPosition position, int depth, int alpha, int beta) {
        if (depth == 0) {
            OthelloAction action = new OthelloAction(0,0);
            action.value = evaluator.evaluate(position);
            return action;
        }

        if (position.getMoves().isEmpty()) {
            OthelloAction action = new OthelloAction(0,0);
            action.value = evaluator.evaluate(position);
            action.pass = true;
            return action;
        }

        if (position.toMove()) {
            return max(position, depth, alpha, beta);
        } else {
            return min(position, depth, alpha, beta);
        }
    }

    /**
     * Performs the minimax algorithm on the max player and recursively calls
     * miniMax. Returns immediately if the time limit is up.
     *
     * @param position the board state
     * @param depth the current depth in the minimax tree expansion
     * @param alpha alpha value for alpha-beta pruning
     * @param beta beta value for alpha-beta pruning
     * @return a dummy action if cancelled early, otherwise the evaluation for
     * the max player.
     */
    private OthelloAction max(OthelloPosition position, int depth, int alpha, int beta) {
        int currentTime;
        OthelloAction maxEval = new OthelloAction(0,0);
        maxEval.value = minInf;

        for (OthelloAction action : position.getMoves()) {
            currentTime = (int)System.currentTimeMillis();
            if (currentTime - time > timeLimit) {
                finished = true;
                return new OthelloAction(0,0);
            }

            try {
                OthelloAction a = miniMax(position.makeMove(action), depth - 1, alpha, beta);

                if (maxEval.value < a.value) {
                    maxEval = action;
                    maxEval.value = a.value;
                }

                alpha = Math.max(alpha, a.value);
                if (beta <= alpha)
                    break;
            } catch (IllegalMoveException e) {
                System.err.println(e.getAction() + " is an invalid action");
            }
        }

        return maxEval;
    }

    /**
     * Performs the minimax algorithm on the min player and recursively calls
     * miniMax. Returns immediately if the time limit is up.
     *
     * @param position the board state
     * @param depth the current depth in the minimax tree expansion
     * @param alpha alpha value for alpha-beta pruning
     * @param beta beta value for alpha-beta pruning
     * @return a dummy action if cancelled early, otherwise the evaluation for
     * the min player.
     */
    private OthelloAction min(OthelloPosition position, int depth, int alpha, int beta) {
        int currentTime;
        OthelloAction minEval = new OthelloAction(0,0);
        minEval.value = maxInf;

        for (OthelloAction action : position.getMoves()) {
            currentTime = (int)System.currentTimeMillis();
            if (currentTime - time > timeLimit) {
                finished = true;
                return new OthelloAction(0,0);
            }

            try {
                OthelloAction a = miniMax(position.makeMove(action), depth - 1, alpha, beta);

                if (minEval.value > a.value) {
                    minEval = action;
                    minEval.value = a.value;
                }

                beta = Math.min(beta, a.value);
                if (beta <= alpha)
                    break;
            } catch (IllegalMoveException e) {
                System.err.println(e.getAction() + " is an invalid action");
            }
        }

        return minEval;
    }

    public void setEvaluator(OthelloEvaluator evaluator) {
        this.evaluator = evaluator;
    }

    public void setSearchDepth(int depth) {
        searchDepth = depth;
    }

    public void setTimeLimit(int timeLimit) {
        this.timeLimit = timeLimit;
    }
}
