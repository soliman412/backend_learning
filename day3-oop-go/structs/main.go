package main

import "fmt"

type Car struct {
	Brand string
	Year  int
}

func main() {
	car1 := Car{Brand: "TOYOTA", Year: 2020}
	car2 := Car{"BMW", 2022}
	var car3 Car

	fmt.Println(car1)
	fmt.Println(car2.Brand)
	fmt.Println(car3)

	car3.Brand = "KIA"
	car3.Year = 2018

	fmt.Println(car3)
}
