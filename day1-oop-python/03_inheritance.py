class BankAccount:
    def __init__(self, owner, balance):
        self.owner = owner
        self.balance = balance

    def deposit(self, amount):
        self.balance = self.balance + amount
        print("Deposited", amount)

    def withdraw(self, amount):
        if amount > self.balance:
            print("Not enough money")
        else:
            self.balance = self.balance - amount
            print("Withdrew", amount)

    def show_balance(self):
        print(self.owner, "has", self.balance)


class SavingsAccount(BankAccount):
    def add_interest(self):
        interest = self.balance * 0.1
        self.balance = self.balance + interest
        print("Interest added", interest)


class StudentAccount(BankAccount):
    def gift(self):
        self.balance = self.balance + 20
        print("Gift added", 20)


savings = SavingsAccount("sami", 1000)
savings.show_balance()
savings.deposit(200)
savings.add_interest()
savings.show_balance()

student = StudentAccount("emad", 200)
student.show_balance()
student.gift()
student.show_balance()