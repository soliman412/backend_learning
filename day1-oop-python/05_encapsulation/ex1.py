class BankAccount:
    def __init__(self, owner, balance):
        self.owner = owner
        self.__balance = balance

    def deposit(self, amount):
        if amount > 0:
            self.balance = self.__balance + amount
            print("Deposited", amount)
        else:
            print("Invalid amount")

    def withdraw(self, amount):
        if amount > self.__balance:
            print("Not enough money")
        else:
            self.balance = self.__balance - amount
            print("Withdrew", amount)

    def show_balance(self):
        print(self.owner, "has", self.__balance)


account = BankAccount("soliman", 100)
account.deposit(50)
account.show_balance()
account.deposit(-20)

account.__balance = 1000000
account.show_balance()