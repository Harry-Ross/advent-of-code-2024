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

	res = part2(lines)
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

			if isUnsafe(lastDiff, diff) {
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
		for i, num := range splitNums[1:] {

			currNum, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			diff := currNum - lastNum

			if isUnsafe(lastDiff, diff) {
				if i != len(splitNums[1:])-1 {
					nextNum, err := strconv.Atoi(splitNums[i+2])
					if err != nil {
						panic(err)
					}

					diff = nextNum - lastNum
					if isUnsafe(lastDiff, diff) {
						unsafe = true
						break
					}
				} else {
					unsafe = true
					break
				}
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

func isUnsafe(lastDiff int, diff int) bool {
	return (diff > 3 || diff == 0 || diff < -3) || (diff > 0 && lastDiff < 0) || (diff < 0 && lastDiff > 0)
}
