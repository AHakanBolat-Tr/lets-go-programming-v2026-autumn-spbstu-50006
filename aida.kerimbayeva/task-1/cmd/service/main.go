package main

import (
	"fmt"
)

func main() {
	var x, y int
	var op string

	_, err := fmt.Scanln(&x)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanln(&y)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scanln(&op)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch op {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		if y == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(x / y)
	default:
		fmt.Println("Invalid operation")
	}
}
