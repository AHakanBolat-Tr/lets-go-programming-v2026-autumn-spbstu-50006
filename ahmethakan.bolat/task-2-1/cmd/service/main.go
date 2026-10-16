package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	minTemp = 15
	maxTemp = 30
)

func processDepartment(reader *bufio.Reader) bool {
	var employees int

	_, err := fmt.Fscan(reader, &employees)
	if err != nil {
		fmt.Println("Invalid number of employees")

		return false
	}

	low, high := minTemp, maxTemp

	for range employees {
		var (
			sign string
			temp int
		)

		_, err = fmt.Fscan(reader, &sign, &temp)
		if err != nil {
			fmt.Println("Invalid temperature")

			return false
		}

		switch sign {
		case ">=":
			low = max(low, temp)
		case "<=":
			high = min(high, temp)
		default:
			fmt.Println("Invalid operation")

			return false
		}

		if low > high {
			fmt.Println(-1)
		} else {
			fmt.Println(low)
		}
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var departments int

	_, err := fmt.Fscan(reader, &departments)
	if err != nil {
		fmt.Println("Invalid number of departments")

		return
	}

	for range departments {
		if !processDepartment(reader) {
			return
		}
	}
}
