package util

import (
	"math"

	"github.com/obscurelyme/aoc2024/day1/types"
)

func Sum(input *types.ListInput) float64 {
	sum := 0.0

	for i := 0; i < len(input.List1); i++ {
		sum += math.Abs(float64(input.List1[i] - input.List2[i]))
	}

	return sum
}
