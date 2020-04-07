package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math/rand"
)

// TODO: clean up the code
// TODO: got WA in test set 3
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T, B int

	if scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		T, B = getInt(input[0]), getInt(input[1])
	}

	for t := 0; t < T; t++ {
		
		ans := make([]bit, B)
		for i := 0; i < 5; i++ {
			fmt.Println(i+1)
			if scanner.Scan() {
				ans[i] = fromInt(getInt(scanner.Text()))
			}
		}
		for i := B-1; i> B-5-1; i-- {
			fmt.Println(i+1)
			if scanner.Scan() {
				ans[i] = fromInt(getInt(scanner.Text()))
			}
		}

		variations := make([]answer, 4)  //0-origin, 1-reverse, 2-complement, 3-r+c
		variations[0] = answer{data: ans}
		variations[1] = reverse(variations[0])
		variations[2] = complement(variations[0])
		variations[3] = reverse(variations[2])

		current := 0

		checkpoints := make(map[int]bit)
		determined := false
		i := 0
		head := 5
		tail := B-5-1

		for tail >= head {

			if i % 10 == 0 {
				determined = false
				checkpoints[getCheckpoint(variations[0].data, variations[1].data)] = unknown
				checkpoints[getCheckpoint(variations[0].data, variations[2].data)] = unknown
				checkpoints[getCheckpoint(variations[0].data, variations[3].data)] = unknown
			}

			// fmt.Println(variations)
			// fmt.Println(checkpoints)
			// fmt.Println("current is ", current)
			// fmt.Println("head: ", head, ", tail: ", tail)

			if !determined {
				// check complement first as it's always different
				for k := range checkpoints {
					fmt.Println(k+1)
					i++
					if scanner.Scan() {
						checkpoints[k] = fromInt(getInt(scanner.Text()))
					}
				}
				
				for j := 0; j < 4; j++ {
					match := true
					for k, v := range checkpoints {
						if variations[j].data[k] != v {
							match = false
						}
					}
					if match {
						current = j
						determined = true

						for k := 0; k < B; k++ {
							if variations[current].data[k] == unknown {
								head = k
								break
							}
						}

						for k := B-1; k >= 0; k-- {
							if variations[current].data[k] == unknown {
								tail = k
								break
							}
						}


						break
					}
				}

			} else {
				if head >= (B-tail) {
					// query tail
					// fmt.Println("query tail")
					fmt.Println(tail+1)
					i++
					

					var d bit
					if scanner.Scan() {
						d = fromInt(getInt(scanner.Text()))
					}
					switch current {
					case 0:
						variations[0].data[tail] = d
						variations[1].data[B-1-tail] = d
						variations[2].data[tail] = not(d)
						variations[3].data[B-1-tail] = not(d)
					case 1:
						variations[1].data[tail] = d
						variations[0].data[B-1-tail] = d
						variations[3].data[tail] = not(d)
						variations[2].data[B-1-tail] = not(d)
					case 2:
						variations[2].data[tail] = d
						variations[3].data[B-1-tail] = d
						variations[0].data[tail] = not(d)
						variations[1].data[B-1-tail] = not(d)
					case 3:
						variations[3].data[tail] = d
						variations[2].data[B-1-tail] = d
						variations[1].data[tail] = not(d)
						variations[0].data[B-1-tail] = not(d)
						
					}

					tail--
				} else {
					// query head
					// fmt.Println("query head")
					fmt.Println(head+1)
					i++
					

					var d bit
					if scanner.Scan() {
						d = fromInt(getInt(scanner.Text()))
					}
					switch current {
					case 0:
						variations[0].data[head] = d
						variations[1].data[B-1-head] = d
						variations[2].data[head] = not(d)
						variations[3].data[B-1-head] = not(d)
					case 1:
						variations[1].data[head] = d
						variations[0].data[B-1-head] = d
						variations[3].data[head] = not(d)
						variations[2].data[B-1-head] = not(d)
					case 2:
						variations[2].data[head] = d
						variations[3].data[B-1-head] = d
						variations[0].data[head] = not(d)
						variations[1].data[B-1-head] = not(d)
					case 3:
						variations[3].data[head] = d
						variations[2].data[B-1-head] = d
						variations[1].data[head] = not(d)
						variations[0].data[B-1-head] = not(d)
						
					}

					head++
				}
			}
		}


		fmt.Println(variations[current].String())

		if scanner.Scan() {
			if scanner.Text() != "Y" {
				// failed, exit
				return
			}
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

type bit int
const (
	unknown bit = iota
	zero
	one
)

func fromInt(x int) bit {
	if x == 0 {
		return zero
	}
	if x == 1 {
		return one
	}
	return unknown
}

func not(d bit) bit {
	if d == one {
		return zero
	}
	if d == zero {
		return one
	}
	return unknown
}

type answer struct {
	data []bit
}

func(a answer) String() string {
	builder := make([]byte, len(a.data))
	for k := 0; k < len(a.data); k++ {
		if a.data[k] == zero {
			builder[k] = '0'
		} else if a.data[k] == one {
			builder[k] = '1'
		} else {
			builder[k] = 'x'
		}
	}
	return string(builder)
}

func(a answer) firstUnknown() int {
	for i := range a.data {
		if a.data[i] == unknown {
			return i
		}
	}
	return len(a.data)
}

func reverse(a answer) answer {
	data := make([]bit, len(a.data))
	for i := range a.data {
		data[len(a.data)-i-1] = a.data[i]
	}
	return answer{
		data: data, 

	}
}

func complement(a answer) answer {
	data := make([]bit, len(a.data))
	for i := range a.data {
		if a.data[i] == zero {
			data[i] = one
		} else if a.data[i] == one {
			data[i] = zero
		}
	}
	return answer{
		data: data, 

	}
}

func divergent(B int, answerQueue []answer) int {
	start := rand.Intn(B)
	for i := 0; i < B; i++ {
		for _, ans := range answerQueue {
			ii := (start+i)%B
			if ans.data[ii] != answerQueue[0].data[ii] {
				return ii
			}
		}
	}

	return -1
}


func getCheckpoint(a, b []bit) int {
	for i := range a {
		if a[i] == unknown ||  b[i] == unknown {
			break
		}
		if a[i] != b[i] {
			return i
		}
	}
	return 0
}