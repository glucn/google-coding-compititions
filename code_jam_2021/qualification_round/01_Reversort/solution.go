package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Scan(&N)

		L := make([]int, N)
		for i := range L {
			fmt.Scan(&L[i])
		}

		cost := 0

		for i := 0; i < N-1; i++ {
			minIndex := -1
			min := math.MaxInt64
			for j := i; j < N; j++ {
				if L[j] < min {
					minIndex = j
					min = L[j]
				}
			}

			cost += minIndex - i + 1

			reverse(L[i : minIndex+1])
			// fmt.Printf("%v\n", L)
			// fmt.Println("cost:", cost, ", i:", i, ", minIndex:", minIndex)

		}

		fmt.Printf("Case #%d: %d\n", t+1, cost)
	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
