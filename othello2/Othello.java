import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Othello {
    public static void main(String[] args) {
        Pattern p = Pattern.compile("[WB][OXE]{64}");

        /* Error handling */

        if (args.length < 2) {
            System.err.println("Too few arguments. Expected: position time_limit");
            return;
        } else if (args[0].length() < 65) {
            System.err.println("Position string too short. Expected length: 65");
            return;
        } else if (args[0].length() > 65) {
            System.err.println("Position string too long. Expected length: 65");
            return;
        }

        Matcher m = p.matcher(args[0]);

        if (!m.matches()) {
            System.err.println("Position string format error.");
            return;
        }

        try {
            Integer.parseInt(args[1]);
        } catch (Exception e) {
            System.err.println("Time format error.");
            return;
        }

        /* Create a position and evaluate an action */

        OthelloPosition pos = new OthelloPosition(args[0]);
        AlphaBeta ab = new AlphaBeta();
        OthelloEvaluator ev = new ComponentWiseHeuristic();
        ab.setSearchDepth(12);
        ab.setEvaluator(ev);
        ab.setTimeLimit(Integer.parseInt(args[1]) * 1000);

        OthelloAction a;
        a = ab.evaluate(pos);
        a.print();
    }
}
