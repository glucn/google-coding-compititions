package main

// TODO: Test Set 3 runtime error??

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T int

	if scanner.Scan() {
		input := scanner.Text()
		T = getInt(input)
	}

	for t := 0; t < T; t++ {
		var N int
		var existingPath string
		if scanner.Scan() {
			input := scanner.Text()
			N = getInt(input)
		}
		_ = N

		if scanner.Scan() {
			existingPath = scanner.Text()
		}

		output := make([]byte, len(existingPath))
		for i := range existingPath {
			if existingPath[i] == 'S' {
				output[i] = 'E'
			} else {
				output[i] = 'S'
			}
		}

		fmt.Printf("Case #%d: %s\n", t+1, string(output[:]))

	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}
