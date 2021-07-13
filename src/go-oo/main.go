package main

import (
	"fmt"
	"go-oo/savingsaccount"
)

func main() {
	amorimSa := savingsaccount.SavingsAccount{}
	amorimSa.Deposit(100)
	fmt.Println(amorimSa.Statement())
}

func PayBillet(account verifyAccount, cash float64) {
	account.Withdraw(cash)
}

type verifyAccount interface {
	Withdraw(cash float64) string
}
