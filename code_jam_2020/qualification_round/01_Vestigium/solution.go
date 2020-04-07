package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TODO: clean up the code
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T int

	if scanner.Scan() {
		input := scanner.Text()
		T = getInt(input)
	}

	for t := 0; t < T; t++ {
		var N int
		if scanner.Scan() {
			input := scanner.Text()
			N = getInt(input)
		}

		var k, r, c int

		colMap := make([]map[int]struct{}, N)
		for i := range colMap {
			colMap[i] = make(map[int]struct{})
		}
		colRepeated := make([]bool, N)

		for n := 0; n < N; n++ {
			inputRow := make([]int, N)
			if scanner.Scan() {
				input := scanner.Text()
				numbers := strings.Split(input, " ")
				for i := range numbers {
					inputRow[i] = getInt(numbers[i])
				}
			}

			// trace
			k += inputRow[n]

			// row repeat
			rowMap := make(map[int]struct{})
			for i := range inputRow {
				if _, found := rowMap[inputRow[i]]; found {
					r++
					break
				}
				rowMap[inputRow[i]] = struct{}{}
			}

			// col repeat
			for i := range inputRow {
				if colRepeated[i] {
					continue
				}
				if _, found := colMap[i][inputRow[i]]; found {
					c++
					colRepeated[i] = true
					continue
				}
				colMap[i][inputRow[i]] = struct{}{}
			}
		}

		fmt.Printf("Case #%d: %d %d %d\n", t+1, k, r, c)

	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}
