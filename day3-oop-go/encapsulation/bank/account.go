package bank
import "fmt"

type Account struct {
	Owner string
	balance float64
}

func NewAccount (owner string , balance float64) *Account {
	return &Account {Owner: owner , balance: balance}
}

func (a *Account) Balance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Invalid amount")
		return
	}
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) {
	if amount > a.balance {
	fmt.Println("not enough balance")
	return
	}
	a.balance -= amount
}

