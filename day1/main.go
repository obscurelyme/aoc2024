package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var list1 []int
var list2 []int

func main() {
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

		list1 = append(list1, firstValue)
		list2 = append(list2, secondValue)
	}

	sort.Slice(list1, func(a int, b int) bool {
		return list1[a] > list1[b]
	})

	sort.Slice(list2, func(a int, b int) bool {
		return list2[a] > list2[b]
	})

	sum := 0.0
	for i := 0; i < len(list1); i++ {
		sum += math.Abs(float64(list1[i] - list2[i]))
	}

	fmt.Printf("The sum is: %.0f\n", sum)
}
