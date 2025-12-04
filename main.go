package main

import "fmt"

func main() {

	for {

		var a, b float64
		var op string

		fmt.Print("Enter first num: ")
		fmt.Scan(&a)
		fmt.Print("Enter second num: ")
		fmt.Scan(&b)
		fmt.Print("Input operation (+, -, *, /): ")
		fmt.Scan(&op)

		switch op {
		case "+":
			fmt.Println("Res:", a+b)
		case "-":
			fmt.Println("Res: ", a-b)
		case "*":
			fmt.Println("Res: ", a*b)
		case "/":
			if b == 0 {
				fmt.Println("Error! devision by zero!")
			} else {
				fmt.Println("Res: ", a/b)
			}
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
