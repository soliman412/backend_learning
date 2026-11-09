package encapsulation;
public class BankAccount {
    private String owner;
    private double balance;

    public BankAccount(String owner , double balance){
        this.owner=owner;
        this.balance=balance;
    }
    public double getBalance(){
        return balance;
    }
    public void deposit (double amount){
        if (amount <= 0){
            System.out.println("Invaild amount");
            return ;
        }
        balance=balance+amount;
    }
    public void withdraw (double amount){
        if (amount > balance){
            System.out.println("not enough balance ");
            return ;
        }
        balance=balance-amount;
    }
}
