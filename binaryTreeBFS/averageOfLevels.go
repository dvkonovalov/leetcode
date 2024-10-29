package binaryTreeBFS

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

func averageOfLevels(root *TreeNode) []float64 {
	curLevel := []*TreeNode{root}
	avarageList := []float64{}

	for len(curLevel) > 0 {
		nextLevel := []*TreeNode{}
		curSum := 0

		for _, node := range curLevel {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}

			curSum += node.Val
		}
		avarageList = append(avarageList, float64(curSum)/float64(len(curLevel)))
		curLevel = nextLevel
	}

	return avarageList
}
