package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Scan(&N)

		nodes := make([]Node, N)

		for i := 0; i < N; i++ {
			var fun int
			fmt.Scan(&fun)
			nodes[i].fun = fun
			// nodes[i].index = i + 1
		}

		rootNode := Node{
			fun: 0,
			// index: 0,
		}

		for i := 0; i < N; i++ {
			var parentIndex int
			fmt.Scan(&parentIndex)
			if parentIndex == 0 {
				rootNode.children = append(rootNode.children, &nodes[i])
			} else {
				nodes[parentIndex-1].children = append(nodes[parentIndex-1].children, &nodes[i])
			}
		}

		dft(&rootNode)

		// fmt.Println(nodes)
		// fmt.Println(rootNode)

		fmt.Printf("Case #%d: %d \n", t+1, rootNode.unconnectedFun+rootNode.connectedFun)
	}
}

type Node struct {
	fun int
	// index int
	children       []*Node
	connectedFun   int
	unconnectedFun int
}

func dft(node *Node) (int, int) {
	if len(node.children) == 0 {
		node.connectedFun = node.fun
		node.unconnectedFun = 0
		return node.connectedFun, node.unconnectedFun
	}

	minConnectedFromChild := math.MaxInt64
	totalUnconnected := 0

	for _, child := range node.children {
		connect, unconnect := dft(child)

		// fmt.Println("Processing", connect, unconnect, child.fun)

		if minConnectedFromChild > connect {
			minConnectedFromChild = connect
		}
		totalUnconnected += connect + unconnect
	}
	totalUnconnected -= minConnectedFromChild
	connected := max(minConnectedFromChild, node.fun)

	node.connectedFun = connected
	node.unconnectedFun = totalUnconnected
	return connected, totalUnconnected
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
