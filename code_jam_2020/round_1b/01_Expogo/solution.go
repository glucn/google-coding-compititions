package main

import (
	"fmt"
)

func main() {

	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var X, Y int
		fmt.Scan(&X)
		fmt.Scan(&Y)

		fmt.Printf("Case #%d: %s\n", t+1, findPath(X, Y))
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func findPath(X, Y int) string {
	if (X+Y) % 2 == 0 {
		return "IMPOSSIBLE"
	}

	// fmt.Println("findPath", X, Y)

	switch {
	case X==1 && Y==0:
		return "E"
	case X==-1 && Y==0:
		return "W"
	case X==0 && Y==1:
		return "N"
	case X==0 && Y==-1:
		return "S"
	}
	
	if abs(X) % 2 == 1 {
		p := findPath((X+1)/2, Y/2)
		if p != "IMPOSSIBLE" {
			return "W" + p
		}
		p = findPath((X-1)/2, Y/2)
		if p != "IMPOSSIBLE" {
			return "E" + p
		}
		return "IMPOSSIBLE"
	} 

	p := findPath(X/2, (Y+1)/2)
	if p != "IMPOSSIBLE" {
		return "S" + p
	}
	p = findPath(X/2, (Y-1)/2)
	if p != "IMPOSSIBLE" {
		return "N" + p
	}
	return "IMPOSSIBLE"
}