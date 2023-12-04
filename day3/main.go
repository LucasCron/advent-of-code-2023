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
)

type Number struct {
	value           int
	row             int
	startIndex      int
	endIndex        int
	shouldBeCounted bool
}

type Symbol struct {
	row   int
	index int
	value string
}

func main() {
	for _, path := range files {
		lines := util.ReadTextFileToArray(path)

		var numbers []Number
		var symbols []Symbol

		for row, line := range lines {
			chars := strings.Split(line, "")

			trackingNumber := false
			var number Number
			var numberArray []string
			for i, c := range chars {
				if !trackingNumber && util.IsNumber(c) {
					number = Number{
						startIndex: i,
						row:        row,
					}
					numberArray = append(numberArray, c)
					trackingNumber = true
				} else if trackingNumber && util.IsNumber(c) {
					numberArray = append(numberArray, c)
				}
				if trackingNumber && (!util.IsNumber(c) || i+1 == len(chars)) {
					endOffset := 0
					if util.IsNumber(c) && i+1 == len(chars) {
						endOffset = 1
					}
					number.endIndex = i - 1 + endOffset
					number.value = numCharArrayToInt(numberArray)
					numbers = append(numbers, number)

					trackingNumber = false
					numberArray = nil
				}

				if c != "." && !util.IsNumber(c) {
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

func numCharArrayToInt(s []string) int {
	var numString string
	for _, v := range s {
		numString = numString + v
	}
	num, err := strconv.Atoi(numString)
	util.Check(err)
	return num
}
