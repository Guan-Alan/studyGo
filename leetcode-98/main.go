package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 用中序遍历，看是不是升序
func isValidBST(root *TreeNode) bool {
	elem := math.MinInt
	res := true
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil || !res {
			return
		}

		f(node.Left)

		if node.Val <= elem {
			res = false
			return
		}
		elem = node.Val

		f(node.Right)
	}
	f(root)
	return res
}

// 解法1 能通过
func _isValidBST(root *TreeNode) bool {
	return valid(root, math.MinInt, math.MaxInt)
}

func valid(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min {
		return false
	}

	if node.Val >= max {
		return false
	}

	return valid(node.Left, min, node.Val) && valid(node.Right, node.Val, max)
}
