package main

import (
	"bufio"
	"fmt"
	"os"
)

const LINES = 100
const CHARACTERS = 10000

// This solution got Wrong Answer
// TODO: implement the correct solution refering to the analysis
func main() {
	var T int
	var P int
	fmt.Scan(&T)
	fmt.Scan(&P)

	scanner := bufio.NewScanner(os.Stdin)

	for t := 0; t < T; t++ {
		lines := make([]string, LINES)

		for i := 0; i < LINES; i++ {
			if scanner.Scan() {
				lines[i] = scanner.Text()
			}
		}

		correctCounts := make([]int, LINES)
		for i, line := range lines {
			for _, s := range line {
				if s == '1' {
					correctCounts[i] = correctCounts[i] + 1
				}
			}
		}

		cheaterCountAbs := CHARACTERS
		cheaterIndex := -1

		for i := range correctCounts {
			if abs(correctCounts[i]-CHARACTERS/2) < cheaterCountAbs {
				cheaterCountAbs = abs(correctCounts[i] - CHARACTERS/2)
				cheaterIndex = i
			}
		}

		fmt.Printf("Case #%d: %d\n", t+1, cheaterIndex+1)
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
