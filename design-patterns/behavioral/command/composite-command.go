package main

import "fmt"

var overdraftLimit2 = -500

type BankAccount2 struct {
	balance int
}

func (b *BankAccount2) Deposit2(amount int) {
	b.balance += amount
	fmt.Println("Deposit2ed", amount,
		"\b, balance is now", b.balance)
}

func (b *BankAccount2) Withdraw2(amount int) bool {
	if b.balance-amount >= overdraftLimit2 {
		b.balance -= amount
		fmt.Println("Withdrew", amount,
			"\b, balance is now", b.balance)
		return true
	}
	return false
}

type Command2 interface {
	Call()
	Undo()
	Succeeded() bool
	SetSucceeded(value bool)
}

type Action2 int

const (
	Deposit2 Action2 = iota
	Withdraw2
)

type BankAccount2Command2 struct {
	account   *BankAccount2
	action    Action2
	amount    int
	succeeded bool
}

func (b *BankAccount2Command2) SetSucceeded(value bool) {
	b.succeeded = value
}

// additional member
func (b *BankAccount2Command2) Succeeded() bool {
	return b.succeeded
}

func (b *BankAccount2Command2) Call() {
	switch b.action {
	case Deposit2:
		b.account.Deposit2(b.amount)
		b.succeeded = true
	case Withdraw2:
		b.succeeded = b.account.Withdraw2(b.amount)
	}
}

func (b *BankAccount2Command2) Undo() {
	if !b.succeeded {
		return
	}
	switch b.action {
	case Deposit2:
		b.account.Withdraw2(b.amount)
	case Withdraw2:
		b.account.Deposit2(b.amount)
	}
}

type CompositeBankAccount2Command2 struct {
	command2s []Command2
}

func (c *CompositeBankAccount2Command2) Succeeded() bool {
	for _, cmd := range c.command2s {
		if !cmd.Succeeded() {
			return false
		}
	}
	return true
}

func (c *CompositeBankAccount2Command2) SetSucceeded(value bool) {
	for _, cmd := range c.command2s {
		cmd.SetSucceeded(value)
	}
}

func (c *CompositeBankAccount2Command2) Call() {
	for _, cmd := range c.command2s {
		cmd.Call()
	}
}

func (c *CompositeBankAccount2Command2) Undo() {
	// undo in reverse order
	for idx := range c.command2s {
		c.command2s[len(c.command2s)-idx-1].Undo()
	}
}

func NewBankAccount2Command2(account *BankAccount2, action Action2, amount int) *BankAccount2Command2 {
	return &BankAccount2Command2{account: account, action: action, amount: amount}
}

type MoneyTransferCommand2 struct {
	CompositeBankAccount2Command2
	from, to *BankAccount2
	amount   int
}

func NewMoneyTransferCommand2(from *BankAccount2, to *BankAccount2, amount int) *MoneyTransferCommand2 {
	c := &MoneyTransferCommand2{from: from, to: to, amount: amount}
	c.command2s = append(c.command2s,
		NewBankAccount2Command2(from, Withdraw2, amount))
	c.command2s = append(c.command2s,
		NewBankAccount2Command2(to, Deposit2, amount))
	return c
}

func (m *MoneyTransferCommand2) Call() {
	ok := true
	for _, cmd := range m.command2s {
		if ok {
			cmd.Call()
			ok = cmd.Succeeded()
		} else {
			cmd.SetSucceeded(false)
		}
	}
}

func main2() {
	ba := &BankAccount2{}
	cmdDeposit2 := NewBankAccount2Command2(ba, Deposit2, 100)
	cmdWithdraw2 := NewBankAccount2Command2(ba, Withdraw2, 1000)
	cmdDeposit2.Call()
	cmdWithdraw2.Call()
	fmt.Println(ba)
	cmdWithdraw2.Undo()
	cmdDeposit2.Undo()
	fmt.Println(ba)

	from := BankAccount2{100}
	to := BankAccount2{0}
	mtc := NewMoneyTransferCommand2(&from, &to, 100) // try 1000
	mtc.Call()

	fmt.Println("from=", from, "to=", to)

	fmt.Println("Undoing...")
	mtc.Undo()
	fmt.Println("from=", from, "to=", to)
}
