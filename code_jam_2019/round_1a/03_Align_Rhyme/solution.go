package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T int

	if scanner.Scan() {
		input := scanner.Text()
		T = getInt(input)
	}

	for t := 0; t < T; t++ {
		var N int
		if scanner.Scan() {
			input := scanner.Text()
			N = getInt(input)
		}

		root := &Node{
			parent:   nil,
			children: nil,
			suffix:   "",
			count:    0,
		}

		for i := 0; i < N; i++ {
			if scanner.Scan() {
				s := scanner.Text()
				Add(root, s)
			}
		}

		Print(root)
		fmt.Printf("Case #%d: %d\n", t+1, 2*GetPairCount(root))

	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}

type Node struct {
	parent   *Node
	children []*Node
	suffix   string
	count    int
}

func Add(node *Node, s string) {
	node.count++
	if node.suffix == s {
		return
	}
	for _, n := range node.children {
		if strings.HasSuffix(s, n.suffix) {
			Add(n, s)
			return
		}
	}
	newNode := &Node{
		parent:   node,
		children: nil,
		suffix:   s[len(s)-len(node.suffix)-1:],
		count:    0,
	}
	Add(newNode, s)
	node.children = append(node.children, newNode)
}

func GetPairCount(node *Node) int {
	if node.count < 2 {
		return 0
	}
	c := 0
	total := node.count
	for _, n := range node.children {
		p := GetPairCount(n)
		if p != 0 {
			c += p
			total -= 2
		}
	}

	if total >= 2 && node.suffix != "" {
		return c + 1
	}
	return c
}

func Print(n *Node) {
	fmt.Println(n.suffix, n.count)
	for _, nn := range n.children {
		Print(nn)
	}
}
