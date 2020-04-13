package main

import (
	"fmt"
)

// This was the code I submitted during round 1a,
// it got MLE(memory limit exceeded) in the invisible test set 2
// TODO: clean up the code
// TODO: solve the pproblem for test set 2
func main() {

	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var R, C int
		fmt.Scan(&R)
		fmt.Scan(&C)

		cell := make([][]int, R)
		up := make([][]int, R)
		down := make([][]int, R)
		left := make([][]int, R)
		right := make([][]int, R)

		for r := 0; r < R; r++ {
			cell[r] = make([]int, C)
			up[r] = make([]int, C)
			down[r] = make([]int, C)
			left[r] = make([]int, C)
			right[r] = make([]int, C)

			for c := 0; c < C; c++ {
				up[r][c] = r - 1
				down[r][c] = r + 1
				left[r][c] = c - 1
				right[r][c] = c + 1

				fmt.Scan(&cell[r][c])
			}
		}

		sum := 0
		eliminate := make([][]int, R)
		for i := range eliminate {
			eliminate[i] = make([]int, C)
		}

		check := make(map[node]struct{})
		for r := 0; r < R; r++ {
			for c := 0; c < C; c++ {
				check[node{x: r, y: c}] = struct{}{}
			}
		}

		round := 1
		for {
			// fmt.Println("round start")
			// fmt.Println("left", left)
			// fmt.Println("right", right)

			for r := 0; r < R; r++ {
				for c := 0; c < C; c++ {
					if eliminate[r][c] == 0 {
						// interest level of the competition
						sum += cell[r][c]
					}
				}
			}

			nextCheck := make(map[node]struct{})

			for n := range check {
				r := n.x
				c := n.y
				if eliminate[r][c] != 0 {
					continue
				}
				sumNei := 0
				countNei := 0
				if left[r][c] != -1 {
					sumNei += cell[r][left[r][c]]
					countNei++
				}
				if right[r][c] != C {
					sumNei += cell[r][right[r][c]]
					countNei++
				}
				if up[r][c] != -1 {
					sumNei += cell[up[r][c]][c]
					countNei++
				}
				if down[r][c] != R {
					sumNei += cell[down[r][c]][c]
					countNei++
				}
				if countNei == 0 {
					continue
				}
				if cell[r][c]*countNei < sumNei {
					eliminate[r][c] = round

					if left[r][c] != -1 {
						nextCheck[node{x: r, y: left[r][c]}] = struct{}{}
					}

					if right[r][c] != C {
						nextCheck[node{x: r, y: right[r][c]}] = struct{}{}
					}

					if up[r][c] != -1 {
						nextCheck[node{x: up[r][c], y: c}] = struct{}{}
					}

					if down[r][c] != R {
						nextCheck[node{x: down[r][c], y: c}] = struct{}{}
					}
				}
			}

			// fmt.Println("eliminate", eliminate)
			// fmt.Println("nextCheck", nextCheck)

			for n := range check {
				r := n.x
				c := n.y
				if eliminate[r][c] == round {
					if left[r][c] != -1 {
						right[r][left[r][c]] = right[r][c]
					}

					if right[r][c] != C {
						left[r][right[r][c]] = left[r][c]
					}

					if up[r][c] != -1 {
						down[up[r][c]][c] = down[r][c]
					}

					if down[r][c] != R {
						up[down[r][c]][c] = up[r][c]
					}
				}
			}

			check = make(map[node]struct{})
			for k, v := range nextCheck {
				if eliminate[k.x][k.y] == 0 {
					check[k] = v
				}
			}
			if len(check) == 0 {
				break
			}

			round++
		}

		fmt.Printf("Case #%d: %d\n", t+1, sum)
	}
}

type node struct {
	x, y int
}

type neighbor struct {
	left, right, up, down int
}
