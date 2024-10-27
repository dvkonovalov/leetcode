/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 2 {
		return &TreeNode{
			Val: nums[0],
			Left: nil,
			Right: sortedArrayToBST(nums[1:]),
		}
	}
	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
			Left: nil,
			Right: nil,
		}
	}

	mid := len(nums) / 2
	tree := TreeNode{
		Val: nums[mid],
		Right: nil,
		Left: nil,
	}

	tree.Left = sortedArrayToBST(nums[:mid])
	tree.Right = sortedArrayToBST(nums[mid+1:])
	return &tree

}