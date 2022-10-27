package internals

import (
	"fmt"
)

var balance float64 // The amount of money a user has

// AccountBalance() prints out the balance of the user
func AccountBalance() {
	fmt.Printf("Your account Balance is %.2f Ksh", balance)
	newLine(1)
}

// WithdrawFunds() handles the logic of withdrawing funds from the user account
func WithdrawFunds() {
	newLine(1)
	fmt.Println("Withdraw funds")
	var amount float64

	if balance == 0 {
		fmt.Println("Withdrawal not Completed due to Insufficient Balance")
		return
	}
	fmt.Printf("Enter amount: ")
	fmt.Scan(&amount)
	if amount < 0 {
		fmt.Println("invalid amount, enter amount greater than 0")
		return
	}

	if amount > balance {
		fmt.Println("insufficient balance")
		fmt.Printf("Current balance: %.2f Ksh", balance)
		return
	}
	prev_balance := balance
	balance -= amount

	fmt.Printf("\nPrevious balance: %.2f \n", prev_balance)
	fmt.Printf("Amount withdrawn: %.2f  Ksh,\nCurrent balance: %.2f  Ksh", amount, balance)
}

// DepositFunds() function handles the Deposit operations
func DepositFunds() {
	newLine(1)
	fmt.Println("deposit funds")
	var amount float64
	fmt.Printf("Enter amount: ")
	fmt.Scan(&amount)
	if amount < 0 {
		fmt.Println("invalid amount")
		fmt.Println("You cannot deposit less than 0 Ksh")
		return
	}
	prev_balance := balance
	balance += amount

	newLine(1)
	fmt.Printf("Previous Balance: %.2f  Ksh,\nCurrent balance: %.2f Ksh", prev_balance, balance)
}
