package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T int

	if scanner.Scan() {
		input := scanner.Text()
		T = getInt(input)
	}

	for t := 0; t < T; t++ {
		var P, Q int
		if scanner.Scan() {
			input := scanner.Text()
			s := strings.Split(input, " ")
			P = getInt(s[0])
			Q = getInt(s[1])
		}

		//count := make([][]int, Q+1)
		//for i := 0; i <= Q; i++ {
		//	count[i] = make([]int, Q+1)
		//}

		countX := make([]int, Q+1)
		countY := make([]int, Q+1)
		for p := 0; p < P; p++ {
			var x, y int
			var direction string
			if scanner.Scan() {
				input := scanner.Text()
				s := strings.Split(input, " ")
				x = getInt(s[0])
				y = getInt(s[1])
				direction = s[2]
			}

			switch direction {
			case "N":
				//for i := 0; i <= Q; i++ {
				//	for j := 0; j <= Q; j++ {
				//		if j > y {
				//			count[i][j]++
				//		}
				//	}
				//}
				for i := y + 1; i <= Q; i++ {
					countY[i]++
				}
			case "S":
				//for i := 0; i <= Q; i++ {
				//	for j := 0; j <= Q; j++ {
				//		if j < y {
				//			count[i][j]++
				//		}
				//	}
				//}
				for i := 0; i < y; i++ {
					countY[i]++
				}

			case "W":
				//for i := 0; i <= Q; i++ {
				//	for j := 0; j <= Q; j++ {
				//		if i < x {
				//			count[i][j]++
				//		}
				//	}
				//}
				for i := 0; i < x; i++ {
					countX[i]++
				}
			case "E":
				//for i := 0; i <= Q; i++ {
				//	for j := 0; j <= Q; j++ {
				//		if i > x {
				//			count[i][j]++
				//		}
				//	}
				//}
				for i := x + 1; i <= Q; i++ {
					countX[i]++
				}
			}
		}

		//max := 0
		//tarX, tarY := 0, 0
		//for i := 0; i <= Q; i++ {
		//	for j := 0; j <= Q; j++ {
		//		if count[i][j] > max {
		//			max = count[i][j]
		//			tarX, tarY = i, j
		//		}
		//	}
		//}
		maxX, maxY, tarX, tarY := 0, 0, 0, 0
		for i := 0; i <= Q; i++ {
			if countX[i] > maxX {
				maxX = countX[i]
				tarX = i
			}
			if countY[i] > maxY {
				maxY = countY[i]
				tarY = i
			}
		}
		fmt.Printf("Case #%d: %d %d\n", t+1, tarX, tarY)
	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}
