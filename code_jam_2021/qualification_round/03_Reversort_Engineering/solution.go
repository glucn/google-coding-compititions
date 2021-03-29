package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N int
		var C int
		fmt.Scan(&N)
		fmt.Scan(&C)

		if C < N-1 || C > N*(N+1)/2 - 1 {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", t+1)
			continue
		}

		L := make([]int, N)
		for i := range L {
			L[i] = i + 1
		}

		steps := make([]bool, N-1)
		for i := 0; i< N-1; i++ {
			// fmt.Println(C, i, N-i)
			if N-i <= C-(N-i-2) {
				steps[i] = true
				C -= N-i
			} else {
				C -= 1
			}
		}

		// fmt.Println(steps)

		for i := N-2; i>= 0; i-- {
			if steps[i] {
				reverse(L[i:])
			}
		}


		fmt.Printf("Case #%d: %v\n", t+1, toStr(L))
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func toStr(a []int) string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, " ")
}
