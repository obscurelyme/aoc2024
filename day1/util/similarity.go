package util

import (
	"github.com/obscurelyme/aoc2024/day1/types"
)

func incrementInMap(key int, m map[int]int) {
	_, ok := m[key]
	if ok {
		// value does exist
		m[key]++
	} else {
		// value does not exist
		m[key] = 1
	}
}

// Runtime O(n)
func Similarity(input *types.ListInput) int {
	occurances := map[int]int{}
	similarity := 0

	for i := 0; i < len(input.List2); i++ {
		occurances[input.List2[i]]++
	}

	for i := 0; i < len(input.List1); i++ {
		current := input.List1[i]
		similarity += occurances[current] * current
	}

	return similarity
}

// My original answer, non-optimal runtime O(n^2)
// func Similarity(input *types.ListInput) int {
// 	occurances := map[int]int{}
// 	similarity := 0

// 	for i := 0; i < len(input.List1); i++ {
// 		current := input.List1[i]
// 		for j := 0; j < len(input.List2); j++ {
// 			if current == input.List2[j] {
// 				incrementInMap(current, occurances)
// 			}
// 		}
// 	}

// 	for key, value := range occurances {
// 		similarity += key * value
// 	}

// 	return similarity
// }
