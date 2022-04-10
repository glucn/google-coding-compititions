package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)
	print := 1000000

	for t := 0; t < T; t++ {
		var min = []int{print, print, print, print}
		for i := 0; i < 3; i++ {
			for j := 0; j < 4; j++ {
				var U int
				fmt.Scan(&U)
				if min[j] > U {
					min[j] = U
				}
			}
		}

		sum := min[0] + min[1] + min[2] + min[3]

		if sum < print {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", t+1)
			continue
		}

		for i := 0; i < 4; i++ {
			if (sum - print) < min[i] {
				min[i] -= (sum - print)
				break
			}
			sum -= min[i]
			min[i] = 0
		}

		fmt.Printf("Case #%d: %d %d %d %d\n", t+1, min[0], min[1], min[2], min[3])
	}
}
