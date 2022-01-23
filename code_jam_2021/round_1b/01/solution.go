package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var A, B, C int
		fmt.Scan(&A)
		fmt.Scan(&B)
		fmt.Scan(&C)

		fmt.Println(A, B, C)
		fmt.Printf("Case #%d: %d %d %d %d\n", t+1, 0, 0, 0, 0)
	}
}
