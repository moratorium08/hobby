package main

import "fmt"

type Heap struct {
	Size int
	Buffer []int
}

func (h *Heap) Dump() {
	for _, item := range h.Buffer[1:]{
		fmt.Printf("%d ", item)
	}
	fmt.Println()
}

func (h *Heap) Insert(val int) {
	h.Size += 1

	if h.Size >= len(h.Buffer) {
		h.Buffer = append(h.Buffer, 0)
	}
	h.Buffer[h.Size] = val

	idx := h.Size
	for idx > 0 {
		par := idx / 2
		if h.Buffer[par] > h.Buffer[idx] {
			tmp := h.Buffer[par]
			h.Buffer[par] = h.Buffer[idx]
			h.Buffer[idx] = tmp
		}
		idx = par
	}
}

func (h *Heap) Pop() int {
	ret := h.Buffer[1]

	tail := h.Buffer[h.Size]

	h.Buffer[1] = tail
	idx := 1

	for 2 * idx < h.Size {
		c1, c2 := 2 * idx, 2 * idx + 1
		if c2 >= h.Size {
			c2 = c1
		}
		if h.Buffer[c1] < h.Buffer[c2] {
			if h.Buffer[c1] < h.Buffer[idx] {
				tmp := h.Buffer[c1]
				h.Buffer[c1] = h.Buffer[idx]
				h.Buffer[idx] = tmp
				idx = c1
			} else {
				break
			}
		} else {
			if h.Buffer[c2] < h.Buffer[idx] {
				tmp := h.Buffer[c2]
				h.Buffer[c2] = h.Buffer[idx]
				h.Buffer[idx] = tmp
				idx = c2
			} else {
				break
			}
		}
	}

	h.Size -= 1
	return ret
}

func (h Heap) Top() int {
	return h.Buffer[1]
}

func CreateHeap() *Heap {
	return &Heap{0, []int{0}}
}

func main() {
	heap := CreateHeap()

	heap.Insert(100)
	heap.Insert(80)
	heap.Insert(1)
	heap.Insert(10)

	heap.Dump()

	for i := 0; i < 2; i++ {
		fmt.Printf("%d\n", heap.Pop())
	}

	heap.Dump()

	heap.Insert(99)
	for i := 0; i < 3; i++ {
		fmt.Printf("%d\n", heap.Pop())
	}
	heap.Insert(101)
	heap.Insert(101)
	heap.Insert(101)
	heap.Insert(101)
	heap.Insert(101)
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", heap.Pop())
	}
}
