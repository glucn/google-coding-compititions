package main

import (
	"fmt"
	"sort"
)

const inputSize = 10000

// TODO: clean up the code, some variables can have better names
func main() {

	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var U int
		fmt.Scan(&U)

		firstDigits := make(map[byte]int)
		digits := make(map[byte]struct{})

		for i:=0; i< inputSize; i++ {
			var Q int
			var R string
			fmt.Scan(&Q)
			fmt.Scan(&R)

			c := R[0]
			if _, found := firstDigits[c]; found {
				firstDigits[c] = firstDigits[c] + 1
			} else {
				firstDigits[c] = 1
			}

			if len(digits) < 10 {
				for i := range R {
					if _, found := digits[R[i]]; !found {
						digits[R[i]] = struct{}{}
					}
				}
			}
		}


		var dd []digit
		for k, v := range firstDigits {
			dd = append(dd, digit{c: k, count: v})
			delete(digits, k)
		}
		var zero byte
		for k := range digits {
			// there should be only one key left, which is the charactor for 0
			zero = k
		}

		sort.Sort(byCount(dd))

		var output []byte
		output = append(output, zero)
		for _, d := range dd {
			output = append(output, d.c)
		}


		fmt.Printf("Case #%d: %s\n", t+1, string(output))
		
		
	}
}

type digit struct {
	c byte
	count int
}

type byCount []digit

func (a byCount) Len() int           { return len(a) }
func (a byCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCount) Less(i, j int) bool { return a[i].count > a[j].count }

