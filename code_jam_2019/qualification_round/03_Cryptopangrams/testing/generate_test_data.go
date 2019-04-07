package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	T := 1000

	f, _ := os.Create("test.txt")
	defer f.Close()

	_, _ = f.WriteString(fmt.Sprintf("%d\n", T))

	for t := 0; t < T; t++ {
		N := rand.Intn(10001-101) + 101
		L := rand.Intn(100+1-25) + 25

		var cipher []int
		for i := 2; i <= N; i++ {
			if !IsPrime(i) {
				continue
			}
			if rand.Int()%26 == 0 {
				cipher = append(cipher, i)
			}
			if len(cipher) == 26 {
				break
			}
		}

		if len(cipher) < 26 {
			fmt.Println("not enough")
			t--
			continue
		}

		_, _ = f.WriteString(fmt.Sprintf("%d %d\n", N, L))

		origin := make([]int, L+1)
		originStr := make([]byte, L+1)
		output := make([]int, L)
		for i := range origin {
			n := rand.Intn(26)
			originStr[i] = byte(n + 65)
			origin[i] = cipher[n]

			if i != 0 {
				output[i-1] = origin[i-1] * origin[i]
			}
		}
		fmt.Println("origin")
		fmt.Println(string(originStr[:]))
		fmt.Println("output")
		fmt.Println(output)
		for i, o := range output {
			if i == 0 {
				_, _ = f.WriteString(fmt.Sprintf("%d", o))

			} else {
				_, _ = f.WriteString(fmt.Sprintf(" %d", o))
			}
		}
		if t != T-1 {
			_, _ = f.WriteString(fmt.Sprintf("\n"))
		}
	}

}

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
