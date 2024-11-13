package binaryTreeBFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	curLevel := []*TreeNode{root}
	ans := make([]int, 0)

	for len(curLevel) > 0 {
		nextLevel := make([]*TreeNode, 0)
		for _, node := range curLevel {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		ans = append(ans, curLevel[len(curLevel)-1].Val)
		curLevel = nextLevel

	}
	return ans
}
