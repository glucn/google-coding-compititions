package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// TODO: Test Set 2 Runtime Error

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T int

	if scanner.Scan() {
		input := scanner.Text()
		T = getInt(input)
	}

	for t := 0; t < T; t++ {
		var N int
		var P int

		if scanner.Scan() {
			input := scanner.Text()
			s := strings.Split(input, " ")
			N = getInt(s[0])
			P = getInt(s[1])
		}

		//fmt.Println("N: ", N, " P: ", P)
		skills := make([]int, N)

		if scanner.Scan() {
			input := scanner.Text()
			s := strings.Split(input, " ")

			for j := 0; j < N; j++ {
				skills[j] = getInt(s[j])
			}
		}

		sort.Ints(skills)
		//fmt.Println("skills: ", skills)

		best := math.MaxInt64
		for i := 0; i < N-P+1; i++ {
			sum := 0
			for j := i; j < i+P; j++ {
				sum += skills[j]
			}
			//fmt.Println("sum: ", sum)
			training := skills[i+P-1]*P - sum
			//fmt.Println("training: ", training)

			if training < best {
				best = training
			}
		}
		fmt.Printf("Case #%d: %d\n", t+1, best)
	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}
