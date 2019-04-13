package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TODO: test set 2 is TLE with this solution

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T int

	if scanner.Scan() {
		input := scanner.Text()
		T = getInt(input)
	}

	for t := 0; t < T; t++ {
		var R, C int
		if scanner.Scan() {
			input := scanner.Text()
			s := strings.Split(input, " ")
			R = getInt(s[0])
			C = getInt(s[1])
		}

		visited := make([][]bool, R)

		for i := 0; i < R; i++ {
			visited[i] = make([]bool, C)
		}

		path, err := visitAll(visited)
		if err != nil {
			fmt.Printf("Case #%d: %s\n", t+1, err.Error())
		} else {
			fmt.Printf("Case #%d: POSSIBLE\n", t+1)
			fmt.Printf("1 1\n")
			for i := range path {
				fmt.Println(path[i])
			}
		}
	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}

func visitAll(visited [][]bool) ([]string, error) {
	return visit(visited, 0, 0)
}

func visit(visited [][]bool, x, y int) ([]string, error) {
	//fmt.Println("visiting", x, y, visited)
	visited[x][y] = true

	count := 0

	for i := 0; i < len(visited); i++ {

		for j := 0; j < len(visited[i]); j++ {
			//fmt.Println("try", i, j)
			if visited[i][j] {
				count++
				continue
			}
			if i == x || j == y || i-j == x-y || i+j == x+y {
				continue
			}
			path, err := visit(visited, i, j)
			if err != nil {
				continue
			}
			return append([]string{fmt.Sprintf("%d %d", i+1, j+1)}, path...), nil
		}
	}
	if count == len(visited)*len(visited[0]) {
		return []string{}, nil
	}
	visited[x][y] = false
	return nil, fmt.Errorf("IMPOSSIBLE")
}
