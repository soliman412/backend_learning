from abc import ABC, abstractmethod


class Employee(ABC):
    def __init__(self, name, salary):
        self.name = name
        self.__salary = salary

    def get_salary(self):
        return self.__salary

    def give_raise(self, percent):
        if percent < 0:
            print("Invalid percent")
        else:
            self.__salary = self.__salary + self.__salary * percent / 100
            print(self.name, "got a raise of", percent, "%")

    def show_info(self):
        print(self.name, "salary:", self.__salary)

    @abstractmethod
    def calculate_pay(self):
        pass


class Manager(Employee):
    def calculate_pay(self):
        return self.get_salary() + 500


class Developer(Employee):
    def calculate_pay(self):
        return self.get_salary() + 200


manager = Manager("Sami", 1200)
developer = Developer("Rami", 1000)

employees = [manager, developer]

for employee in employees:
    print(employee.name, "earns", employee.calculate_pay())

developer.give_raise(10)
print(developer.name, "earns", developer.calculate_pay())

developer.give_raise(-5)