package main

import (
	"fmt"
	"strconv"
	"strings"

	"lucascron.com/advent-of-code-2023/util"
)

type Number struct {
	val  string
	word string
}

var (
	files = [2]string{
		"example.txt",
		"input.txt",
	}
	numbers = [9]Number{
		{
			val:  "1",
			word: "one",
		},
		{
			val:  "2",
			word: "two",
		},
		{
			val:  "3",
			word: "three",
		},
		{
			val:  "4",
			word: "four",
		},
		{
			val:  "5",
			word: "five",
		},
		{
			val:  "6",
			word: "six",
		},
		{
			val:  "7",
			word: "seven",
		},
		{
			val:  "8",
			word: "eight",
		},
		{
			val:  "9",
			word: "nine",
		},
	}
)

func main() {
	for _, path := range files {
		lines := util.ReadTextFileToArray(path)
		var sum int
		for _, line := range lines {
			front := getFirstNumber(line, false)
			back := getFirstNumber(line, true)

			calibrationValueStr := fmt.Sprintf("%s%s", front, back)
			calibrationValue, err := strconv.Atoi(calibrationValueStr)
			util.Check(err)

			fmt.Println(fmt.Sprintf("Line: %s, Front: %s, Back: %s", line, front, back))
			sum = sum + calibrationValue
		}
		fmt.Println(fmt.Sprintf("File path: %s, Sum: %d", path, sum))

	}
}

func getFirstNumber(text string, reversed bool) string {
	chars := strings.Split(text, "")

	if reversed {
		chars = util.Reverse(chars)
	}

	for i, c := range chars {
		for _, n := range numbers {
			word := n.word
			if reversed {
				word = util.ReverseString(n.word)
			}

			if c == string(n.val) {
				return n.val
			}

			firstChar := word[0:1]
			if c == firstChar {
				wordLen := len(word)
				if i+wordLen < len(chars) {
					next := chars[i : i+wordLen]
					nextWord := strings.Join(next, "")
					if nextWord == word {
						return n.val
					}
				}
			}
		}
	}

	return "0"
}
