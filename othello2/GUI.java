
import javax.swing.*;
import java.awt.*;
import java.awt.event.MouseAdapter;
import java.util.ArrayList;

public class GUI {
    private JFrame frame;
    private GamePanel panel;

    public void initialize() {
        frame = new JFrame();
        frame.setLayout(new BorderLayout());
        frame.setPreferredSize(new Dimension(800,850));
        frame.setMinimumSize(new Dimension(800, 850));

        initializeMenuBar();

        initializeGamePanel();
        frame.add(panel, BorderLayout.CENTER);

        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.pack();
    }

    private void initializeGamePanel() {
        panel = new GamePanel();
        panel.redraw();
    }

    private void initializeMenuBar() {
        JMenuBar menuBar = new JMenuBar();

        JMenu menu = new JMenu("File");
        menuBar.add(menu);

        JMenuItem newGame = new JMenuItem("New Game");
        menu.add(newGame);
        JMenuItem quit = new JMenuItem("Quit");
        quit.addActionListener((e)->System.exit(0));
        menu.add(quit);

        frame.setJMenuBar(menuBar);
    }

    public void display() {
        frame.setVisible(true);
    }

    public void addMouseListener(MouseAdapter listener) {
        panel.addMouseListener(listener);
    }

    public void updateBoard(char[][] board) {
        this.panel.updateBoard(board);
    }
}
