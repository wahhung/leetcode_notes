//83. 删除排序链表中的重复元素
package main

type ListNode struct {
	Val  int
	Next *ListNode
}


// Time: O(n)
// Space: O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	prev := head
	for prev != nil && prev.Next != nil {
		if prev.Val == prev.Next.Val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return head
}
