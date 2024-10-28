package binarySearchTree

import "math"

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

var MinDiff, prevValue int

func getMinimumDifference(root *TreeNode) int {
	prevValue = math.MaxInt32
	MinDiff = math.MaxInt32
	findMinimum(root)
	return MinDiff
}

func findMinimum(root *TreeNode) {
	if root == nil {
		return
	}
	findMinimum(root.Left)

	if root.Val > prevValue {
		MinDiff = min(MinDiff, root.Val-prevValue)
	} else {
		MinDiff = min(MinDiff, prevValue-root.Val)
	}
	prevValue = root.Val

	findMinimum(root.Right)
	return
}
