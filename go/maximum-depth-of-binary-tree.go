package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/maximum-depth-of-binary-tree
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := maxDepth(root.Left) + 1
	right := maxDepth(root.Right) + 1

	if left > right {
		return left
	}
	return right
}
