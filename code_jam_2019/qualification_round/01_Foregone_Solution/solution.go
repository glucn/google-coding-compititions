package main

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
		var input string

		if scanner.Scan() {
			input = scanner.Text()
		}

		output1 := make([]byte, len(input))
		output2 := make([]byte, len(input))
		for i := range input {
			d := int(input[i]) - 48
			switch d {
			case 7:
				output1[i] = '2'
				output2[i] = '5'
			case 8:
				output1[i] = '3'
				output2[i] = '5'
			case 9:
				output1[i] = '3'
				output2[i] = '6'
			default:
				output1[i] = byte(d/2 + 48)
				output2[i] = byte(d - d/2 + 48)
			}
		}

		fmt.Printf("Case #%d: %s %s\n", t+1, string(output1[:]), string(output2[:]))

	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}
