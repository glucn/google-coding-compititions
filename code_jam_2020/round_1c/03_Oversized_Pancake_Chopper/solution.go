package main

import (
	"fmt"
	"sort"
)

// Got WA in test set 1
// TODO: debug
// TODO: optimize for test set 3
func main() {

	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N, D int
		fmt.Scan(&N)
		fmt.Scan(&D)

		A := make([]int, N)
		for i := range A {
			fmt.Scan(&A[i])
		}

		sort.Ints(A)

		minCut := D - 1
		for i := range A {
			for j := 1; j <= D; j++ {
				if A[i] % j != 0 {
					continue
				}
				targetSize := fraction{numerator: A[i], denominator: j}

				if okay, countFullyUsed := isPossible(A, targetSize, D); !okay {
					continue
				} else {
					// fmt.Println(targetSize, countFullyUsed)
					cut := D - countFullyUsed
					if cut < minCut {
						minCut = cut
					}
				}
			}
		}

		fmt.Printf("Case #%d: %d\n", t+1, minCut)
		
	}
}

func isPossible(slices []int, targetSize fraction, D int) (isPossible bool, countFullyUsed int) {
	// fmt.Println("isPossible", targetSize, D)
	countAll := 0
	countFullyUsed = 0
	for _, s := range slices {
		countAll += s * targetSize.denominator / targetSize.numerator
		if (s * targetSize.denominator % targetSize.numerator == 0) {
			countFullyUsed++
		}
		if countAll == D {
			return true, countFullyUsed
		}
		if countAll > D {
			return true, countFullyUsed-1
		}
	}

	return false, -1
}

type fraction struct {
	numerator, denominator int
}