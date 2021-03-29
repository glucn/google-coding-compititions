package main

import "fmt"

func main() {
	var T, N, Q int
	fmt.Scan(&T)
	fmt.Scan(&N)
	fmt.Scan(&Q)

	for t := 0; t < T; t++ {
		// init
		answer := tree{root: &node{index: 1, left: &node{index: 2}}}
		todo := make([]int, N-2)
		for i := range todo {
			todo[i] = i + 3
		}

		for len(todo) > 0 {
			n := todo[0]
			todo = todo[1:]

			// TODO: rebalance the tree
			answer.Insert(n)
		}

		fmt.Println(answer.String())

		var result int
		fmt.Scan(&result)

		if result != 1 {
			break
		}

	}
}

type tree struct {
	root *node
}

func (t *tree) String() string {
	return t.root.String()
}

func (t *tree) Insert(idx int) {
	t.root.Insert(idx)
}

type node struct {
	index int
	left  *node
	right *node
}

func (n *node) String() string {
	left := ""
	if n.left != nil {
		left = fmt.Sprintf("%s ", n.left.String())
	}
	right := ""
	if n.right != nil {
		right = fmt.Sprintf(" %s", n.right.String())
	}
	return fmt.Sprintf("%s%d%s", left, n.index, right)
}

func (n *node) Insert(idx int) {
	if n.LeftInsert(idx) {
		return
	}

	n.RightInsert(idx)

}

func (n *node) LeftInsert(idx int) bool {
	if n.left == nil {
		n.left = &node{index: idx}
		return true
	}

	mid := ask(idx, n.index, n.left.index)

	switch mid {
	case n.index:
		// the value to be inserted is "righter" than the root, it should go with RightInsert()
		return false
	case n.left.index:
		// the value is "lefter" than n.left, leftInsert it to n.left
		return n.left.LeftInsert(idx)
	case idx:
		// the value is "righter" than n.left, but "lefter" than n, rightInsert to n.left
		return n.left.RightInsert(idx)
	}

	// This should not happen
	return false
}

func (n *node) RightInsert(idx int) bool {
	if n.right == nil {
		n.right = &node{index: idx}
		return true
	}

	mid := ask(idx, n.index, n.right.index)

	switch mid {
	case n.index:
		// the value to be inserted is "lefter" than the root, it should go with LeftInsert()
		return false
	case n.right.index:
		// the value is "righter" than n.right, rightInsert it to n.right
		return n.right.RightInsert(idx)
	case idx:
		// the value is "lefter" than n.right, but "righter" than n, leftInsert to n.right
		return n.right.LeftInsert(idx)
	}
	// This should not happen
	return false
}

func ask(a, b, c int) int {
	fmt.Printf("%d %d %d\n", a, b, c)

	var resp int
	fmt.Scan(&resp)

	return resp
}
