// 122. 买卖股票的最佳时机 II
package main

import "fmt"

// 贪心算法：每一步都做出当时看起来最佳的选择(局部最优)
// 解题思路：只要今天比昨天高，就可以交易
// Time:O(n)
// Space:O(1)
func maxProfit1(prices []int) int {
	var profit int
	for i := 1; i < len(prices); i++ {
		profit += max(0, prices[i]-prices[i-1])
	}
	return profit
}

// 动态规划
// 状态定义：dp0/dp1:第i天手上有/没有股票的最大利润
// 转移方程：
//		dp0 = max(dp0, dp1 + prices[i]) 前一天没买和前一天买了今天卖出，做比较
//      dp1 = max(dp1, dp0 - prices[i]) 前一天卖了和前一天没卖今天买了，做比较
// Time: O(n)
// Space:O(1)
func maxProfit(prices []int) int {
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		today0 := max(dp0, dp1+prices[i])
		today1 := max(dp1, dp0-prices[i])
		dp0, dp1 = today0, today1
	}
	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(nums))
}
