package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// TODO: Test Set 2 Time Limit Exceeded

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T int

	if scanner.Scan() {
		input := scanner.Text()
		T = getInt(input)
	}

	for t := 0; t < T; t++ {
		var R int
		var C int

		if scanner.Scan() {
			input := scanner.Text()
			s := strings.Split(input, " ")
			R = getInt(s[0])
			C = getInt(s[1])
		}

		grid := make([][]byte, R)
		for i := range grid {
			grid[i] = make([]byte, C)
			if scanner.Scan() {
				input := scanner.Text()
				for j := 0; j < C; j++ {
					grid[i][j] = input[j]
				}
			}
		}

		initial, maxSingle := getTotalCost(grid)
		if initial != 0 {
			_, maxSingle = getUpdatedTotalCost(grid)
		}

		fmt.Printf("Case #%d: %d\n", t+1, maxSingle)
	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}

func getTotalCost(grid [][]byte) (int, int) {
	cost := 0
	maxSingle := 0
	for i := range grid {
		for j := range grid[i] {
			c := getCost(grid, i, j)

			cost += c
			if c > maxSingle {
				maxSingle = c
			}
		}
	}
	//fmt.Printf("getTotalCost: %d\n", cost)

	return cost, maxSingle
}

func getCost(grid [][]byte, x, y int) int {
	min := math.MaxInt64

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				cost := int(math.Abs(float64(x-i)) + math.Abs(float64(y-j)))
				if cost < min {
					min = cost
				}
			}

		}
	}
	//fmt.Printf("getCost: i %d, j %d: %d\n", x, y, min)
	return min
}

func getUpdatedTotalCost(grid [][]byte) (int, int) {
	total := math.MaxInt64
	var maxSingle int

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				continue
			}
			//fmt.Println("updating ", i, j)
			grid[i][j] = '1'
			c, m := getTotalCost(grid)
			if c == 0 {
				return 0, 0
			}
			if c < total {
				total = c
				maxSingle = m
			}
			grid[i][j] = '0'
		}
	}
	return total, maxSingle
}
