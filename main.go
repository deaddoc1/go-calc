package main

import "fmt"

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

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
