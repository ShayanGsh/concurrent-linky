package main

import (
	"fmt"
	"math/rand"
	"sync"

	coll "github.com/ShayanGsh/concurrent-linky/coll"
)

const (
	NumWorkers = 10
	NumValues  = 200
)

// Worker function
func worker(wg *sync.WaitGroup, values <-chan coll.Comparable, cancel <-chan struct{}, ll *coll.LinkedList) {
	defer wg.Done()
	for {
		select {
		case val := <-values:
			ll.Insert(val)
		case <-cancel:
			return
		}
	}
}

func main() {
	// Initialize a large pool of random numbers
	values := make([]coll.Comparable, NumValues)
	for i := range values {
		// values[i] = linkedlist.ComparableString(fmt.Sprintf("%d", i))
		values[i] = coll.ComparableInt(
			rand.Intn(NumValues * 1000000),
		)

	}

	// Initialize the linked list
	ll := coll.NewLinkedList()

	// Create a channel to pass values to workers
	valueChannel := make(chan coll.Comparable)
	cancel := make(chan struct{})

	// Spawn 10 workers
	var wg sync.WaitGroup
	for i := 0; i < NumWorkers; i++ {
		wg.Add(1)
		go worker(&wg, valueChannel, cancel, ll)
	}

	// Pass values to workers
	for _, val := range values {
		valueChannel <- val
	}

	// Close the channel after passing all values
	close(cancel)

	// Wait for all workers to finish
	wg.Wait()

	// Close the values channel
	close(valueChannel)

	// Print the linked list
	for node := ll.Head; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}
