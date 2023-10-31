package coll

import (
	"log"
	"strconv"
	"sync"
)

// TODO: add worker pool for inserting
// TODO: add removing/poping/peeking
// TODO: implement traversing as a separate function

type Comparable interface {
    CompareTo(other Comparable) int
}

type ComparableString string
type ComparableInt int

func (cs ComparableString) CompareTo(other Comparable) int {
    switch v := other.(type) {
    case ComparableString:
        if cs < v {
            return -1
        } else if cs > v {
            return 1
        }
        return 0
    case ComparableInt:
        csInt, err := strconv.Atoi(string(cs))
        if err != nil {
            log.Println("Error: Comparing different types")
            return 0
        }
        if csInt < int(v) {
            return -1
        } else if csInt > int(v) {
            return 1
        }
        return 0
    default:
        log.Println("Error: Comparing different types")
        return 0
    }
}

func (ci ComparableInt) CompareTo(other Comparable) int {
    switch v := other.(type) {
    case ComparableInt:
        if ci < v {
            return -1
        } else if ci > v {
            return 1
        }
        return 0
    case ComparableString:
        vInt, err := strconv.Atoi(string(v))
        if err != nil {
            log.Println("Error: Comparing different types")
            return 0
        }
        if int(ci) < vInt {
            return -1
        } else if int(ci) > vInt {
            return 1
        }
        return 0
    default:
        log.Println("Error: Comparing different types")
        return 0
    }
}

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
	return &LinkedList{
		Head: NewNode(nil),
		Tail: NewNode(nil),
	}
}

func (l *LinkedList) Insert(val Comparable) {
	newNode := NewNode(val)
	inserted := false

	selected := l.Head

	for !inserted {
		selected.Lock.Lock()
		if selected.Val == nil {
			l.Head = newNode
			l.Tail = newNode
			inserted = true
			selected.Lock.Unlock()
			return
		}
		if selected.Val.CompareTo(val) == 0 {
			// log.Println("Duplicate value", val)
			selected.Lock.Unlock()
			return
		}
		if selected.Val.CompareTo(val) == -1 {
			if selected.Next == nil {
				selected.Next = newNode
				newNode.Prev = selected
				l.Tail = newNode
				inserted = true
				selected.Lock.Unlock()
			} else {
				selected.Lock.Unlock()
				selected = selected.Next
			}
		} else {
			if selected.Prev == nil {
				selected.Prev = newNode
				newNode.Next = selected
				l.Head = newNode
				inserted = true
				selected.Lock.Unlock()
			} else {
				selected.Prev.Next = newNode
				newNode.Prev = selected.Prev
				newNode.Next = selected
				selected.Prev = newNode
				inserted = true
				selected.Lock.Unlock()
			}
		}
	}
}
