package main

import (
	"fmt"
	"sort"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Scan(&N)

		dices := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&dices[i])
		}

		sort.Ints(dices)

		straight := 0
		for _, d := range dices {
			if d > straight {
				straight++
			}
		}

		fmt.Printf("Case #%d: %d \n", t+1, straight)
	}
}
