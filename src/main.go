package main

import (
	"fmt"
	"account"
	"customer")


func main() {
	// var erickAcc account.CheckingAccount = CheckingAccount{
	// 	Holder:  "Erick",
	// 	Branch:  123,
	// 	Account: 123123,
	// 	Balance: 500.99}

	amorimAcc := account.CheckingAccount{
		"Amorim",
		123,
		123123,
		550.99,
	}

	erickAcc := account.CheckingAccount{
		"Erick",
		123,
		123123,
		550.99,
	}

	// var erickAcc *CheckingAccount
	// erickAcc = new(CheckingAccount)

	// fmt.Println(erickAcc)

	limaCust := customer.Holder{
			Name: "Lima",
			CPF: "12312312300",
			Occupation: "Dev Go"}

	limaAcc := account.CheckingAccount{
		Holder: limaCust,
		Branch: 789,
		Account: 123456,
		Balance: 0.0,
	}



}

