package main

import (
	"fmt"
	"strconv"
	"strings"

	"lucascron.com/advent-of-code-2023/util"
)

var (
	files = [2]string{
		"example.txt",
		"input.txt",
	}
	digits = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
)

type Number struct {
	value           int
	row             int
	startIndex      int
	endIndex        int
	shouldBeCounted bool
}

func (n Number) print() {
	fmt.Println(fmt.Sprintf("Value: %d, Row: %d, Start Index: %d, End Index: %d, ShouldBeCounted: %t", n.value, n.row, n.startIndex, n.endIndex, n.shouldBeCounted))
}

type Symbol struct {
	row   int
	index int
	value string
}

func (s Symbol) print() {
	fmt.Println(fmt.Sprintf("Row: %d, Index: %d, Value: %s", s.row, s.index, s.value))
}

func main() {
	for _, path := range files {
		lines := util.ReadTextFileToArray(path)

		var numbers []Number
		var symbols []Symbol

		for row, line := range lines {
			// fmt.Println(line)
			chars := strings.Split(line, "")

			trackingNumber := false
			var number Number
			var numberArray []string
			for i, c := range chars {
				if !trackingNumber && isNumber(c) {
					number = Number{
						startIndex: i,
						row:        row,
					}
					numberArray = append(numberArray, c)
					trackingNumber = true
				} else if trackingNumber && isNumber(c) {
					numberArray = append(numberArray, c)
				}
				if trackingNumber && (!isNumber(c) || i+1 == len(chars)) {
					endOffset := 0
					if isNumber(c) && i+1 == len(chars) {
						endOffset = 1
					}
					number.endIndex = i - 1 + endOffset
					number.value = concatNumberCharacterArrayToString(numberArray)
					numbers = append(numbers, number)

					trackingNumber = false
					numberArray = nil
				}

				if !isPeriod(c) && !isNumber(c) {
					symbols = append(symbols, Symbol{
						row:   row,
						index: i,
						value: c,
					},
					)
				}
			}

		}

		var partSum int

		for _, n := range numbers {
			for _, s := range symbols {
				if (n.row == s.row && (n.startIndex-1 == s.index || n.endIndex+1 == s.index)) ||
					((s.row == n.row-1 || s.row == n.row+1) && s.index <= n.endIndex+1 && s.index >= n.startIndex-1) {
					n.shouldBeCounted = true
					partSum += n.value
				}
			}
		}

		var gearRatioSum int
		for _, s := range symbols {
			if s.value == "*" {
				var gearsAdjacent int
				gearRatio := 1
				for _, n := range numbers {
					if (n.row == s.row && (s.index == n.endIndex+1 || s.index == n.startIndex-1)) ||
						((s.row == n.row+1 || s.row == n.row-1) && (s.index >= n.startIndex-1 && s.index <= n.endIndex+1)) {
						gearRatio *= n.value
						gearsAdjacent += 1
					}

				}
				if gearsAdjacent == 2 {
					gearRatioSum += gearRatio
				}
			}
		}

		fmt.Println(partSum)
		fmt.Println(gearRatioSum)
	}
}

func isNumber(c string) bool {
	for _, d := range digits {
		if c == d {
			return true
		}
	}
	return false
}

func isPeriod(c string) bool {
	return c == "."
}

func concatNumberCharacterArrayToString(s []string) int {
	var numString string
	for _, v := range s {
		numString = numString + v
	}
	num, err := strconv.Atoi(numString)
	util.Check(err)
	return num
}
