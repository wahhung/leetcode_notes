package main

import "fmt"

// 33. 搜索旋转排序数组
// 解题思路：
//       0：二分查找
//       1：取数组中间元素 nums[mid]，如果等于target，则返回，否则将数组分成两段nums[0~mid] nums[mid+1, n]
//		 2：其中肯定有一点是有序的，假如前半段是有序的，则须满足 nums[0] <= mid[mid]（否则后半段有序）
//  	 3：判断 target 是否在有序段中，nums[0] <= target < nums[mid-1]，如果不满足，则target在后半段中
//		 4: 在存在段中递归重复 1～3。直到结束，没有找到返回-1
// Time:O(logn) Space:O(1)
func search(nums []int, target int) int {
	if len(nums) < 2 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		fmt.Println(i, j, mid)
		if nums[mid] == target {
			return mid
		}
		if nums[i] <= nums[mid] {
			if nums[i] <= target && target < nums[mid] {
				j = mid
			} else {
				i = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[j] {
				i = mid + 1
			} else {
				j = mid
			}
		}
	}
	return -1
}

func main() {
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	//fmt.Println(search(nums, 0))
	fmt.Println(search([]int{1,3}, 3))
}
