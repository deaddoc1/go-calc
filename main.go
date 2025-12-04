package main

import (
	"fmt"
	"math"
)

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

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}

func sqrt(a float64) float64 {
	return math.Sqrt(a)
}

func mod(a, b float64) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return int(a) % int(b), nil
}

func square(a float64) float64 {
	return a * a
}

func abs(a float64) float64 {
	return math.Abs(a)
}

func main() {

	for {

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
