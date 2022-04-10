package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var R, C int
		fmt.Scan(&R)
		fmt.Scan(&C)

		if R < 2 || C < 2 {
			continue
		}

		fmt.Printf("Case #%d:\n", t+1)
		for r := 0; r < 2*R+1; r++ {
			for c := 0; c < 2*C+1; c++ {
				if r <= 1 && c <= 1 {
					fmt.Print(".")
				} else if r%2 == 0 && c%2 == 0 {
					fmt.Print("+")
				} else if r%2 == 1 && c%2 == 0 {
					fmt.Print("|")
				} else if r%2 == 0 && c%2 == 1 {
					fmt.Print("-")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	}
}
