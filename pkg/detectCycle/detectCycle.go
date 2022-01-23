package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func main() {
	l3 := ListNode{3, nil}
	l2 := ListNode{2, nil}
	l0 := ListNode{0, nil}
	l4 := ListNode{4, nil}

	l3.Next = &l2
	l2.Next = &l0
	l0.Next = &l4
	l4.Next = &l2

	fmt.Println(detectCycle(&l3))

}
