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

	if isIncreasing(report) || isDecreasing(report) {
		return true
	}

	if allowRemoval {
		for i := 0; i < len(report); i++ {
			removed := remove(report, i)
			if isIncreasing(removed) || isDecreasing(removed) {
				fmt.Println("removed", report[i])
				return true
			}
		}
	}

	return false
}

func isIncreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isDecreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff > -1 || diff < -3 {
			return false
		}
	}
	return true
}

func remove(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
