package main

import (
	"fmt"
	"strconv"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Scan(&N)

		numbers := make([]int64, N)
		strs := make([]string, N)

		for i := 0; i < N; i++ {
			fmt.Scan(&numbers[i])
			strs[i] = strconv.FormatInt(numbers[i], 10)
		}

		count := 0
		for i := 1; i < N; i++ {
			if numbers[i] > numbers[i-1] {
				continue
			}

			if len(strs[i]) == len(strs[i-1]) {
				numbers[i] = numbers[i] * 10
				strs[i] = strs[i] + "0"
				count++
			}

			less := false
			equal := true
			for j := 0; j < len(strs[i-1]); j++ {
				if !less && j < len(strs[i]) && strs[i][j] < strs[i-1][j] {
					less = true
				}

				if equal && j < len(strs[i]) && strs[i][j] != strs[i-1][j] {
					equal = false
				}

				if j < len(strs[i]) {
					continue
				}

				if less {
					strs[i] = strs[i] + "0"
					numbers[i] = numbers[i] * 10
					count++

				} else {
					if j != len(strs[i-1])-1 {
						strs[i] = strs[i] + string(strs[i-1][j])
						digit := int64(strs[i-1][j] - '0')
						numbers[i] = numbers[i]*10 + digit
						count++
					} else {
						if equal {
							if strs[i-1][j] == '9' {
								strs[i] = strs[i] + string(strs[i-1][j])
								digit := int64(strs[i-1][j] - '0')
								numbers[i] = numbers[i]*10 + digit
								count++
							} else {
								strs[i] = strs[i] + string(byte(strs[i-1][j]+1))
								digit := int64(strs[i-1][j] - '0' + 1)
								numbers[i] = numbers[i]*10 + digit
								count++
							}
						} else {
							strs[i] = strs[i] + string(strs[i-1][j])
							digit := int64(strs[i-1][j] - '0')
							numbers[i] = numbers[i]*10 + digit
							count++
						}
					}

				}
			}
			if less {
				strs[i] = strs[i] + "0"
				numbers[i] = numbers[i] * 10
				count++
			}

			if numbers[i] == numbers[i-1] {
				numbers[i] = numbers[i] * 10
				strs[i] = strs[i] + "0"
				count++
			}

		}

		// fmt.Println(strs)
		fmt.Printf("Case #%d: %d\n", t+1, count)

	}

}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
