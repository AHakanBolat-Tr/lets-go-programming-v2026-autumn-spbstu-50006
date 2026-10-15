package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AHakanBolat-Tr/task-2-1/internal/temperature"
)

func main() {
	err := run(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(reader *bufio.Reader, writer *bufio.Writer) error {
	var departments int

	_, err := fmt.Fscan(reader, &departments)
	if err != nil {
		return fmt.Errorf("read number of departments: %w", err)
	}

	for range departments {
		err = processDepartment(reader, writer)
		if err != nil {
			return err
		}
	}

	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("flush output: %w", err)
	}

	return nil
}

func processDepartment(reader *bufio.Reader, writer *bufio.Writer) error {
	var employees int

	_, err := fmt.Fscan(reader, &employees)
	if err != nil {
		return fmt.Errorf("read number of employees: %w", err)
	}

	allowed := temperature.NewRange()

	for range employees {
		var (
			operator string
			value    int
		)

		_, err = fmt.Fscan(reader, &operator, &value)
		if err != nil {
			return fmt.Errorf("read constraint: %w", err)
		}

		err = allowed.Apply(operator, value)
		if err != nil {
			return fmt.Errorf("apply constraint: %w", err)
		}

		_, err = fmt.Fprintln(writer, allowed.Optimal())
		if err != nil {
			return fmt.Errorf("write answer: %w", err)
		}
	}

	return nil
}
