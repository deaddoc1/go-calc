package main

import (
	"fmt"
)

func main() {

	for {

		printMenu()

		var a, b float64
		var op string

		fmt.Print("Input operation (+, -, *, /, ^, sqrt, %, square, abs): ")
		fmt.Scan(&op)

		fmt.Print("Enter first num: ")
		fmt.Scan(&a)

		if op != "sqrt" && op != "square" && op != "abs" {
			fmt.Print("Enter second num: ")
			fmt.Scan(&b)
		}

		switch op {
		case "+":
			fmt.Println("Res:", add(a, b))

		case "-":
			fmt.Println("Res: ", sub(a, b))

		case "*":
			fmt.Println("Res: ", mul(a, b))

		case "/":
			result, err := div(a, b)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Res: ", result)
			}

		case "^":
			fmt.Println("Res: ", pow(a, b))

		case "sqrt":
			fmt.Println("Res: ", sqrt(a))

		case "%":
			result, err := mod(a, b)
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("Res: ", result)
			}

		case "square":
			fmt.Println("Res: ", square(a))

		case "abs":
			fmt.Println("Res: ", abs(a))

		default:
			fmt.Println("Unknown opreation. Try again.")
		}

		var cont string
		fmt.Print("Continue? (y/n): ")
		fmt.Scan(&cont)

		if cont != "y" && cont != "Y" {
			fmt.Println("Goodbye!")
			break
		}
	}
}
