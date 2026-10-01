package main

import (
	"errors"
	"fmt"
)

func readOperand() (string, error) {
	var operand string

	_, err := fmt.Scan(&operand)
	if err != nil {
		return "", errors.New("Invalid operation")
	}
	return operand, nil
}

func add(a, b int) (int, error) {
	return a + b, nil
}

func subtract(a, b int) (int, error) {
	return a - b, nil
}

func multiply(a, b int) (int, error) {
	return a * b, nil
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}

func calculate(a, b int, operand string) (int, error) {
	switch operand {
	case "+":
		return add(a, b)
	case "-":
		return subtract(a, b)
	case "*":
		return multiply(a, b)
	case "/":
		return divide(a, b)
	default:
		return 0, errors.New("Invalid operation")
	}
}

func main() {
	var a, b int
	var operation string
	var err error

	_, err = fmt.Scan(&a)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	operation, err = readOperand()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result int
	result, err = calculate(a, b, operation)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
}
