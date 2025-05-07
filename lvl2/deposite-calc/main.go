package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	balance float64
}

func (b BankAccount) Balance() float64 {
	return b.balance
}

func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
}

func (b *BankAccount) Withdraw(amount float64) error {
	if b.balance < amount {
		return errors.New("not enought money")
	}

	b.balance -= amount

	return nil
}

func main() {
	account := BankAccount{}

	account.Deposit(100)
	account.Deposit(100)
	account.Deposit(500)
	fmt.Println(account.Balance())

	err := account.Withdraw(500)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account.Balance())

	err = account.Withdraw(500)

	if err != nil {
		fmt.Println(err)
	}
}
