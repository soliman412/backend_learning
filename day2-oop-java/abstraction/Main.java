package abstraction;

public class Main {
    public static void main(String[] args) {
        Shape[] shapes = {
            new Circle(2),
            new Rectangle(3, 4)
        };
        for (Shape s : shapes) {
            s.describe();
        }

        Drawable[] drawables = {
            new Circle(1),
            new Rectangle(2, 2)
        };
        for (Drawable d : drawables) {
            d.draw();
        }

        Circle c = new Circle(1);
        c.resize(3);
        c.describe();
    }
}