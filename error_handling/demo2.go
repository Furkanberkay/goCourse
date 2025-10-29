package errorhandling

import "fmt"

type BankAccount interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	CurrentBalance() float64
}

type CheckingAccount struct {
	balance float64
}

func (c *CheckingAccount) Deposit(amount float64) {
	c.balance += amount
}

func (c *CheckingAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be > 0")
	}
	if amount > c.balance {
		return fmt.Errorf("insufficient funds")
	}
	c.balance -= amount
	return nil
}

func (c *CheckingAccount) CurrentBalance() float64 {
	return c.balance
}

func Demo3() {
	var acc BankAccount = &CheckingAccount{balance: 200} // ← interface'e bağla
	if err := acc.Withdraw(50); err != nil {
		fmt.Println("ÇEKİM HATASI:", err)
		return
	}
	fmt.Println("Kalan bakiye:", acc.CurrentBalance())
}
