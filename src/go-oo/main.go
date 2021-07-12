package main

import "fmt"

type CheckingAccount struct {
	holder  string
	branch  int
	account int
	balance float64
}

func (c *CheckingAccount) Withdraw(cash float64) string {
	if cash <= c.balance && cash > 0 {
		c.balance -= cash
		return fmt.Sprintf("%.2f successfully withdrawn. Yor new balance is: %.2f.", cash, c.balance)
	} else {
		return fmt.Sprintf("Withdraw failed. Your balance is only %.2f.", c.balance)
	}
}

func (c *CheckingAccount) Deposit(cash float64) string {
	if cash > 0 {
		c.balance += cash
		return fmt.Sprintf("%.2f deposited successfully. Your new balance is: %.2f.",
			cash, c.balance)
	} else {
		return fmt.Sprintf("Deposit failed.")
	}
}

func (c *CheckingAccount) Transfer(cash float64, destinationAccount *CheckingAccount) string {
	if cash <= c.balance && cash > 0 {
		c.balance -= cash
		destinationAccount.balance += cash
		return fmt.Sprintf("%.2f successfully transferred to %v.", cash, destinationAccount.holder)
	} else {
		return fmt.Sprintf("Transfer failed.")
	}
}

func (c CheckingAccount) String() string {
	return fmt.Sprintf("%s is the holder. %d-%d is the branch and account. %.2f is the balance.",
		c.holder, c.branch, c.account, c.balance)
}

func main() {
	// var erickAcc CheckingAccount = CheckingAccount{
	// 	holder:  "Erick",
	// 	branch:  123,
	// 	account: 123123,
	// 	balance: 500.99}

	amorimAcc := CheckingAccount{
		"Amorim",
		123,
		123123,
		550.99,
	}

	erickAcc := CheckingAccount{
		"Erick",
		123,
		123123,
		550.99,
	}

	// var erickAcc *CheckingAccount
	// erickAcc = new(CheckingAccount)

	// fmt.Println(erickAcc)

	fmt.Println(erickAcc.balance)
	fmt.Println(erickAcc.Withdraw(600))
	fmt.Println(erickAcc.balance)

	fmt.Println(erickAcc.Deposit(500))
	fmt.Println(erickAcc.balance)

	fmt.Println(amorimAcc.balance)
	fmt.Println(erickAcc.Transfer(100, &amorimAcc))
	fmt.Println(amorimAcc.balance)
}
