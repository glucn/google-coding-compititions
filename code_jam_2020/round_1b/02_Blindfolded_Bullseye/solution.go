package main

import (
	"fmt"
)

// TODO: actually write the code
func main() {

	var T int
	var A, B int
	fmt.Scan(&T)
	fmt.Scan(&A)
	fmt.Scan(&B)

	for t := 0; t < T; t++ {
		
		var x, y int
		for {
			fmt.Println(x, y)

			var input string
			fmt.Scan(&input)

			switch input {
			case "CENTER":
				//success
				break
			case "HIT":
				// do something
				continue
			case "MISS":
				// do something else
				continue
			case "WRONG":
				// game over
				return
			default:
				// unknown input
				return
			}
		}

		fmt.Printf("Case #%d: %s\n", t+1, "IMPOSSIBLE")
	}
}
