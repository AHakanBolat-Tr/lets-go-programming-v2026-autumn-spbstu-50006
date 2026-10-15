package temperature_test

import (
	"errors"
	"testing"

	"github.com/AHakanBolat-Tr/task-2-1/internal/temperature"
)

type constraint struct {
	operator string
	value    int
}

type rangeTestCase struct {
	name        string
	constraints []constraint
	expected    []int
}

func rangeTestCases() []rangeTestCase {
	return []rangeTestCase{
		{
			name:        "assignment example second department",
			constraints: []constraint{{">=", 18}, {"<=", 23}, {">=", 20}, {"<=", 27}, {"<=", 21}, {">=", 28}},
			expected:    []int{18, 18, 20, 20, 20, -1},
		},
		{
			name:        "lower boundary",
			constraints: []constraint{{"<=", 15}, {">=", 15}},
			expected:    []int{15, 15},
		},
		{
			name:        "upper boundary",
			constraints: []constraint{{">=", 30}, {"<=", 30}},
			expected:    []int{30, 30},
		},
		{
			name:        "contradiction",
			constraints: []constraint{{">=", 25}, {"<=", 20}},
			expected:    []int{25, -1},
		},
		{
			name:        "contradiction stays after more constraints",
			constraints: []constraint{{"<=", 18}, {">=", 22}, {"<=", 29}, {">=", 16}},
			expected:    []int{15, -1, -1, -1},
		},
		{
			name:        "repeated compatible constraints",
			constraints: []constraint{{">=", 17}, {">=", 16}, {"<=", 28}, {">=", 21}, {"<=", 25}, {">=", 21}},
			expected:    []int{17, 17, 17, 21, 21, 21},
		},
		{
			name:        "lower bound above physical range",
			constraints: []constraint{{">=", 100}},
			expected:    []int{-1},
		},
		{
			name:        "upper bound below physical range",
			constraints: []constraint{{"<=", -10}},
			expected:    []int{-1},
		},
		{
			name:        "bounds outside physical range do not narrow it",
			constraints: []constraint{{">=", -10}, {"<=", 100}},
			expected:    []int{15, 15},
		},
	}
}

func TestRangeOptimalAfterEachConstraint(t *testing.T) {
	t.Parallel()

	for _, testCase := range rangeTestCases() {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			allowed := temperature.NewRange()

			for step, current := range testCase.constraints {
				err := allowed.Apply(current.operator, current.value)
				if err != nil {
					t.Fatalf("step %d: unexpected error: %v", step, err)
				}

				got := allowed.Optimal()
				if got != testCase.expected[step] {
					t.Fatalf("step %d (%s %d): got %d, want %d",
						step, current.operator, current.value, got, testCase.expected[step])
				}
			}
		})
	}
}

func TestRangeUnknownOperator(t *testing.T) {
	t.Parallel()

	allowed := temperature.NewRange()

	err := allowed.Apply("==", 20)
	if !errors.Is(err, temperature.ErrUnknownOperator) {
		t.Fatalf("got error %v, want %v", err, temperature.ErrUnknownOperator)
	}
}
