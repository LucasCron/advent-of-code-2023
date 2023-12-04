package main

import (
	"fmt"

	"lucascron.com/advent-of-code-2023/util"
)

var (
	files = [1]string{
		"example.txt",
	}
)

func main() {
	for _, path := range files {
		lines := util.ReadTextFileToArray(path)
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}
