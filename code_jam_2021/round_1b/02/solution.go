package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N, A, B int
		fmt.Scan(&N)
		fmt.Scan(&A)
		fmt.Scan(&B)

		target := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&target[i])
		}

		fmt.Println(N, A, B)
		fmt.Println(target)
		fmt.Printf("Case #%d: %d\n", t+1, 0)
	}
}
