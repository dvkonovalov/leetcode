package binarySearchTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	arr := treeInArray(root)
	return arr[k-1]
}

func treeInArray(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := append(treeInArray(root.Left), root.Val)
	ans = append(ans, treeInArray(root.Right)...)
	return ans
}
