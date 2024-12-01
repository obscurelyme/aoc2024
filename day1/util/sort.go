package util

import (
	"sort"

	"github.com/obscurelyme/aoc2024/day1/types"
)

func Sort(input *types.ListInput) {
	sort.Slice(input.List1, func(a int, b int) bool {
		return input.List1[a] > input.List1[b]
	})

	sort.Slice(input.List2, func(a int, b int) bool {
		return input.List2[a] > input.List2[b]
	})
}
