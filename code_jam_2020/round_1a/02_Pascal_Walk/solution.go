package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const maxSize = 500 // the length of the walk cannot be longer than 500, so we can use up to 500 rows
var dp = make(map[int][]node)
var val = make([][]int, maxSize)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T int
	if scanner.Scan() {
		T, _ = strconv.Atoi(scanner.Text())
	}

	for i := 0; i < maxSize; i++ {
		val[i] = make([]int, i+1)
		val[i][0] = 1
		val[i][i] = 1
		if i > 1 {
			for j := 1; j < i; j++ {
				val[i][j] = val[i-1][j-1] + val[i-1][j]
			}
		}

	}

	dp[1] = []node{{row: 1, col: 1}}

	for t := 0; t < T; t++ {
		var N int
		if scanner.Scan() {
			N, _ = strconv.Atoi(scanner.Text())
		}

		if solution, found := dp[N]; found {
			// solved before
			fmt.Printf("Case #%d:\n", t+1)
			output(solution)
			continue
		}

		visited := make([][]bool, maxSize)
		for i := 0; i < maxSize; i++ {
			visited[i] = make([]bool, i+1)
		}

		sum := 1
		var path = []node{node{row: 1, col: 1}}
		visited[0][0] = true

		dfs(path, visited, N, sum, t)
	}
}

type node struct {
	row, col int
	val      int
}

type nodes []node

func (n nodes) Len() int           { return len(n) }
func (n nodes) Less(i, j int) bool { return n[i].val > n[j].val }
func (n nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

var moves [][]int = [][]int{
	[]int{-1, -1}, []int{-1, 0}, []int{0, -1}, []int{0, 1}, []int{1, 0}, []int{1, 1},
}

func dfs(path []node, visited [][]bool, N int, sum int, t int) bool {
	current := path[len(path)-1]

	var nextSteps []node

	for _, m := range moves {
		nn := node{row: current.row + m[0], col: current.col + m[1]}
		if nn.row > 0 && nn.col > 0 && nn.row <= maxSize && nn.col <= nn.row && !visited[nn.row-1][nn.col-1] {
			ss := sum + val[nn.row-1][nn.col-1]

			if _, found := dp[ss]; !found {
				// use dp to optimize performance in the case of multiple runs
				dp[ss] = append(append(path[:0:0], path...), nn)
			}

			if ss == N {
				// solved
				fmt.Printf("Case #%d:\n", t+1)
				output(append(path, nn))
				return true
			}

			if ss < N {
				nn.val = ss
				nextSteps = append(nextSteps, nn)
			}
		}
	}

	// prioritize nodes that has higher value to shorten the walk
	sort.Sort(nodes(nextSteps))

	for _, n := range nextSteps {
		visited[n.row-1][n.col-1] = true
		if dfs(append(path, n), visited, N, sum+val[n.row-1][n.col-1], t) {
			return true
		}
		visited[n.row-1][n.col-1] = false
	}

	return false
}

func output(path []node) {
	// sum := 0
	for _, n := range path {
		fmt.Println(n.row, n.col)
		// sum += val[n.row-1][n.col-1]
	}
	// fmt.Println("length", len(path))
	// fmt.Println("sum", sum)
}
