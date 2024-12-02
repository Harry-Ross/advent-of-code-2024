package main

import (
	filehelpers "aoc-2024/src/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := filehelpers.GetLines("example.txt")
	res := part1(lines)
	fmt.Println(res)

	res = part2(lines)
	fmt.Println(res)
}

func part1(lines []string) int {
	unsafeCount := 0

	for _, line := range lines {
		splitNums := strings.Split(line, " ")

		report := make([]int, 0)
		for _, num := range splitNums {
			currNum, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			report = append(report, currNum)
		}
		if !isReportSafe(report, false) {
			unsafeCount++
		}
	}

	return len(lines) - unsafeCount
}

func part2(lines []string) int {
	unsafeCount := 0

	for _, line := range lines {
		splitNums := strings.Split(line, " ")

		report := make([]int, 0)
		for _, num := range splitNums {
			currNum, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			report = append(report, currNum)
		}

		if !isReportSafe(report, true) {
			unsafeCount++
			fmt.Println("unsafe", report)
		} else {
			fmt.Println("safe", report)
		}
	}

	return len(lines) - unsafeCount
}

func isReportSafe(report []int, allowRemoval bool) bool {
	if len(report) < 2 {
		return true
	}

	direction := 0
	for i, num := range report {
		if i == 0 {
			continue
		}

		diff := num - report[i-1]

		if direction == 0 {
			if diff > 0 {
				direction = 1
			} else if diff < 0 {
				direction = -1
			}
		}

		if isUnsafe(direction, diff) {
			if allowRemoval {
				if isReportSafe(remove(report, i-1), false) {
					fmt.Println("removed", report[i-1])
					return true
				}
			}
			return false
		}

	}

	return true
}

func isUnsafe(direction int, diff int) bool {
	return (diff > 3 || diff == 0 || diff < -3) || (direction != 0 && (diff > 0 && direction < 0) || (diff < 0 && direction > 0))
}

func remove(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
