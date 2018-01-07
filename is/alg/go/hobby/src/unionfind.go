package main

import "fmt"

type UnionFind struct {
	parents []int
	ranks []int
	size int
}

func (uf *UnionFind) Append() {
	uf.parents = append(uf.parents, uf.size)
	uf.ranks = append(uf.ranks, 0)
	uf.size += 1
}

func (uf *UnionFind) Same(x1 int, x2 int) bool {
	return uf.Find(x1) == uf.Find(x2)
}

func (uf *UnionFind) Find(index int) int {
	if index == uf.parents[index] {
		return index
	} else {
		uf.parents[index] = uf.Find(uf.parents[index])
		return uf.parents[index]
	}
}

func (uf *UnionFind) Unite(x1 int, x2 int) {
	x1, x2 = uf.Find(x1), uf.Find(x2)
	if x1 == x2 {
		return
	}

	if uf.ranks[x1] < uf.ranks[x2] {
		uf.parents[x1] = x2
	} else {
		uf.parents[x2] = x1
		if uf.ranks[x1] == uf.ranks[x2] {
			uf.ranks[x1] += 1
		}
	}
}

func NewUnionFind(size int) *UnionFind {
	uf := UnionFind{make([]int, size), make([]int, size), size}
	for idx := range uf.parents {
		uf.parents[idx] = idx
	}
	return &uf
}

func Check(uf *UnionFind, x int , y int) {
	fmt.Printf("%d, %d -> %v\n", x, y, uf.Same(x, y))
}

func main() {
	uf := NewUnionFind(10)

	// 1 = 4 = 7, 0 = 3 = 6 = 9, 2 = 5 = 8
	uf.Unite(1, 4)
	uf.Unite(4, 7)

	uf.Unite(0, 3)
	uf.Unite(6, 9)
	uf.Unite(9, 0)

	uf.Unite(2, 5)
	uf.Unite(5, 8)

	Check(uf, 1, 7)
	Check(uf, 0, 7)
	Check(uf, 9, 3)
	Check(uf, 2, 5)
}

