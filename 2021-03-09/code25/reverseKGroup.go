// 25. K 个一组翻转链表
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n) 遍历一次链表
// Space: O(1)
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 反转部分链表，建立root结点，用于记录开始位置
	root := &ListNode{Val: -1, Next: head}
	tail := root
	start, end := head, head
	for end != nil {
		// 1. 先判断是剩余部分使用有k个结点，如果没有，循环结束
		for i := 0; i < k-1 && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		// 2. 分出start到end一段子链表，用于反转整个子链表
		next := end.Next
		end.Next = nil
		reverse(start)
		// 3. 将子链表的放回原来位置
		tail.Next = end
		start.Next = next
		tail = start
		start = next
		end = next
	}

	return root.Next
}

// Time: O(n)
// Space: O(1)
func reverse(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		// 将当前结点的下一个结点保存起来, 并将Next指针指向 prev
		next, head.Next = head.Next, prev
		// 指向下移
		prev, head = head, next
	}
	return prev
}
