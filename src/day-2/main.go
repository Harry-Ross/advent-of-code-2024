package main

import (
	filehelpers "aoc-2024/src/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := filehelpers.GetLines("input.txt")
	res := part1(lines)
	fmt.Println(res)
}

func part1(lines []string) int {
	unsafeCount := 0

	for _, line := range lines {
		splitNums := strings.Split(line, " ")

		firstNum, err := strconv.Atoi(splitNums[0])
		if err != nil {
			panic(err)
		}
		lastNum := firstNum
		lastDiff := 0
		unsafe := false
		for _, num := range splitNums[1:] {
			currNum, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			diff := currNum - lastNum

			if diff > 3 || diff == 0 || diff < -3 {
				unsafe = true
				break
			}

			if (diff > 0 && lastDiff < 0) || (diff < 0 && lastDiff > 0) {
				unsafe = true
				break
			}

			lastNum = currNum
			lastDiff = diff
		}

		if unsafe {
			unsafeCount++
		}
	}

	return len(lines) - unsafeCount
}

func part2(lines []string) int {
	return 0
}
