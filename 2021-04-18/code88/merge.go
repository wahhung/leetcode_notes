// 88. 合并两个有序数组
// 归并
// 1.因为 nums1 后面都是无用数据，所以遍历从后往前
// 2. 比较 nums1[i] = max(nums1[m-i],nums2[n-i]) 
// 3. 当 m - i == 0 or n -i == 0，循环结束
// 4.且当 n-i > 0 时，需要顺序拷贝数据到nums1
// Time:O(n+m) Space:O(1)
func merge(nums1 []int, m int, nums2 []int, n int)  {
    i := len(nums1) - 1
    for n > 0 && m > 0 {
        if nums1[m-1] > nums2[n-1] {
            nums1[i] = nums1[m-1]
            m--
        } else {
            nums1[i] = nums2[n-1]
            n--
        }
        i--
    }
    for i := 0; i < n; i++ {
        nums1[i] = nums2[i]
    }
}
