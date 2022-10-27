package main

import (
	"fmt"
	"os"

	g "github.com/onukajiuche/altschool-go-task-3/internals"
)

// welcome prints welcome message for the user
func welcome() {
	fmt.Println("*******{ Welcome to the ATM CLI app }*******")
}

// exitProgram() is used to exit the cli application
func exitProgram() {
	fmt.Println("Thank You For Banking With Us.")
	os.Exit(0)
}

/*
newLine takes in an integer as an argument and returns the integer
number of new lines
*/
func newLine(n int) {
	for n > 0 {
		fmt.Println("")
		n--
	}
}

/*
anotherOperation() is used to ask the user if they want to perform
another  operation or exit the program
*/
func anotherOperation() {
	var yesOrNo string
	newLine(2)
	fmt.Println("Do you want to perform another operation")
	fmt.Println("if Yes type y, else type anything")
	fmt.Scan(&yesOrNo)
	if yesOrNo == "y" {
		menu()
	}
	exitProgram()
}

/*
menu function handles the display and selection of operation in
the cli application
*/
func menu() {
	newLine(1)
	fmt.Println("Select Operation:")
	fmt.Println("1. Change pin \n2. Account Balance")
	fmt.Println("3. Withdraw funds \n4. Deposit funds")
	fmt.Println("0. Exit program \t")
	newLine(1)

	var menuNumber int
	fmt.Printf("Enter operation number: ")
	_, err := fmt.Scan(&menuNumber)
	if err != nil {
		fmt.Println("Error:, please select the correct menu item")
		//menu()
	}

	switch menuNumber {
	case 1:
		g.ChangePin()
		anotherOperation()
	case 2:
		g.AccountBalance()
		anotherOperation()
	case 3:
		g.WithdrawFunds()
		anotherOperation()
	case 4:
		g.DepositFunds()
		anotherOperation()
	case 0:
		exitProgram()
	default:
		fmt.Println("Invalid input ")
		anotherOperation()
	}
}

func main() {
	g.Login()
	welcome()
	menu()
}
