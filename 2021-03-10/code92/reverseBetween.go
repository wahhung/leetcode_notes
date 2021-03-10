// 92. 反转链表 II
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解题思路：先找到需要反转的子链表，然后反转子链表
// Time: O(n)
// Space: O(1)
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 1. 可能头部变化，增加一个root
	root := &ListNode{Next: head}
	prev := root
	// 2.找到 left, right 结点
	i := 1
	for ; i < left; i++ {
		prev = head
		head = head.Next
	}
	start := head
	for ; i < right; i++ {
		head = head.Next
	}
	end := head
	// 3.截取子链表，反转子链表
	next := end.Next
	end.Next = nil
	reverse(start)
	// 4. 将子链表拼回来
	prev.Next, start.Next = end, next
	return root.Next
}

func reverse(head *ListNode) {
	var prev, next *ListNode
	for head != nil {
		next, head.Next = head.Next, prev
		prev, head = head, next
	}
}
