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
}

func part1(lines []string) int {
	total := 0
	for _, line := range lines {
		r := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
		matches := r.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			num1, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}

			num2, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}

			total += num1 * num2
		}
	}
	return total
}

func part2() {

}
