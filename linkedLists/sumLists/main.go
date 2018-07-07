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

func sum(mm, nn linkedList) linkedList {
	if mm.head == nil {
		return nn
	}

	if nn.head == nil {
		return mm
	}

	m := mm.head
	n := nn.head

	sumHead := &node{
		data: (m.data + n.data) % 10,
		next: nil,
	}

	carry := (m.data + n.data) / 10

	sumTail := sumHead

	for m.next != nil || n.next != nil || carry > 0 {
		t := carry

		if m.next != nil {
			t += m.next.data
			m = m.next
		}

		if n.next != nil {
			t += n.next.data
			n = n.next
		}

		newTail := node{
			data: t % 10,
			next: nil,
		}

		carry = t / 10

		sumTail.next = &newTail

		sumTail = &newTail
	}

	return linkedList{head: sumHead}
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
