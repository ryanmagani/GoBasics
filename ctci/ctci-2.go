package main

import(
	"fmt"
)

type Node struct {
	value int
	next * Node
}

func main() {
	fmt.Println("\nStart CTCI-2")
	removeDuplicates(makeBasicList())
	removeDuplicates(makeDupList())
	removeDuplicatesStructEquality(makeDupList())
	findKthElement(makeBasicList(), 0)
	findKthElement(makeBasicList(), 1)
	findKthElement(makeBasicList(), 2)
	findKthElement(makeBasicList(), 3)
	fmt.Println("End CTCI-2\n")
}

// helpers
func printList(n * Node) {
	cur := n

	for cur.next != nil {
		fmt.Print(cur.value)
		fmt.Print(" -> ")
		cur = cur.next
	}
	fmt.Println(cur.value)
}

func makeBasicList() * Node {
	head := &Node{value: 0, next: nil}
	head.next = &Node{value: 1, next: nil}
	head.next.next = &Node{value: 2, next: nil}
	return head
}

func makeDupList() * Node {
	head := makeBasicList()
	head.next.next.next = &Node{value: 0, next: nil}
	return head
}

// 2.1
func removeDuplicates(head * Node) {
	fmt.Println("\nStart removeDuplicates with input:")
	printList(head)

	cur := head

	// see 1.1 for why we use a map
	seen := make(map[int]byte)
	seen[cur.value] = 1

	for cur != nil && cur.next != nil {
		_, exists := seen[cur.next.value]

		if exists {
			cur.next = cur.next.next
		} else {
			seen[cur.next.value] = 1
		}

		cur = cur.next
	}

	printList(head)
	fmt.Println("End removeDuplicates\n")
}

// 2.1 continued
// from the GoLang documentation:

/*
 * Struct values are comparable if all their fields are comparable.
 * Two struct values are equal if their corresponding non-blank fields are equal.
 * https://golang.org/ref/spec#Comparison_operators
 */

// therefore, this function does not work appropriately! There may be
// some way to hack through equality using interfaces, but I currently
// do not know how to do this
func removeDuplicatesStructEquality(head * Node) {
	fmt.Println("\nStart removeDuplicatesStructEquality with input:")
	printList(head)

	cur := head

	// see 1.1 for why we use a map
	seen := make(map[Node]byte)
	seen[*cur] = 1

	for cur != nil && cur.next != nil {
		_, exists := seen[*(cur.next)]

		if exists {
			cur.next = cur.next.next
		} else {
			seen[*(cur.next)] = 1
		}

		cur = cur.next
	}

	printList(head)
	fmt.Println("End removeDuplicatesStructEquality\n")
}

// 2.2
func findKthElement(head * Node, k int) * Node {
	fmt.Println("\nStart findKthElement inputs:")
	fmt.Println("\tk:", k)
	fmt.Print("\tlist:")
	printList(head)
	fast := head
	slow := head
	for i := 0; i < k && fast.next != nil; i++ {
		fast = fast.next
	}

	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}

	fmt.Println("Found:", slow.value)
	fmt.Println("End findKthElement\n")
	return slow
}