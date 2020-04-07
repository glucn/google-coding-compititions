package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"bufio"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("1 20")

	ans := make([]string, 4) //0 - original, 1 - reverse, 2 - complement, 3 - r + c
	for i := 0; i< 20; i++ {
		d := rand.Intn(2)
		v := strconv.Itoa(d)
		c := strconv.Itoa(1-d)
		ans[0] = ans[0] + v
		ans[1] = v + ans[1]
		ans[2] = ans[2] + c
		ans[3] = c + ans[3]
	}
	fmt.Println(ans)

	count := 0
	current := 0
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if count > 0 && count % 10 == 0 {
			current = rand.Intn(4)
			fmt.Println("rotating, current is", current, ans[current])
		}

		if scanner.Scan() {
			
			input := scanner.Text()
			if len(input) == 20 {
				if input == ans[current] {
					fmt.Println("Y")
					break
				} else {
					fmt.Println("N")
					break
				}
			}
			count++
			index := getInt(input) - 1
			fmt.Println(ans[current][index]-'0')
		
		}

	}
	
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}