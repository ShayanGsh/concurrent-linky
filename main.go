package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	coll "github.com/ShayanGsh/concurrent-linky/coll"
)

const (
	NumWorkers = 100
	NumValues  = 100000
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
	startTimestamp := time.Now();
	// numbers := []int{5, 3, 1, 4, 9, 7, 10, 8, 2, 6}
	// Initialize a large pool of random numbers
	values := make([]coll.Comparable, NumValues)
	for i := range values {
		// values[i] = linkedlist.ComparableString(fmt.Sprintf("%d", i))
		values[i] = coll.ComparableInt(
			rand.Intn(NumValues * 1000000),
		)
		// values[i] = coll.ComparableInt(
		// 	numbers[i],
		// )
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

	endTimestamp := time.Now();
	// Close the values channel
	close(valueChannel)
	fmt.Println("Time taken: ", endTimestamp.Sub(startTimestamp))
	// Print the linked list
	// for node := ll.Head; node != nil; node = node.Next {
	// 	fmt.Println(node.Val)
	// }

	// Validate the linked list
	fmt.Println("Is sorted: ", sortValidation(ll))
}

func sortValidation(ll *coll.LinkedList) bool {
	wrongSort := []coll.Comparable{}
	for node := ll.Head; node.Next != nil; node = node.Next {
		if node.Val.CompareTo(node.Next.Val) > 0 {
			wronglySorted := []coll.Comparable{node.Val, node.Next.Val}
			wrongSort = append(wrongSort, wronglySorted...)
		}
	}
	if len(wrongSort) > 0 {
		fmt.Println("Wrongly sorted: ", wrongSort)
		return false
	}
	return true
}
