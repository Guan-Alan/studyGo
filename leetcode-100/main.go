package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameNode(p, q)
}

func isSameNode(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	isSameLeft := isSameNode(p.Left, q.Left)
	isSameRight := isSameNode(p.Right, q.Right)
	return isSameLeft && isSameRight
}
