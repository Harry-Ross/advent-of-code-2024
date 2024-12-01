package helpers

import (
	"os"
	"strings"
)

func GetLines(filename string) []string {
	data, error := os.ReadFile(filename)

	if error != nil {
		panic(error)
	}

	var strData string = string(data)
	lines := strings.Split(strData, "\n")

	return lines
}
