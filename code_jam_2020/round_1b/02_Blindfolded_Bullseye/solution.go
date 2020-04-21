package main

import (
	"fmt"
)

const half = 1000000000

// TODO: clean up the code, it's soo messy now
func main() {

	var T int
	var A, B int
	fmt.Scan(&T)
	fmt.Scan(&A)
	fmt.Scan(&B)

	for t := 0; t < T; t++ {
		foundCenter := false
		var x0, y0 int

	outer:
		// look for a HIT point
		for xx := -half + A; xx < half; xx += A {
			for yy := -half + A; yy < half; yy += A {
				fmt.Println(xx, yy)
				var input string
				fmt.Scan(&input)
				switch input {
				case "CENTER":
					//success by accident
					foundCenter = true
					break outer
				case "HIT":
					x0, y0 = xx, yy
					break outer
				}
			}
		}

		if foundCenter {
			// already got it, skip to the next test case
			continue
		}

		var xLeft, xRight, yUp, yDown int
		// look for the left-most HIT point
		left, right := -half, x0
		for left <= right {
			mid := left + (right-left)>>1

			fmt.Println(mid, y0)
			var input string
			fmt.Scan(&input)
			switch input {
			case "CENTER":
				//success by accident
				foundCenter = true
				break
			case "HIT":
				right = mid - 1
				continue
			case "MISS":
				left = mid + 1
				continue
			}
		}
		if foundCenter {
			continue
		}
		xLeft = left
		// fmt.Println("leftmost", xLeft, y0)

		// look for the right-most HIT point
		left, right = x0, half
		for left <= right {
			mid := left + (right-left)>>1

			fmt.Println(mid, y0)
			var input string
			fmt.Scan(&input)
			switch input {
			case "CENTER":
				//success by accident
				foundCenter = true
				break
			case "HIT":
				left = mid + 1
				continue
			case "MISS":
				right = mid - 1
				continue
			}
		}
		if foundCenter {
			continue
		}
		xRight = right
		// fmt.Println("rightmost", xRight, y0)

		// look for the up-most HIT point
		down, up := y0, half
		for down <= up {
			mid := down + (up-down)>>1

			fmt.Println(x0, mid)
			var input string
			fmt.Scan(&input)
			switch input {
			case "CENTER":
				//success by accident
				foundCenter = true
				break
			case "HIT":
				down = mid + 1
				continue
			case "MISS":
				up = mid - 1
				continue
			}
		}
		if foundCenter {
			continue
		}
		yUp = up
		// fmt.Println("upmost", x0, yUp)

		// look for the down-most HIT point
		down, up = -half, y0
		for down <= up {
			mid := down + (up-down)>>1

			fmt.Println(x0, mid)
			var input string
			fmt.Scan(&input)
			switch input {
			case "CENTER":
				//success by accident
				foundCenter = true
				break
			case "HIT":
				up = mid - 1
				continue
			case "MISS":
				down = mid + 1
				continue
			}
		}
		if foundCenter {
			continue
		}
		yDown = down
		// fmt.Println("downMost", x0, yDown)

		// look for the center
		// this could be wrong due to accuracy, need to check points around it
		xPossible, yPossible := xLeft+(xRight-xLeft)>>1, yDown+(yUp-yDown)>>1

	outer2:
		for xx := xPossible - 3; xx <= xPossible+3; xx++ {
			for yy := yPossible - 3; yy <= yPossible+3; yy++ {
				fmt.Println(xx, yy)
				var input string
				fmt.Scan(&input)
				switch input {
				case "CENTER":
					//success
					foundCenter = true
					break outer2
				}
			}
		}
	}
}
