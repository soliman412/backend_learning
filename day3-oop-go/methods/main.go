package main
import "fmt"

type Car struct {
	Brand string 
 	Year int 
}

func (c Car) Info() {
	fmt.Println (c.Brand, "-" , c.Year)
}
func (c Car) Age(currentYear int) int {
	return currentYear - c.Year
}
func (c Car) TryUpgrade() {
	c.Year = 2025
}
func (c *Car) Upgrade() {
	c.Year = 2025
}

func main () {
	myCar := Car {Brand: "toyota" , Year: 2020}
	myCar.Info()
	fmt.Println("Age:", myCar.Age(2026))
	myCar.TryUpgrade()
	myCar.Info()
	myCar.Upgrade()
	myCar.Info()
}