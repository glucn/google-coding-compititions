package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This was the code I submitted during round 1a, it only sovled the problem for test set 1
// TODO: clean up the code
// TODO: solve the problem for test set 2 & 3
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T int
	if scanner.Scan() {
		T = getInt(scanner.Text())
	}

	for t := 0; t < T; t++ {
		var N int
		if scanner.Scan() {
			N = getInt(scanner.Text())
		}

		patterns := make([]string, N)

		for n := 0; n<N; n++ {
			if scanner.Scan() {
				patterns[n] = scanner.Text()
			}
		}

		if ok, result := hasSameSuffix(patterns); ok {
			fmt.Printf("Case #%d: %s\n", t+1, result)
		} else {
			fmt.Printf("Case #%d: %s\n", t+1, "*")
		}

	
	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}

func hasSameSuffix(patterns []string) (bool, string) {

	
	for i := 0; i<len(patterns); i++ {
		s1 := patterns[i][1:]
		for j := 0; j<len(patterns); j++ {
			s2 := patterns[j][1:]
			if !strings.HasSuffix(s1, s2) && !strings.HasSuffix(s2, s1) {
				return false, ""
			}
		}
	}

	maxPatternIndex := -1
	maxLen := -1
	for i, p := range patterns {
		if len(p) > maxLen {
			maxLen = len(p)
			maxPatternIndex = i
		}
	}
	return true, patterns[maxPatternIndex][1:]
}
