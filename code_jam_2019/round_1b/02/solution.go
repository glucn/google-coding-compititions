package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var invA = [][]float64{
	{0.1, -0.1, -0.05, 0, 0, 0.025},
	{-1.2, 1.2, 0.1, 0, 0, -0.05},
	{-0.4, -0.6, 1.2, 0, 0, -0.1},
	{1.6, -1.6, -0.8, 1, 0, -0.1},
	{-1.6, 1.6, 0.8, -1, 1, -0.4},
	{2.4, -0.4, -1.2, 0, -1, 0.6},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T int
	var W int

	if scanner.Scan() {
		input := scanner.Text()
		s := strings.Split(input, " ")
		T = getInt(s[0])
		W = getInt(s[1])
	}
	_ = W

	for t := 0; t < T; t++ {
		D := make([]int64, 6)
		for i := 0; i < 6; i++ {
			fmt.Println(i + 1)
			if scanner.Scan() {
				input := scanner.Text()
				D[i] = getInt64(input)
			}
		}

		R := make([]float64, 6)
		for i := 0; i < 6; i++ {
			for j := 0; j < 6; j++ {
				//fmt.Println(D[j])
				//fmt.Println(float32(D[j]))
				//fmt.Println(float32(D[j]) * invA[i][j])
				R[i] += float64(D[j]) * invA[i][j]
			}
		}
		fmt.Println(int64(R[0]), int64(R[1]), int64(R[2]), int64(R[3]), int64(R[4]), int64(R[5]))
		if scanner.Scan() {
			input := scanner.Text()
			if input == "-1" {
				break
			}
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

func getInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return n
}
