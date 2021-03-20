// linkedlist.go implements a concurrency safe doubly linked list in go.

package main

import (
	"fmt"
	"sync"
)

// Node is a single node
type Node struct {
	data string
	next *Node
	prev *Node
}

// LinkedList is a list of Nodes
type LinkedList struct {
	head *Node
	tail *Node
	size int
	lock sync.RWMutex
}

// append adds an Item to the end of the linked list
func (ll *LinkedList) append(node *Node) {
	ll.lock.Lock() // prevent this linked list from being modified by another process

	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {
		node.prev = ll.tail
		ll.tail.next = node
		ll.tail = node
	}

	ll.size++        // increment size of list
	ll.lock.Unlock() // unlock linked list before returning
}

// insert inserts a node at a given position i
func (ll *LinkedList) insert(i int, node *Node) {
	ll.lock.Lock() // prevent this list from being modified by another process

	currentNode := ll.head
	if i == 0 {
		currentNode.prev = node
		ll.head = node
		ll.head.next = currentNode
	} else {
		for j := 0; j < i; j++ {
			currentNode = currentNode.next
		}
		node.prev = currentNode.prev
		node.next = currentNode
		currentNode.prev = node
	}

	ll.size++        // increment size of list
	ll.lock.Unlock() // unlock linked list before returning
}

// removeAt removes a node from the linked list at position i
func (ll *LinkedList) removeAt(i int) {
	ll.lock.Lock() // prevent this list from being modified by another process

	currentNode := ll.head
	if i == 0 {
		ll.head = ll.head.next
	} else {
		for j := 0; j < i; j++ {
			currentNode = currentNode.next
		}
		currentNode.prev.next = currentNode.next
		currentNode.next.prev = currentNode.prev
	}

	ll.size--
	ll.lock.Unlock() // unlock linked list before returning
}

// indexOf searches for a data value, if found returns its index, otherwise returns -1
func (ll *LinkedList) indexOf(search string) int {
	ll.lock.Lock() // prevent this list from being modified by another process

	if ll.head == nil {
		ll.lock.Unlock()
		return -1
	}
	currentNode := ll.head
	index := 0

check:
	if currentNode.data == search {
		ll.lock.Unlock()
		return index
	} else if currentNode.next != nil {
		currentNode = currentNode.next
		index++
		goto check
	}

	ll.lock.Unlock() // unlock linked list before returning
	return -1        // if data not in list
}

// isEmpty returns true if linkedlist has a head node
func (ll *LinkedList) isEmpty() bool {
	if ll.head == nil {
		return true
	}
	return false
}

func main() {
	fmt.Println("Concurrency Safe Doubly Linked List in Go")

	linkedList := LinkedList{}
	dataList := []string{"A", "B", "C", "D", "E", "F", "G"}

	// Add all strings from dataList to each node in linkedlist
	for i := range dataList {
		linkedList.append(&Node{data: dataList[i], next: nil, prev: nil})
	}

	// Validate all nodes have been inserted and size updated correctly
	fmt.Println(linkedList.size)                               // => 7
	fmt.Printf("indexOf('A') = %d\n", linkedList.indexOf("A")) // => 0
	fmt.Printf("indexOf('C') = %d\n", linkedList.indexOf("C")) // => 2
	fmt.Printf("indexOf('G') = %d\n", linkedList.indexOf("G")) // => 6

	// Insertion at a given position
	linkedList.insert(0, &Node{data: "Z", next: nil, prev: nil})
	fmt.Println(linkedList.size)                               // => 8
	fmt.Printf("indexOf('Z') = %d\n", linkedList.indexOf("Z")) // => 0
	fmt.Printf("indexOf('A') = %d\n", linkedList.indexOf("A")) // => 1

	// Remove node at a given position
	linkedList.removeAt(4)
	fmt.Println(linkedList.size)                            // => 7
	fmt.Printf("indexOf('D) = %d", linkedList.indexOf("D")) // => -1
}
