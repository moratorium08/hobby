package main

import (
	"fmt"
	"errors"
)

type AVLTree struct {
	Height int
	Left *AVLTree
	Right *AVLTree
	Elem Element
}

func (at *AVLTree) RotateRight() *AVLTree {
	if at.Left == nil {
		fmt.Errorf("なんかおかしいよ\n")
		return nil
	}

	ret := at.Left
	at.Left = ret.Right
	ret.Right = at
	return ret
}

func (at *AVLTree) RotateLeft() *AVLTree {
	if at.Right == nil {
		fmt.Errorf("なんかおかしいよ\n")
		return nil
	}

	ret := at.Right
	at.Right = ret.Left
	ret.Left = at
	return ret
}

func ShowTabs(tab int) {
	for i := 0; i < tab; i++ {
		fmt.Printf(" ")
	}
}

func (at *AVLTree) Dump(tab int) {
	ShowTabs(tab)
	fmt.Printf("%v:%v <- %d\n", at.Elem.Tag, at.Elem.Value, at.Height)
	if at.Left != nil {
		at.Left.Dump(tab + 2)
	}
	if at.Right != nil {
		at.Right.Dump(tab + 2)
	}

	if tab == 0 {
		fmt.Println()
	}
}

func (at *AVLTree) Insert(tag Tag, val interface{}) (*AVLTree, bool) {
	if at == nil {
		elem := Element{tag, val}
		return &AVLTree{0, nil, nil, elem}, true
	}

	v := tag.Compare(at.Elem.Tag)
	if v == 0 {
		at.Elem.Value = val
		return at, false
	}
	//fmt.Printf("%v %v\n", tag, val)

	if v < 0 {
		left, flag := at.Left.Insert(tag, val)
		at.Left = left
		if !flag {
			return at, false
		}
		if at.Height == -1 {
			at.Height = 0
			return at, false
		}
		if at.Height == 0 {
			at.Height = 1
			return at, true
		}
		if at.Left.Height == 1 {
			at.Height = 0
			at = at.RotateRight()
			at.Height = 0
			return at, false
		}
		if at.Left.Height == -1 {
			at.Left.Height = 0
			at.Left = at.Left.RotateLeft()
			at = at.RotateRight()
			at.Height = 0
			return at, false
		}
		fmt.Errorf("なんかおかしい")
		return at, false
	} else {
		right, flag := at.Right.Insert(tag, val)
		at.Right= right
		if !flag {

			return at, false
		}
		if at.Height == 1 {
			at.Height = 0
			return at, false
		}
		if at.Height == 0 {
			at.Height = -1
			return at, true
		}

		if at.Right.Height == -1 {
			at.Height = 0
			at = at.RotateLeft()
			at.Height = 0
			return at, false
		}
		if at.Left.Height == 1 {
			at.Right.Height = 0
			at.Left = at.Right.RotateRight()
			at = at.RotateLeft()
			at.Height = 0
			return at, false
		}
		fmt.Errorf("なんかおかしい")
		return at, false
	}
}

func (at *AVLTree) Delete(tag Tag) *AVLTree {
	// 飽きたまた今度
	return at
}

func (at *AVLTree) find(t Tag) (Element, error) {
	var val = t.Compare(at.Elem.Tag)
	if val == 0 {
		return at.Elem, nil
	}

	if val < 0 {
		if at.Left == nil {
			return Element{}, errors.New("No such entity")
		}
		return at.Left.find(t)
	} else {
		if at.Right== nil {
			return Element{}, errors.New("No such entity")
		}
		return at.Right.find(t)
	}
}

func (at *AVLTree) Find(t Tag) (interface{}, error) {
	elem, err := at.find(t)
	if err != nil {
		return nil, err
	}
	return elem.Value, nil
}

func CreateAVLTree() *AVLTree {
	return nil
}

func main() {
	at := CreateAVLTree()
	at, _ = at.Insert(String("c"), 1)
	at.Dump(0)
	at, _ = at.Insert(String("b"), 5)
	at.Dump(0)
	at, _ = at.Insert(String("a"), 6)
	at.Dump(0)
	at, _ = at.Insert(String("d"), 7)
	at.Dump(0)
	at, _ = at.Insert(String("e"), 8)
	at.Dump(0)
}
