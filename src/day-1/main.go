package main

import (
	filehelpers "aoc-2024/src/helpers"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	res := part1()
	fmt.Println(res)

	part2()
}

func part1() int {
	lines := filehelpers.GetLines("input.txt")

	list1 := make([]int, len(lines))
	list2 := make([]int, len(lines))

	for i, line := range lines {
		stringNums := strings.Split(line, "   ")
		num1, err := strconv.Atoi(stringNums[0])
		if err != nil {
			fmt.Println(i, stringNums)
			panic(err)
		}
		num2, err := strconv.Atoi(stringNums[1])
		if err != nil {
			fmt.Println(i, stringNums)
			panic(err)
		}

		list1[i] = num1
		list2[i] = num2
	}

	slices.Sort(list1)
	slices.Sort(list2)

	total := 0
	for i, num1 := range list1 {
		num2 := list2[i]

		diff := math.Abs(float64(num1 - num2))
		total += int(diff)
	}

	return total
}

func part2() {

}
