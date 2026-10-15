package temperature

import (
	"errors"
	"fmt"
)

const (
	MinAllowed = 15
	MaxAllowed = 30
	NoSolution = -1
)

var ErrUnknownOperator = errors.New("unknown operator")

// Range is the interval of temperatures that satisfies every constraint applied so far.
type Range struct {
	lower int
	upper int
}

func NewRange() *Range {
	return &Range{
		lower: MinAllowed,
		upper: MaxAllowed,
	}
}

// Apply narrows the range with a constraint "<= value" or ">= value".
func (r *Range) Apply(operator string, value int) error {
	switch operator {
	case ">=":
		r.lower = max(r.lower, value)
	case "<=":
		r.upper = min(r.upper, value)
	default:
		return fmt.Errorf("%w: %q", ErrUnknownOperator, operator)
	}

	return nil
}

// Optimal returns the lowest suitable temperature or NoSolution if the range is empty.
func (r *Range) Optimal() int {
	if r.lower > r.upper {
		return NoSolution
	}

	return r.lower
}
