package utils

import (
	"math"
)

const epsilonConstraint = 1e-9

func Epsilon (a, b float64) bool {
	return math.Abs(a -b) < epsilonConstraint;
}