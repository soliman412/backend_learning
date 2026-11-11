package main

import (
	"day3/employees/company"
	"fmt"
)

func main() {
	staff := []company.Employee{
		company.NewFullTime(1, "Ahmad", 3000),
		company.NewPartTime(2, "Sara", 80, 15),
		company.NewManager(3, "Soliman", 5000, 1000),
	}

	total := 0.0

	for _, e := range staff {
		fmt.Println("#", e.ID(), e.Name(), "- Salary:", e.Salary())
		total += e.Salary()
	}

	fmt.Println("Total salaries:", total)
}