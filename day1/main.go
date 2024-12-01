package main

import (
	"fmt"

	"github.com/obscurelyme/aoc2024/day1/types"
	"github.com/obscurelyme/aoc2024/day1/util"
)

var input types.ListInput = types.ListInput{}

func main() {
	util.ReadInput(&input)
	util.Sort(&input)
	sum := util.Sum(&input)
	similarity := util.Similarity(&input)

	fmt.Printf("The sum is: %.0f\n", sum)
	fmt.Printf("The similarity is: %d\n", similarity)
}
