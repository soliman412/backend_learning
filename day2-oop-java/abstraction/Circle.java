package abstraction;

public class Circle extends Shape implements Drawable,
Resizable {
    private double radius;

    public Circle(double radius) {
        super ("Circle");
        this.radius = radius;
    }
    @Override 
    public double area() {
        return Math.PI * radius * radius;
    }
    @Override 
    public void draw(){
        System.out.println("Drawing a circle");
    }
    @Override 
    public void resize (double factor){
        radius =  radius * factor;
    }
}
