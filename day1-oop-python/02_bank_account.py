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


account = BankAccount("soliman", 100)
account.show_balance()
account.deposit(50)
account.show_balance()
account.withdraw(100)
account.show_balance()
account.withdraw(300)