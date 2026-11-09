package employees;

public class Manager extends FullTimeEmployee implements Bonusable {
    private double bonus;

    public Manager(int id, String name, double monthlySalary, double bonus) {
        super(id, name, monthlySalary);
        this.bonus = bonus;
    }

    @Override
    public double getBonus() {
        return bonus;
    }

    @Override
    public double calculateSalary() {
        return super.calculateSalary() + bonus;
    }
}