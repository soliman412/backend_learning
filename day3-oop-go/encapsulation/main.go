package main

import (
	"day3/encapsulation/bank"
	"fmt"
)

func main() {
	acc := bank.NewAccount("Soliman", 500)

	acc.Deposit(200)
	acc.Withdraw(1000)
	acc.Deposit(-50)
	

	fmt.Println("Owner:", acc.Owner)
	fmt.Println("Balance:", acc.Balance())
}