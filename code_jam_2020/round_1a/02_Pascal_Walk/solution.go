package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var maxSize = 1 << 10
var dp = make(map[int][]node)
var val = make([][]int, maxSize)

// This was the code I submitted during round 1a,
// it got WA (wrong answer) in the invisible test set 3
// TODO: clean up the code
// TODO: solve it for test set 3
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T int

	if scanner.Scan() {
		T = getInt(scanner.Text())
	}

	for i := 0; i<maxSize; i++ {
		val[i] = make([]int, i+1)
		val[i][0] = 1
		val[i][i] = 1
		if i > 1 {
			for j := 1; j<i; j++ {
				val[i][j] = val[i-1][j-1] + val[i-1][j]
			}
		}
		
	}

	dp[1] = []node{{row: 1, col: 1}}

	// fmt.Println(val)

	for t := 0; t < T; t++ {
		var N int
		if scanner.Scan() {
			N = getInt(scanner.Text())
		}

		if solution, found := dp[N]; found {
			// solved
			fmt.Printf("Case #%d:\n", t+1)
			output(solution)
			continue
		}

		visited := make([][]bool, maxSize)
		for i := 0; i< maxSize; i++ {
			visited[i] = make([]bool, i+1)
		}

		sum := 1
		var path = []node{node{row: 1, col: 1}}
		visited[0][0] = true

		dfs(path, visited, N, sum, t)
	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}

type node struct {
	row, col int
}

var moves [][]int = [][]int{
	[]int{-1, -1}, []int{-1, 0}, []int{0, -1}, []int{0, 1}, []int{1, 0}, []int{1, 1},
}

func dfs(path []node, visited [][]bool, N int, sum int, t int) bool {
	// fmt.Println(path)
	current := path[len(path) - 1]

	for _, m := range moves {
		nn := node{row: current.row + m[0], col: current.col + m[1]}

		if nn.row > 0 && nn.col > 0 && nn.col <= nn.row && !visited[nn.row-1][nn.col-1] {
			ss := sum + val[nn.row-1][nn.col-1]

			if _, found := dp[ss]; !found {
				dp[ss] = append(append(path[:0:0], path...), nn)
			}
			
			if ss == N {
				// solved
				fmt.Printf("Case #%d:\n", t+1)
				output(append(path, nn))
				return true
			}
			if ss < N {
				visited[nn.row-1][nn.col-1] = true
				if dfs(append(path, nn), visited, N, sum + val[nn.row-1][nn.col-1], t) {
					return true
				}
				visited[nn.row-1][nn.col-1] = false
			}
		}
	}

	return false
}

func output(path []node) {
	// sum := 0
	// fmt.Println("length", len(path))
	for _, n := range path {
		fmt.Println(n.row, n.col)
		// sum += val[n.row-1][n.col-1]
	}
	// fmt.Println("sum", sum)
}

