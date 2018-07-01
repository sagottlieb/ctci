package main

import (
	"fmt"
)

func main() {
	fmt.Println("Run go test and go test -bench=.")
}

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func removeDups(mm linkedList) linkedList {
	n := mm.head
	if n == nil {
		return mm
	}

	seen := map[int]bool{n.data: true}

	for n.next != nil {
		if _, exists := seen[n.next.data]; exists {
			// skip n.next
			n.next = n.next.next
		} else {
			seen[n.next.data] = true
			n = n.next
		}
	}

	return mm
}

// for each node `curr` as you traverse the linked list,
// do a "look-ahead" pass through the list checking [and skipping]
// any nodes with value equal to curr.data.
// O(n^2)
func removeDupsNoBuffer(mm linkedList) linkedList {

	curr := mm.head

	if curr == nil {
		return mm
	}

	for curr.next != nil {

		n := curr

		for n.next != nil {
			if curr.data == n.next.data {
				// skip it!
				n.next = n.next.next
			} else {
				n = n.next
			}
		}

		curr = curr.next
	}


	return mm
}

func formatLL(mm linkedList) string {
	return formatNode(mm.head)
}

func formatNode(n *node) string {
	if n == nil {
		return ""
	}
	out := fmt.Sprintf("%d->", n.data)
	for n.next != nil {
		out += fmt.Sprintf("%d->", n.next.data)
		n = n.next
	}
	out += "nil"
	return out
}
