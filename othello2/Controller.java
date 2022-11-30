import javax.swing.*;
import java.awt.event.MouseAdapter;
import java.awt.event.MouseEvent;
import java.lang.reflect.InvocationTargetException;

public class Controller {
    private GUI gui;
    private OthelloPosition position;

    private Controller(GUI gui) {
        this.gui = gui;
    }

    private void initialize() {
        try {
            SwingUtilities.invokeAndWait(() -> {
                gui.initialize();
                gui.display();
                gui.addMouseListener(new MouseClick());
            });

            position = new OthelloPosition("WEEEEEEEEEEEEEEEEEEEEEEEEEEEEEOEEOXOXOXEEEEEEEEEEEEEEEEEEEEEEEEEE");
            position.initialize();
            position.illustrate();

            gui.updateBoard(position.board);

        } catch (InterruptedException | InvocationTargetException e) {
            throw new RuntimeException("Initialization error: " + e.getMessage());
        }
    }

    public static void main(String args[]) {
        GUI gui = new GUI();
        Controller controller = new Controller(gui);
        controller.initialize();
    }

    private class MouseClick extends MouseAdapter {
        @Override
        public void mouseClicked(MouseEvent e) {
            int col = (e.getX() - 60) / 85 + 1;
            int row = (e.getY() - 60) / 85 + 1;
            OthelloAction a = new OthelloAction(row, col);
            OthelloPosition newPos = null;
            try {
                newPos = position.makeMove(a);
            }
            catch (IllegalMoveException ex) {
                System.out.println("Not valid move");
            }
            if (newPos != null) {
                position = newPos;
                gui.updateBoard(newPos.board);
            }
        }
    }
}
