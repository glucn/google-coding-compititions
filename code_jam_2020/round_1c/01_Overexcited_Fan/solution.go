package main

import (
	"fmt"
)

func main() {

	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var X, Y int
		var M string
		fmt.Scan(&X)
		fmt.Scan(&Y)
		fmt.Scan(&M)

		if isPossible(X, Y, 0) {
			fmt.Printf("Case #%d: %d\n", t+1, 0)
			continue
		}

		possible := false
		for i, m := range M {
			switch m {
			case 'N':
				Y++
			case 'S':
				Y--
			case 'W':
				X--
			case 'E':
				X++
			}
			if isPossible(X, Y, i+1) {
				fmt.Printf("Case #%d: %d\n", t+1, i+1)
				possible = true
				break
			}
		}

		if !possible {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", t+1)
		}
		
	}
}

func isPossible(x, y int, time int) bool {
	return abs(x) + abs(y) <= time 
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}