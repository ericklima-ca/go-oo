package savingsaccount

import ("fmt"
	"customer")

type SavingsAccount struct {
	Holder customer.Holder
	Branch int
	Account int
	Operation int
	balance float64
}

func (c *SavingsAccount) Withdraw(cash float64) string {
	if cash <= c.balance && cash > 0 {
		c.balance -= cash
		return fmt.Sprintf("%.2f successfully withdrawn. Yor new balance is: %.2f.", cash, c.balance)
	} else {
		return fmt.Sprintf("Withdraw failed. Your balance is only %.2f.", c.balance)
	}
}

func (c *SavingsAccount) Deposit(cash float64) string {
	if cash > 0 {
		c.balance += cash
		return fmt.Sprintf("%.2f deposited successfully. Your new balance is: %.2f.",
			cash, c.balance)
	} else {
		return fmt.Sprintf("Deposit failed.")
	}
}


func (c *SavingsAccount) Statement() string {
	return fmt.Sprintf("Your current balance is %.2f", c.balance)
}
