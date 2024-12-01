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

	res := part1(list1, list2)
	fmt.Println(res)

	res = part2(list1, list2)
	fmt.Println(res)
}

func part1(list1 []int, list2 []int) int {
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

func part2(list1 []int, list2 []int) int {
	countMap := make(map[int]int)
	for _, num2 := range list2 {
		countMap[num2]++
	}

	total := 0
	for _, num1 := range list1 {
		val, ok := countMap[num1]
		if ok {
			total += (num1 * val)
		}
	}

	return total
}
