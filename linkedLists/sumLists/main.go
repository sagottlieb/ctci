package main

import (
	"fmt"

	"github.com/sagottlieb/ctci/linkedLists/sumLists/datastructs"
)

func main() {
	fmt.Println("Run go test and go test -bench=.")
}

// 1's digit (least significant) is at the head
func sum(mm, nn datastructs.LinkedList) datastructs.LinkedList {
	if mm.Head == nil {
		return nn
	}

	if nn.Head == nil {
		return mm
	}

	m := mm.Head
	n := nn.Head

	sumHead := &datastructs.Node{
		Data: (m.Data + n.Data) % 10,
		Next: nil,
	}

	carry := (m.Data + n.Data) / 10

	sumTail := sumHead

	for m.Next != nil || n.Next != nil || carry > 0 {
		t := carry

		if m.Next != nil {
			t += m.Next.Data
			m = m.Next
		}

		if n.Next != nil {
			t += n.Next.Data
			n = n.Next
		}

		newTail := datastructs.Node{
			Data: t % 10,
			Next: nil,
		}

		carry = t / 10

		sumTail.Next = &newTail

		sumTail = &newTail
	}

	return datastructs.LinkedList{Head: sumHead}
}

// most significant digit is at the head
// convert LLs to stacks then take sum
func sumReverse(mm, nn datastructs.LinkedList) datastructs.LinkedList {
	mmStack := llToStack(mm)
	nnStack := llToStack(nn)

	var sumHead *datastructs.Node
	carry := 0
	for {
		if mmStack.IsEmpty() && nnStack.IsEmpty() && carry == 0 {
			break
		}
		mval := 0
		if !mmStack.IsEmpty() {
			mval = mmStack.Pop()
		}
		nval := 0
		if !nnStack.IsEmpty() {
			nval = nnStack.Pop()
		}
		newSumHead := &datastructs.Node{
			Data: (mval + nval + carry) % 10,
			Next: sumHead,
		}
		carry = (mval + nval + carry) / 10
		sumHead = newSumHead
	}
	return datastructs.LinkedList{Head: sumHead}
}

func llToStack(mm datastructs.LinkedList) datastructs.Stack {
	out := datastructs.Stack{
		Top: nil,
	}

	m := mm.Head
	for m != nil {
		out.Push(m.Data)
		m = m.Next
	}

	return out
}

func formatLL(mm datastructs.LinkedList) string {
	return formatLLNode(mm.Head)
}

func formatLLNode(n *datastructs.Node) string {
	if n == nil {
		return ""
	}
	out := fmt.Sprintf("%d->", n.Data)
	for n.Next != nil {
		out += fmt.Sprintf("%d->", n.Next.Data)
		n = n.Next
	}
	out += "nil"
	return out
}
