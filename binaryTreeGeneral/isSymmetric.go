package binaryTreeGeneral

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isMirror(leftRoot *TreeNode, rightRoot *TreeNode) bool {
	if leftRoot == nil && rightRoot == nil {
		return true
	}
	if leftRoot == nil || rightRoot == nil {
		return false
	}

	if leftRoot.Val != rightRoot.Val {
		return false
	}
	return isMirror(leftRoot.Left, rightRoot.Right) && isMirror(leftRoot.Right, rightRoot.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}
