
package account

import ("fmt"
	"customer")

type CheckingAccount struct {
	Holder  customer.Holder
	Branch  int
	Sccount int
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
		return fmt.Sprintf("%.2f successfully transferred to %v.", cash, destinationAccount.Holder)
	} else {
		return fmt.Sprintf("Transfer failed.")
	}
}

func (c CheckingAccount) String() string {
	return fmt.Sprintf("%s is the holder. %d-%d is the branch and account. %.2f is the balance.",
		c.Holder, c.Branch, c.Account, c.balance)
}

func (c *CheckingAccount) Statement() string {
	return fmt.Sprintf("Your current balance is %.2f", c.balance)
}
