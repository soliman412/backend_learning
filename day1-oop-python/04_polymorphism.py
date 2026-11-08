class Animal:
    def speak(self):
        print("Some sound")


class Cat(Animal):
    def speak(self):
        print("Meow")


class Dog(Animal):
    def speak(self):
        print("Woof")


class Cow(Animal):
    def speak(self):
        print("Moo")


animals = [Cat(), Dog(), Animal(), Cow()]

for animal in animals:
    animal.speak()