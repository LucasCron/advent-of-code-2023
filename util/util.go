package util

import (
	"bufio"
	"os"
)

var (
	digits = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
)

func IsNumber(c string) bool {
	for _, d := range digits {
		if c == d {
			return true
		}
	}
	return false
}

func ReadTextFileToArray(path string) []string {
	file, err := os.Open(path)
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
	}

	if err = scanner.Err(); err != nil {
		Check(err)
	}
	return lines
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
