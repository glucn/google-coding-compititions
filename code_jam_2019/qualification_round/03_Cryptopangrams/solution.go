package main

// TODO: Test Set 2 runtime error (number size)

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
		var N, L int
		if scanner.Scan() {
			input := scanner.Text()
			s := strings.Split(input, " ")
			N = getInt(s[0])
			L = getInt(s[1])
		}

		text := make([]int, L)

		if scanner.Scan() {
			input := scanner.Text()
			numbers := strings.Split(input, " ")
			for i, n := range numbers {
				text[i] = getInt(n)
			}
		}

		var n int
		numbers := make([]bool, N+1) // false: prime, true: not prime

		for i := 2; i < N; i++ {
			if numbers[i] {
				continue
			}
			for j := 2; j < N/i; j++ {
				numbers[j*i] = true
			}
			if text[0]%i == 0 {
				n = i
				break
			}
		}

		fmt.Printf("Case #%d: %s\n", t+1, decode(text, n))
	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}

func insert(s []int, num int) (res []int) {
	for i, v := range s {
		if v == num {
			return s
		}
		if v > num {
			res := append(s, 0)
			copy(res[i+1:], res[i:])
			res[i] = num
			return res
		}
	}
	return append(s, num)
}

func findIndex(s []int, num int) int {
	for i, ss := range s {
		if ss == num {
			return i
		}
	}
	return -1
}

func decode(input []int, init int) string {
	cipher := []int{init, input[0] / init}
	output := make([]int, len(input)+1)

	if input[1]%init == 0 {
		output[0] = input[0] / init
		output[1] = init
	} else {
		output[0] = init
		output[1] = input[0] / init
	}

	for i := 2; i <= len(input); i++ {
		output[i] = input[i-1] / output[i-1]
		if len(cipher) < 26 {
			cipher = insert(cipher, output[i])
		}
	}

	outputStr := make([]byte, len(output))

	for i, o := range output {
		index := findIndex(cipher, o)
		outputStr[i] = byte(index + 65)
	}

	return string(outputStr[:])
}
