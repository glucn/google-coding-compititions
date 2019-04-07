package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TODO: Test Set 2 error

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T int

	if scanner.Scan() {
		input := scanner.Text()
		T = getInt(input)
	}

	for t := 0; t < T; t++ {
		var N int
		var B int
		var F int
		if scanner.Scan() {
			input := scanner.Text()
			s := strings.Split(input, " ")
			N = getInt(s[0])
			B = getInt(s[1])
			F = getInt(s[2])
		}

		root := &node{
			start:  0,
			end:    N - 1,
			broken: B,
		}

		node := root

		_ = F
		for f := 0; f < F; f++ {
			fmt.Println(node.makeGuess())
			var response string
			if scanner.Scan() {
				response = scanner.Text()
				if response == "-1" {
					// wrong answer or something, terminate
					return
				}
			}
			node.checkResponse(response)

			result, err := node.getFinalResult()
			if err != nil {
				continue
			} else {
				fmt.Println(getOutputString(result))
				break
			}

		}

		if scanner.Scan() {
			s := scanner.Text()
			if s == "1" {
				// correct answer
				continue
			}
			// wrong answer
			return
		}

		//fmt.Printf("Case #%d: %s\n", t+1, "")

	}
}

func getOutputString(input []int) string {
	var s string
	for i := range input {
		s = s + strconv.Itoa(input[i])
		if i != len(input)-1 {
			s = s + " "
		}
	}
	return s
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}

type node struct {
	start  int
	end    int
	broken int
	left   *node
	right  *node
}

func (n *node) makeGuess() string {
	if n.start == n.end {
		return "1"
	}
	if n.left == nil && n.right == nil {
		mid := (n.start + n.end) >> 1

		n.left = &node{
			start:  n.start,
			end:    mid,
			broken: -1,
		}
		n.right = &node{
			start:  mid + 1,
			end:    n.end,
			broken: -1,
		}

		return getGuessString(n.end - n.start + 1)
	}
	return n.left.makeGuess() + n.right.makeGuess()
}

func (n *node) checkResponse(response string) {
	//fmt.Println("checking response for node, ", *n, " response ", response)
	if n.broken == n.end-n.start+1 || n.broken == 0 {
		return
	}
	if n.broken != -1 {
		//if len(response) == 0 {
		//	return
		//}

		if n.left.broken == -1 || n.right.broken == -1 {
			endOfOne := getIndex(response)
			n.left.checkResponse(response[:endOfOne])
			n.right.checkResponse(response[endOfOne:])
			return
		}

		endIndex := n.left.end - n.left.start - n.left.broken + 1
		n.left.checkResponse(response[:endIndex])
		n.right.checkResponse(response[endIndex:])

		return
	}
	n.broken = n.end - n.start + 1 - len(response)
	//fmt.Println("updated node, ", *n)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (n *node) getFinalResult() ([]int, error) {
	if n.broken == n.end-n.start+1 {
		result := make([]int, n.broken)
		for i := range result {
			result[i] = n.start + i
		}
		return result, nil
	}
	if n.broken == 0 {
		return nil, nil
	}
	if n.left == nil || n.right == nil {
		return nil, fmt.Errorf("guess again")
	}
	left, err := n.left.getFinalResult()
	if err != nil {
		return nil, err
	}
	right, err := n.right.getFinalResult()
	if err != nil {
		return nil, err
	}
	return append(left, right...), nil
}

func getIndex(s string) int {
	for i := range s {
		if s[i] == '0' {
			return i
		}
	}
	return len(s)
}

func getGuessString(length int) string {
	res := make([]byte, length)
	for i := 0; i < length; i++ {
		if i < (length+1)/2 {
			res[i] = '1'
		} else {
			res[i] = '0'
		}
	}
	return string(res[:])
}

func getPlainString(length int) string {
	res := make([]byte, length)
	for i := 0; i < length; i++ {
		res[i] = '1'
	}
	return string(res[:])
}
