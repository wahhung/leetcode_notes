// 102. 二叉树的层序遍历
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 解题思路：BFS，遍历父结点，现将子结点存起来，遍历完父结点，子节点变成父结点继续循环，直到没有子结点
// Time: O(n) 所有结点走一遍
// Space: O(n)
func levelOrder(root *TreeNode) [][]int {
	vals := [][]int{}
	if root == nil {
		return vals
	}
	var idx int
	parent := []*TreeNode{root}
	for len(parent) > 0 {
		child := []*TreeNode{}
		vals = append(vals, []int{})
		for i := 0; i < len(parent); i++ {
			vals[idx] = append(vals[idx], parent[i].Val)
			if parent[i].Left != nil {
				child = append(child, parent[i].Left)
			}
			if parent[i].Right != nil {
				child = append(child, parent[i].Right)
			}
		}
		parent = child
		idx++
	}
	return vals
}