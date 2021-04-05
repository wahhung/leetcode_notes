 // 142. 环形链表 II
 // 解题方法：快慢指针(龟兔赛跑)
 //     1: 利用快慢指针找到是否有环，无环则 return nil
 //     2: 从head走和slow指针开始，每次走一步，相遇则为入环点
 // 具体公式：
 //     1：假设整个链表长度为s = a(环外) + b(环内), c 为环内的第c个点相遇
 //     2: 由于快指针是慢指针的两倍，且快指针比慢指针多走n环
 //         即 a + (n+m)b + c = 2(a + mb + c)  => a = kb-c, k=abs(m-n)
 //     3: head到入环点走了a步，slow也走了a步，则 slow 走的为 
 //         a+mb+c + a = a+mb+c + kb-c = a + (m+k)b 即 slow 指针刚好走完环回到入环点
 // Time:O(n) Space:O(1)
func detectCycle(head *ListNode) *ListNode {
    fast, slow := head, head
    var hasCycle bool
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if fast == slow {
            hasCycle = true
            break
        }
    }
    if !hasCycle {
        return nil
    }
    for head != nil {
        if head == slow {
            return head
        }
        head = head.Next
        slow = slow.Next
    }
    return nil
}
