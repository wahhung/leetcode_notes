// 189. 旋转数组
package main

// 方法三：数组翻转
// 解题思路：
//		右移动k次，即将 k%n 个元素放到头部，其他顺位往下
//	 	先旋转整个数组，确保 k%n 个元素在头部，再分布旋转[0,k%n-1]和[k%n,-1]旋转回来
// Time: O(n)
// Space: O(1)
func rotate3(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

// 方法二：环形替换
// 解题思路：
//		tmp = nums[0], tmp 目标存放位置是 idx=(0+k)%n
// 		tmp替换idx位置的值并保存，然后找新的temp存放位置，重新回到开始位置为一圈
//		共右 gcb(k, n) 圈
// Time: O(n) 所有位置都遍历一遍
// Space: O(1)
func rotate2(nums []int, k int) {
	n := len(nums)
	k %= n
	count := gcb(k, n) // count 个环
	for i := 0; i < count; i++ {
		tmp, idx := nums[i], i
		for {
			next := (idx + k) % n
			nums[next], tmp, idx = tmp, nums[next], next
			if i == idx {
				break
			}
		}
	}
}

// 最大公约数
func gcb(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// 方法一：使用额外数组，通过 (i + k) % n 获取最终存放位置
// Time: O(n)
// Space: O(n)
func rotate1(nums []int, k int) {
	n := len(nums)
	tmp := append([]int{}, nums...)
	for i := 0; i < n; i++ {
		idx := (i + k) % n
		nums[idx] = tmp[i]
	}
}
