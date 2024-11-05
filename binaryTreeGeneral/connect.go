package binaryTreeGeneral

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	curLevel := []*Node{root}
	for len(curLevel) > 0 {
		nextLevel := make([]*Node, 0)

		for i := 0; i < len(curLevel); i++ {
			if curLevel[i].Left != nil {
				nextLevel = append(nextLevel, curLevel[i].Left)
			}
			if curLevel[i].Right != nil {
				nextLevel = append(nextLevel, curLevel[i].Right)
			}
			if i < len(curLevel)-1 {
				curLevel[i].Next = curLevel[i+1]
			}
		}

		curLevel = nextLevel

	}

	return root
}
