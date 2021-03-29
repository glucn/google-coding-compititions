package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var N int
	fmt.Scan(&N)

	rand.Seed(time.Now().UnixNano())
	values := make([]int, N)

	for i := range values {
		values[i] = rand.Intn(100)
	}

	fmt.Println(values)

	s := NewSlice(values)
	sort.Sort(s)
	fmt.Println(s.idx)

	// for {
	// TODO: implement the body
	// }

}

type Slice struct {
	sort.IntSlice
	idx []int
}

func (s Slice) Swap(i, j int) {
	s.IntSlice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func NewSlice(values []int) *Slice {
	s := &Slice{IntSlice: sort.IntSlice(values), idx: make([]int, len(values))}
	for i := range s.idx {
		s.idx[i] = i + 1
	}
	return s
}
