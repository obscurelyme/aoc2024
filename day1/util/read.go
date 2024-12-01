package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/obscurelyme/aoc2024/day1/types"
)

func ReadInput(input *types.ListInput) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		current := scanner.Text()
		values := strings.Split(current, "   ")
		firstValue, _ := strconv.Atoi(values[0])
		secondValue, _ := strconv.Atoi(values[1])

		input.List1 = append(input.List1, firstValue)
		input.List2 = append(input.List2, secondValue)
	}
}
