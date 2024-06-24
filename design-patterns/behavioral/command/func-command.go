package main

import "fmt"

type BankAccount3 struct {
	Balance int
}

func Deposit3(ba *BankAccount3, amount int) {
	fmt.Println("Depositing", amount)
	ba.Balance += amount
}

func Withdraw3(ba *BankAccount3, amount int) {
	if ba.Balance >= amount {
		fmt.Println("Withdrawing", amount)
		ba.Balance -= amount
	}
}

func main3() {
	ba := &BankAccount3{0}
	var commands []func()
	commands = append(commands, func() {
		Deposit3(ba, 100)
	})
	commands = append(commands, func() {
		Withdraw3(ba, 100)
	})

	for _, cmd := range commands {
		cmd()
	}
}
