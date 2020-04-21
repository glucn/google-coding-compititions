package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	half := 1000000000

	rand.Seed(time.Now().UnixNano())

	maxR := half - 5
	minR := half - 5

	R := rand.Intn(maxR-minR+1) + minR

	fmt.Println("R", R)

	min := -half + R
	max := half - R

	X := rand.Intn(max-min+1) + min
	Y := rand.Intn(max-min+1) + min
	fmt.Println("Center", X, Y)

	for {
		var inputX, inputY int
		fmt.Scan(&inputX)
		fmt.Scan(&inputY)

		if inputX == X && inputY == Y {
			fmt.Println("CENTER")
			return
		}

		if pow2(inputX-X)+pow2(inputY-Y) <= R*R {
			fmt.Println("HIT")
		} else {
			fmt.Println("MISS")
		}
	}
}

func pow2(a int) int {
	return a * a
}
