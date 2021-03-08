// 141. 环形链表
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法1：龟兔赛跑
// 思路分析： 龟兔赛跑问题，当两者相遇，则证明有换，否则链表遍历到末端
// Time: O(n)
// Space: O(n)
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

// 方法2：哈希法
// 思路分析：如果没有环，则证明一个结点只会出现一次
// Time: O(n)
// Space: O(n)
func hasCycle(head *ListNode) bool {
	nodeData := make(map[*ListNode]struct{}) // 空结构体不占用内存 unsafe.Sizeof(s) == 0
	for head != nil {
		if _, ok := nodeData[head]; ok {
			return true
		}
		nodeData[head] = struct{}{}
		head = head.Next
	}
	return false
}

