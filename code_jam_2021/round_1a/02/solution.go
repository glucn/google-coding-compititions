package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var M int
		fmt.Scan(&M)

		primes := make([]Prime, M)

		sum := 0

		for m := 0; m < M; m++ {
			var P, N int
			fmt.Scan(&P)
			fmt.Scan(&N)
			primes[m] = Prime{prime: P, count: N}
			sum += P * N
		}

		score := 0
		// skip even numbers if there is no 2
		for j := sum - 1; j > 0; j-- {

			if okay(j, primes, sum) {
				score = j
				break
			}
		}

		// fmt.Println(primes, sum)
		fmt.Printf("Case #%d: %d\n", t+1, score)

	}

}

type Prime struct {
	prime int
	count int
}

func okay(number int, primes []Prime, target int) bool {
	// fmt.Println("input", number)
	original := number
	sum := 0
	used := make(map[int]int)

	for _, p := range primes {
		// fmt.Println("working on", p)
		for number%p.prime == 0 && used[p.prime] < p.count {
			// fmt.Println("inside", number, used)
			used[p.prime] = used[p.prime] + 1
			number = number / p.prime
			sum += p.prime
		}
	}

	return number == 1 && original+sum == target
}
