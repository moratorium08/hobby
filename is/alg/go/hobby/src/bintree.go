package main

import (
	"errors"
	"fmt"
)

type Tag interface {
	Compare(Tag) int
}

type Element struct {
	Tag Tag
	Value interface{}
}

type BinTree struct {
	Left *BinTree
	Right *BinTree
	Elem Element
}

func (bt *BinTree) find(t Tag) (Element, error) {
	var val = t.Compare(bt.Elem.Tag)
	if val == 0 {
		return bt.Elem, nil
	}

	if val < 0 {
		if bt.Left == nil {
			return Element{}, errors.New("No such entity")
		}
		return bt.Left.find(t)
	} else {
		if bt.Right== nil {
			return Element{}, errors.New("No such entity")
		}
		return bt.Right.find(t)
	}
}

func (bt *BinTree) Find(t Tag) (interface{}, error) {
	elem, err := bt.find(t)
	if err != nil {
		return nil, err
	}
	return elem.Value, nil
}

func (bt *BinTree) Insert(t Tag, v interface{}) *BinTree{
	if bt == nil {
		elem := Element{t, v}
		return &BinTree{nil, nil, elem}
	}

	var val = t.Compare(bt.Elem.Tag)
	if val == 0 {
		bt.Elem.Value = v
		return bt
	}

	if val < 0 {
		if bt.Left == nil {
			elem := Element{t, v}
			bt.Left = &BinTree{nil, nil, elem}
			return bt
		} else {
			return bt.Left.Insert(t, v)
		}
	} else {
		if bt.Right == nil {
			elem := Element{t, v}
			bt.Right = &BinTree{nil, nil, elem}
			return bt
		} else {
			return bt.Right.Insert(t, v)
		}
	}
}

func (bt *BinTree) Delete(t Tag) *BinTree {
	var val = t.Compare(bt.Elem.Tag)
	if val < 0 {
		bt.Left = bt.Left.Delete(t)
		return bt
	} else if val > 0 {
		bt.Right = bt.Right.Delete(t)
		return bt
	}
	if bt.Left == nil && bt.Right == nil {
		return nil
	}
	if bt.Left == nil {
		return bt.Right
	}
	if bt.Right == nil {
		return bt.Left
	}

	min := bt.DeleteMin()
	min.Left = bt.Left
	min.Right = bt.Right
	return min
}

func (bt *BinTree) DeleteMin() *BinTree {
	if bt.Left.Left == nil {
		ret := bt.Left
		bt.Left = nil
		return ret
	}
	return bt.Left.DeleteMin()
}

type String string
type Int int

func (s String) Compare(x Tag) int {
	arg, _ := x.(String)
	if s < arg {
		return -1
	} else if s > arg {
		return 1
	}
	return 0
}

func (x Int) Compare(arg Tag) int {
	y, _ := arg.(Int)
	if x < y {
		return -1
	} else if x > y {
		return 1
	}
	return 0
}

func CreateBinTree() *BinTree {
	return nil
}

func main() {
	bt := CreateBinTree()

	bt = bt.Insert(String("hoge"), 1)
	bt = bt.Insert(String("fuga"), 2)

	found, _ := bt.Find(String("hoge"))
	fmt.Printf("%d\n", found)
	bt = bt.Delete(String("hoge"))
	_, err := bt.Find(String("hoge"))
	fmt.Println(err)
}
