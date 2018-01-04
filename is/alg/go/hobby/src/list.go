package main

import (
	"fmt"
	"bufio"
	"os"
)


func getLine() string {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}


type IntList struct {
	val int
	next *IntList
}

func (l *IntList) Insert(x int, idx int) *IntList {
	if idx == 0 {
		c := IntList{x, l}
		return &c
	}
	item := l
	for i := 0; i < idx - 1; i++ {
		item = item.next
	}
	c := IntList{x, item.next}
	item.next = &c
	return l
}

func (l *IntList) Find(idx int) int {
	if idx == 0 {
		return l.val
	}
	return l.next.Find(idx - 1)
}

func (l *IntList) Delete(idx int) {
	if idx == 0 {
		l = l.next
		return
	}

	item := l

	for i := 0; i < idx - 1; i++ {
		item = item.next
	}
	item.next = item.next.next
}

func (l *IntList) Size() int {
	item := l
	cnt := 0
	for item != nil {
		item = item.next
		cnt += 1
	}
	return cnt
}

func (l *IntList) PushBack(val int) *IntList {
	newItem := IntList{val, nil}
	item := l
	if item == nil {
		return &newItem
	}
	for item.next != nil {
		item = item.next
	}
	item.next = &newItem
	return l
}

func (l *IntList) PushFront(val int) *IntList {
	item := l
	l = &IntList{val, item}
	return l
}

func CreateIntList() *IntList {
	return nil
}

func main() {
	list := CreateIntList()
	list = list.PushBack(0)
	list = list.PushBack(1)
	list = list.PushBack(2)
	list = list.PushBack(3)
	list = list.PushBack(4)

	for iter := list; iter != nil; iter = iter.next {
		fmt.Printf("%d\n", iter.val)
	}

	fmt.Println("Delete 2")
	list.Delete(2)
	for iter := list; iter != nil; iter = iter.next {
		fmt.Printf("%d\n", iter.val)
	}
	fmt.Printf("Size: %d\n", list.Size())

	fmt.Println("PushFront -1")
	list = list.PushFront(-1)
	for iter := list; iter != nil; iter = iter.next {
		fmt.Printf("%d\n", iter.val)
	}
	fmt.Printf("Size: %d\n", list.Size())

	fmt.Printf("value at 2 is %d\n", list.Find(2))
	fmt.Println("Insert value 100 at 3")
	list = list.Insert(100, 3)
	for iter := list; iter != nil; iter = iter.next {
		fmt.Printf("%d\n", iter.val)
	}
	fmt.Printf("Size: %d\n", list.Size())
}

