package main

import "fmt"

func printMenu() {
	fmt.Println("\033[36m==============================\033[0m")
	fmt.Println("\033[36m       GO CALCULATOR CLI       \033[0m")
	fmt.Println("\033[36m==============================\033[0m")
	fmt.Println("Available operations: ")
	fmt.Println("+  -  *  /  ^  sqrt  %  square  abs")
}
