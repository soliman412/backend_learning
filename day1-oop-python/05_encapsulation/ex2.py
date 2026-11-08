class Player():
    def __init__(self, name):
        self.name = name
        self.__health = 100

    def take_damage(self, amount):
        self.health = self.__health - amount
        if self.__health < 0:
            self.__health = 0
        print(self.name, "took", amount, "damage")

    def show_health(self):
        print(self.name, "health:", self.__health)


player = Player("soliman")
player.show_health()
player.take_damage(30)
player.show_health()
player.take_damage(100)
player.show_health()