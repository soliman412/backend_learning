package encapsulation;
public class Main {
    public static void main (String[] args){
        BankAccount acc = new BankAccount("soliman", 500);

        acc.deposit(200);
        acc.withdraw(1000);
        acc.deposit(-50);

        System.out.println("Balance :" + acc.getBalance());
    }
}
