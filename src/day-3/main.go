package main

import (
	filehelpers "aoc-2024/src/helpers"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines := filehelpers.GetLines("input.txt")
	res := part1(lines)
	fmt.Println(res)

	res = part2(lines)
	fmt.Println(res)
}

func part1(lines []string) int {
	total := 0
	for _, line := range lines {
		r := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
		matches := r.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			total += multiplyStrings(match[1], match[2])
		}
	}
	return total
}

func part2(lines []string) int {
	total := 0

	for _, line := range lines {
		r := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)|don't\(\)|do\(\)`)
		matches := r.FindAllStringSubmatch(line, -1)

		execute := true
		for _, match := range matches {
			if match[0] == "don't()" {
				execute = false
			} else if match[0] == "do()" {
				execute = true
			} else {
				if execute {
					total += multiplyStrings(match[1], match[2])
				}
			}
		}
	}

	return total
}

func multiplyStrings(str1 string, str2 string) int {
	num1, err := strconv.Atoi(str1)
	if err != nil {
		panic(err)
	}

	num2, err := strconv.Atoi(str2)
	if err != nil {
		panic(err)
	}

	return num1 * num2
}
