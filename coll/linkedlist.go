package coll

import (
	"sync"
)

// TODO: add worker pool for inserting
// TODO: add removing/poping/peeking
// TODO: implement traversing as a separate function

type Node struct {
	Val  Comparable
	Next *Node
	Prev *Node
	Lock *sync.Mutex
}

func NewNode(val Comparable) *Node {
	return &Node{
		Val:  val,
		Lock: &sync.Mutex{},
	}
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func NewLinkedList() *LinkedList {
	ll := &LinkedList{
		Head: NewNode(nil),
		Tail: NewNode(nil),
	}

	ll.Head.Next = ll.Tail
	ll.Tail.Prev = ll.Head

	return ll
}

func (l *LinkedList) Insert(val Comparable) {
	newNode := NewNode(val)

	selected := l.Head
	nextSelected := selected.Next
	// TODO: add better insertion logic, it should lock prev and next nodes
	for {
		Lock(selected)
		Lock(nextSelected)

		if l.Head.Val == nil && l.Tail.Val == nil {
			l.Head = newNode
			l.Head.Next = l.Tail
			l.Head.Prev = nil
			l.Tail = newNode
			l.Tail.Prev = l.Head
			l.Tail.Next = nil
			Unlock(selected)
			Unlock(nextSelected)
			break
		}

		if (l.Head.Val.CompareTo(val) > 0) {
			newNode.Next = l.Head
			l.Head = newNode
			Unlock(selected)
			Unlock(nextSelected)
			break
		}

		if (l.Tail.Val.CompareTo(val) < 0) {
			newNode.Prev = l.Tail
			l.Tail = newNode
			Unlock(selected)
			Unlock(nextSelected)
			break
		}

		if (nextSelected == nil) {
			selected.Next = newNode
			newNode.Prev = selected
			newNode.Next = nil
			l.Tail = newNode
			Unlock(selected)
			Unlock(nextSelected)
			break
		}

		if selected.Val != nil && selected.Val.CompareTo(val) == 0 {
			Unlock(selected)
			Unlock(nextSelected)
			break
		}

		if nextSelected.Val.CompareTo(val) > 0 {
			selected.Next = newNode
			newNode.Prev = selected
			newNode.Next = nextSelected
			nextSelected.Prev = newNode
			Unlock(selected)
			Unlock(nextSelected)
			break
		}

		Unlock(selected)
		selected = nextSelected
		Unlock(nextSelected)
		nextSelected = nextSelected.Next
	}
}

func Unlock(node *Node) {
	if (node != nil ) {
		node.Lock.Unlock()
	}
}

func Lock (node *Node) {
	if (node != nil) {
		node.Lock.Lock()
	}
}
