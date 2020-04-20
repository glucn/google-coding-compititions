package main

import (
	"fmt"
)

const max = 1000

// TODO: solve the problem for test set 3, got MLE
func main() {
	dp := make(map[node]string)

	var queue []path
	queue = append(queue, path{output: "", x: 0, y: 0, round: 0})

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		// fmt.Println(p)

		xx := p.x + 1<<p.round
		if abs(xx) <= max {
			
			output := p.output + "E"
			queue = append(queue, path{output: output, x: xx, y: p.y, round: p.round+1 })
			if _, found := dp[node{x: xx, y: p.y}]; !found {
				dp[node{x: xx, y: p.y}] = output
			}
			
		}

		xx = p.x - 1<<p.round
		if abs(xx) <= max {
			output := p.output + "W"
			queue = append(queue, path{output: output, x: xx, y: p.y, round: p.round+1 })
			if _, found := dp[node{x: xx, y: p.y}]; !found {

				dp[node{x: xx, y: p.y}] = output
			}
		}

		yy := p.y + 1<<p.round
		if abs(yy) <= max {
			output := p.output + "N"
			queue = append(queue, path{output: output, x: p.x, y: yy, round: p.round+1 })
			if _, found := dp[node{x: p.x, y: yy}]; !found {

				dp[node{x: p.x, y: yy}] = output
			}
		}

		yy = p.y - 1<<p.round
		if abs(yy) <= max {
			output := p.output + "S"
			queue = append(queue, path{output: output, x: p.x, y: yy, round: p.round+1 })
			if _, found := dp[node{x: p.x, y: yy}]; !found {

				dp[node{x: p.x, y: yy}] = output

			}
		}

	}

	// fmt.Println(dp)

	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var X, Y int
		fmt.Scan(&X)
		fmt.Scan(&Y)

		if (X+Y) % 2 == 0 {
			// impossible
			fmt.Printf("Case #%d: %s\n", t+1, "IMPOSSIBLE")
			continue
		}

		if solution, found := dp[node{x: X, y: Y}]; found {
			fmt.Printf("Case #%d: %s\n", t+1, solution)
			continue
		}

		fmt.Printf("Case #%d: %s\n", t+1, "IMPOSSIBLE")
	}
}

type node struct {
	x, y int
}

type path struct {
	output string
	x, y int
	round uint
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}