package PointersAndErrors

import "fmt"

type Wallet struct {
	//balance int
	balance Bitcoin
}

type Bitcoin int

type Stringer interface {
	String() string
}

func (bitcoin Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", bitcoin)
}

func (wallet *Wallet) Deposit(amount /*int*/ Bitcoin) {
	fmt.Printf("address of balance in test is %p \n", &wallet.balance)
	wallet.balance += amount
}

func (wallet *Wallet) Balance() /*int*/ Bitcoin {
	return wallet.balance
}

func (wallet *Wallet) Withdraw(amount Bitcoin) {
	wallet.balance -= amount
}
