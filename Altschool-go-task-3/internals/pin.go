package internals

import "fmt"

var pin string = "4194"

/* newLine takes in an integer as an argument and returns the integer
number of new lines */
func newLine(n int) {
	for n > 0 {
		fmt.Println("")
		n--
	}
}

// Login() handles the user login
func Login() {
	for {

		if !VerifyPin(pin) {
			continue
		}

		break
	}
}

/* VerifyPin - handles pin verification
 arg := pin string
return => bool (true if the input matches the pic and false if otherwise)
*/
func VerifyPin(pin string) bool {
	var input_pin string
	fmt.Printf("ENTER PIN: ")
	fmt.Scan(&input_pin)

	if pin != input_pin {
		fmt.Println("Incorrect pin")
		return false
	}

	return true
}

// changePin() is used to change the user pin
func ChangePin() {
	newLine(1)
	fmt.Println("change pin")
	for {
		fmt.Printf("Enter new pin: ")
		var new_pin string
		fmt.Scan(&new_pin)

		if len(new_pin) != 4 {
			fmt.Println("Pin should be 4 characters long")
			continue
		}

		pin = new_pin
		fmt.Println("Pin has been changed")
		break
	}
}
