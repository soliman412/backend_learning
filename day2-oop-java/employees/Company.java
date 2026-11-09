package employees;

public class Company {
    public static void main(String[] args) {
        Employee[] employees = {
            new FullTimeEmployee(1, "Ahmad", 3000),
            new PartTimeEmployee(2, "Sara", 80, 15),
            new Manager(3, "Soliman", 5000, 1000)
        };

        double total = 0;

        for (Employee e : employees) {
            e.printInfo();
            total = total + e.calculateSalary();
        }

        System.out.println("Total salaries: " + total);
    }
}