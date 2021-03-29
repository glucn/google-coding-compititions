package main

import (
	"fmt"
)

// This does not work for Test Set 3
// TODO: Change the logic to fit Test Set 3
func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var X, Y int
		var S string

		fmt.Scan(&X)
		fmt.Scan(&Y)
		fmt.Scan(&S)

		cost := 0

		var last byte = '?'

		for i := range S {
			if i == 0 {
				if S[i] != '?' {
					last = S[i]
				}
				continue
			}
			if S[i-1] == 'C' && S[i] == 'J' {
				cost += X
			}

			if S[i-1] == 'J' && S[i] == 'C' {
				cost += Y
			}

			if S[i] != '?' && S[i-1] == '?' && last != '?' && last != S[i] {
				if last == 'J' {
					cost += Y
				} else if last == 'C' {
					cost += X
				}
			}
			if S[i] != '?' {
				last = S[i]
			}
		}

		fmt.Printf("Case #%d: %d\n", t+1, cost)
	}
}
