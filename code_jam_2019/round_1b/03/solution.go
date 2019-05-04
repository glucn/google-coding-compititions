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
			input := scanner.Text()
			s := strings.Split(input, " ")
			N = getInt(s[0])
			K = getInt(s[1])
		}

		C := make([]int, N)
		if scanner.Scan() {
			input := scanner.Text()
			s := strings.Split(input, " ")
			for i := 0; i < N; i++ {
				C[i] = getInt(s[i])
			}
		}

		D := make([]int, N)
		if scanner.Scan() {
			input := scanner.Text()
			s := strings.Split(input, " ")
			for i := 0; i < N; i++ {
				D[i] = getInt(s[i])
			}
		}

		//startTime := time.Now()
		maxC := findMax(C)
		maxD := findMax(D)

		count := 0
		for i := 0; i < len(maxC); i++ {
			if abs(maxC[i]-maxD[i]) <= K {
				count++
			}
		}

		//endTime := time.Now()
		//fmt.Println(endTime.UnixNano() - startTime.UnixNano())
		fmt.Printf("Case #%d: %d\n", t+1, count)

	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findMax(num []int) []int {
	l := len(num)
	total := l * (l + 1) / 2
	result := make([]int, total)

	for i := 0; i < l; i++ {
		result[i] = num[i]
	}

	start := 1
	baseIndex := 1
	resultIndex := len(num)
	resultReadIndex := 0

	for resultIndex < total {
		result[resultIndex] = max(result[resultReadIndex], num[baseIndex])
		baseIndex++
		if baseIndex != l {
			resultReadIndex++
		} else {
			start++
			baseIndex = start
			resultReadIndex += 2
		}
		resultIndex++
	}

	//fmt.Println(result)
	return result

}
