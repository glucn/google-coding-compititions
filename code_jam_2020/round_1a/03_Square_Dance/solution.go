package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This was the code I submitted during round 1a,
// it got MLE(memory limit exceeded) in the invisible test set 2
// TODO: clean up the code
// TODO: solve the pproblem for test set 2
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T int
	if scanner.Scan() {
		T = getInt(scanner.Text())
	}

	for t := 0; t < T; t++ {
		var R, C int
		if scanner.Scan() {
			input := strings.Split(scanner.Text(), " ")
			R = getInt(input[0])
			C = getInt(input[1])
		}

		cell := make([][]int, R)
		neighbors := make([][]neighbor, R)

		// total := 0
		for r := 0; r< R; r++ {
			cell[r] = make([]int, C)
			neighbors[r] = make([]neighbor, C)

			if scanner.Scan() {
				input := strings.Split(scanner.Text(), " ")
				for c := 0; c < C; c++ {
					neighbors[r][c] = neighbor{
						up: r-1,
						down: r+1,
						left: c-1,
						right: c+1,
					}
					cell[r][c] = getInt(input[c])
					// total += cell[r][c]
				}
			}
		}
		// fmt.Println("total", total)

		sum := 0
		eliminate := make([][]bool, R)
		for i := range eliminate {
			eliminate[i] = make([]bool, C)
		}
		terminate := false
		for !terminate {
			// copy neibourgh
			for r := 0; r < R; r++ {
				for c := 0; c < C; c++ {
					if eliminate[r][c] {
						if c != C-1 {
							neighbors[r][c+1].left = neighbors[r][c].left
						}
						if r != R-1 {
							neighbors[r+1][c].up = neighbors[r][c].up 
						}
					}
				}
			}
			for r := R-1; r >= 0; r-- {
				for c := C-1; c >= 0; c-- {
					if eliminate[r][c] {
						if c != 0 {
							neighbors[r][c-1].right = neighbors[r][c].right
						}
						if r != 0 {
							neighbors[r-1][c].down = neighbors[r][c].down 
						}
					}
				}
			}

			terminate = true
			// calculate new round
			for r := 0; r < R; r++ {
				for c := 0; c < C; c++ {
					if eliminate[r][c] {
						continue
					}

					// interest level of the competition
					sum += cell[r][c]

					sumNei := 0
					countNei := 0
					if neighbors[r][c].left != -1 {
						sumNei += cell[r][neighbors[r][c].left]
						countNei++
					}
					if neighbors[r][c].right != C {
						sumNei += cell[r][neighbors[r][c].right]
						countNei++
					}
					if neighbors[r][c].up != -1 {
						sumNei += cell[neighbors[r][c].up][c]
						countNei++
					}
					if neighbors[r][c].down != R {
						sumNei += cell[neighbors[r][c].down][c]
						countNei++
					}
					if countNei == 0 {
						continue
					}
					// fmt.Println(r, c, cell[r][c], sumNei, countNei, float64(sumNei)/float64(countNei))
					if float64(sumNei)/float64(countNei) > float64(cell[r][c]) {
						eliminate[r][c] = true
						terminate = false
						// fmt.Println(r, c, "eliminated")
					}
				}
			}

			// fmt.Println(eliminate)

		}

		

		fmt.Printf("Case #%d: %d\n", t+1, sum)
	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}

type neighbor struct {
	left, right, up, down int
}

