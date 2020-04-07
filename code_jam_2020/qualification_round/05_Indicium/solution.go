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
		input := scanner.Text()
		T = getInt(input)
	}

	for t := 0; t < T; t++ {
		var N, K int
		if scanner.Scan() {
			input := strings.Split(scanner.Text(), " ")
			N = getInt(input[0])
			K = getInt(input[1])
		}

		// TODO: solve the problem


		fmt.Printf("Case #%d: %s\n", t+1, "")

	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}