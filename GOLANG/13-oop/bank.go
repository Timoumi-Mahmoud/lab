package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrDraftIncurred       = errors.New("overdraft incurred")
)

type Depositable interface {
	Deposit(float64)
}
type Withdrawable interface {
	Withdraw(float64) (float64, error)
}

type BankAccount struct {
	Owner   string
	balance float64
}

type DraftBankAccount struct {
	BankAccount
	Fee float64
}

func (ba BankAccount) Balance() float64 {
	return ba.balance
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.balance += amount
}

func (ba *BankAccount) Withdraw(amount float64) (float64, error) {
	if ba.balance < amount {
		return 0, ErrInsufficientBalance
	}
	ba.balance -= amount
	return ba.balance, nil

}
func (dba *DraftBankAccount) Withdraw(amount float64) (float64, error) {
	if dba.balance < (amount + dba.Fee) {
		dba.balance -= dba.Fee
		return 0, ErrDraftIncurred
	}

	dba.balance -= dba.Fee
	dba.balance -= amount
	return dba.balance, nil

}

func Transfer(debtor Withdrawable, creditor Depositable, amount float64) error {
	balance, err := debtor.Withdraw(amount)
	switch err {
	case ErrInsufficientBalance:
		return err
	case ErrDraftIncurred:
		log.Printf("debtor incurred overdraft, new balance is %.2f", balance)
	}
	creditor.Deposit(amount)
	return nil
}

func main() {
	b1 := &BankAccount{"bara", 50}

	b1.Deposit(20)
	fmt.Println(b1.Balance())

	balance, err := b1.Withdraw(150)
	if err != nil {
		log.Println("ERROR: ", err)
	}

	fmt.Println("balance: ", balance)

	a1 := &DraftBankAccount{BankAccount{"Khadija", 90}, 10}
	fmt.Println(a1.BankAccount.Owner)
	balancedpa, err := a1.Withdraw(50)
	if err != nil {
		log.Println("ERROR: ", err)
	}

	fmt.Println("balance: ", balancedpa)
	a1.Deposit(1000)

	fmt.Println(a1.Balance())
	err = Transfer(b1, a1, 90)
	if err != nil {
		log.Println("ERROR: ", err)
	}
	fmt.Println()
	fmt.Println("before: ", b1.Balance(), a1.Balance())
	b1.Deposit(100)
	fmt.Println("Account 1 Balance: ", b1.Balance(), " , Account 2 balance: ", a1.Balance())

	err = Transfer(b1, a1, 90)
	if err != nil {
		log.Println("ERROR: ", err)
	}

	fmt.Println("Account 1 Balance: ", b1.Balance(), " , Account 2 balance: ", a1.Balance())
	err = Transfer(b1, a1, 90)
	if err != nil {
		log.Println("ERROR: ", err)
	}

	fmt.Println("Account 1 Balance: ", b1.Balance(), " , Account 2 balance: ", a1.Balance())
	err = Transfer(a1, b1, 90)
	if err != nil {
		log.Println("ERROR: ", err)
	}

	fmt.Println("Account 1 Balance: ", b1.Balance(), " , Account 2 balance: ", a1.Balance())

}

// Value: get a copy of the struct, if I change nothing will happens to the real struct
// pointer:  changing the original instance
