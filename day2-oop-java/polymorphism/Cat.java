package polymorphism;
public class Cat extends Animal {
    public Cat (String name , int age){
        super(name , age);
    }
    @Override 
    public void makesound () {
        System.out.println(name + " : meow ");
    }
}

