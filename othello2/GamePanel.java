import javax.swing.*;
import java.awt.*;
import java.awt.geom.Ellipse2D;
import java.awt.image.BufferedImage;
import java.util.ArrayList;

public class GamePanel extends JPanel {
    private BufferedImage foreGroundImage;
    private BufferedImage backgroundImage;
    private final int width = 800;
    private final int height = 800;

    private final int offset = 60;
    private final int size = 85;

    public GamePanel() {
        super();
        setPreferredSize(new Dimension(width,height));
        setMinimumSize(new Dimension(width, height));
        setUpBackground();
    }

    private void setUpBackground() {
        backgroundImage = new BufferedImage(
                width, height, BufferedImage.TYPE_INT_ARGB);
        Graphics2D g2d = (Graphics2D) backgroundImage.getGraphics();
        Color color = new Color(6, 99, 60);
        g2d.setColor(color);
        g2d.fillRect(0, 0, width, height);
        g2d.setColor(Color.BLACK);
        for (int i = 0; i < 9; i++) {
            g2d.drawLine(i * size + offset, offset, i * size + offset, size * 8 + offset);
            g2d.drawLine(offset, i * size + offset, size * 8 + offset, i * size + offset);
        }

    }

    public void paintComponent(Graphics g) {
        super.paintComponent(g);
        g.drawImage(backgroundImage, 0, 0, null);
        g.drawImage(foreGroundImage, 0, 0, null);
    }

    public void redraw() {
        revalidate();
        repaint();
    }

    public void updateGraphics(BufferedImage image) {
        this.foreGroundImage = image;
    }

    public void updateBoard(char[][] board) {
        BufferedImage image = new BufferedImage(
                width, height, BufferedImage.TYPE_INT_ARGB);
        Graphics2D g2d = (Graphics2D) image.getGraphics();


       for (int i = 0; i < 8; i++) {
           for (int j = 0; j < 8; j++) {
               if (board[i+1][j+1] != 'E') {
                   if (board[i+1][j+1] == 'W') {
                       g2d.setColor(Color.WHITE);
                   }
                   else if (board[i+1][j+1] == 'B') {
                       g2d.setColor(Color.BLACK);
                   }
                   Ellipse2D.Double circle = new Ellipse2D.Double(offset + size * j + 5, offset + size * i + 5, size - 10, size - 10);
                   g2d.fill(circle);
               }

            }
        }

        updateGraphics(image);
        redraw();
    }
}
