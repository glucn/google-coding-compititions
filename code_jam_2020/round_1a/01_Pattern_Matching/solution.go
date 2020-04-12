package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T int
	if scanner.Scan() {
		T, _ = strconv.Atoi((scanner.Text()))
	}

	for t := 0; t < T; t++ {
		var N int
		if scanner.Scan() {
			N, _ = strconv.Atoi((scanner.Text()))
		}

		patterns := make([]string, N)

		for n := 0; n < N; n++ {
			if scanner.Scan() {
				patterns[n] = scanner.Text()
			}
		}

		prefixs := make([]string, N)
		suffixs := make([]string, N)
		mids := make([]string, N)

		for i, p := range patterns {
			sections := strings.Split(p, "*")
			prefixs[i] = sections[0]
			suffixs[i] = sections[len(sections)-1]

			if len(sections) > 2 {
				mids[i] = strings.Join(sections[1:len(sections)-1], "")
			}
		}

		var prefix, suffix string

		if p, ok := findSharedPattern(prefixs, strings.HasPrefix); ok {
			prefix = p
		} else {
			fmt.Printf("Case #%d: %s\n", t+1, "*")
			continue
		}

		if s, ok := findSharedPattern(suffixs, strings.HasSuffix); ok {
			suffix = s
		} else {
			fmt.Printf("Case #%d: %s\n", t+1, "*")
			continue
		}

		fmt.Printf("Case #%d: %s\n", t+1, prefix+strings.Join(mids, "")+suffix)

	}
}

func findSharedPattern(patterns []string, checker func(s1, s2 string) bool) (string, bool) {
	for i := 0; i < len(patterns); i++ {
		for j := 0; j < len(patterns); j++ {
			if !checker(patterns[i], patterns[j]) && !checker(patterns[j], patterns[i]) {
				return "", false
			}
		}
	}

	maxIndex := -1
	maxLen := -1
	for i, p := range patterns {
		if len(p) > maxLen {
			maxLen = len(p)
			maxIndex = i
		}
	}
	return patterns[maxIndex], true
}
