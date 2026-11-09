package basics;
public class Car {
    String brand;
    int year;

    Car (String brand , int year){
        this.brand=brand;
        this.year=year;
    
    }
    void info() {
        System.out.println(brand + "-" + year);
    }
    public static void main(String[] args){
        Car myCar = new Car ("Toyota",2020);
        myCar.info();
    } 
}