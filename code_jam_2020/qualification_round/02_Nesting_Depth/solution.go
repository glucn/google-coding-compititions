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
		var inputRow []int
		if scanner.Scan() {
			input := scanner.Text()
			numbers := strings.Split(input, "")
			for i := range numbers {
				inputRow = append(inputRow, getInt(numbers[i])) 
			}
		}

		var builder []byte
		depth := 0

		for i := range inputRow {
			if depth < inputRow[i] {
				for j := 0; j< inputRow[i]-depth ; j++ {
					builder = append(builder, '(')
				}
			} else if depth > inputRow[i] {
				for j := 0; j< depth-inputRow[i] ; j++ {
					builder =  append(builder, ')')
				}
			}
			depth = inputRow[i]
			builder =  append(builder, byte('0'+inputRow[i]))
		}
		for j := 0; j< depth ; j++ {
			builder =  append(builder, ')')
		}


		fmt.Printf("Case #%d: %s\n", t+1, string(builder))

	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}
