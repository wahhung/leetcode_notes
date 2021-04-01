package main

import "fmt"

// 剑指 Offer 10- II. 青蛙跳台阶问题
// 解题思路：
// 		1：n=0->1 n=1->1 n=2->2 n=3->3 n=4->5
//      2: 有上面可以推出，fn(n) = fn(n-1) + fn(n-2)
func numWays(n int) int {
	cur, prev1, pre2 := 1, 1, 1
	for i := 2; i <= n; i++ {
		cur = (prev1 + pre2) % 1000000007
		prev1, pre2 = cur, prev1
	}
	return cur
}

func main()  {
	fmt.Println(numWays(92))
}
