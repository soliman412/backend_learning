package polymorphism;

public class Zoo {
     public static void main(String[] args){
        Animal[] animals = {
            new Dog("Rex" , 3 , "Husky"),
            new Cat("Mimi" , 2 ),
            new Animal ("Unknown" , 5)
        };

       for (Animal a : animals){
        a.makesound();
        }
    }
}