package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N, Q int
		fmt.Scan(&N)
		fmt.Scan(&Q)

		answers := make([]Answer, N)
		for n := 0; n < N; n++ {
			var answerStr string
			var scr int
			var answer uint64

			fmt.Scan(&answerStr)
			fmt.Scan(&scr)

			for i := range answerStr {
				if answerStr[i] == 'T' {
					answer = setBit(answer, uint(Q-i-1))
				}
			}

			answers[n] = Answer{answer: answer, score: scr}
		}

		possibleAnswers := make(map[uint64]struct{})

		for p := uint64(0); p < (1 << uint(Q)); p++ {
			// fmt.Println("p", p)
			// fmt.Println("score", Q-bits.OnesCount64(p^answers[0].answer))
			// score := Q - bits.OnesCount64(p^answers[0].answer)

			okay := true
			for i := range answers {
				score := Q - bits.OnesCount64(p^answers[i].answer)
				if score != answers[i].score {
					okay = false
					break
				}
			}

			if okay {
				possibleAnswers[p] = struct{}{}
			}
		}

		var output string
		for i := 0; i < Q; i++ {
			countT := 0
			countF := 0

			mask := uint64(1) << uint(i)
			for p := range possibleAnswers {
				if p&mask == 0 {
					countF++
				} else {
					countT++
				}
			}

			if countT > countF {
				output = "T" + output
			} else {
				output = "F" + output
			}
		}

		fmt.Println(answers)
		fmt.Println(possibleAnswers)
		fmt.Printf("Case #%d: %s %d/%d\n", t+1, output, 0, len(possibleAnswers))

	}

}

type Answer struct {
	answer uint64
	score  int
}

func setBit(n uint64, pos uint) uint64 {
	n |= (1 << pos)
	return n
}
