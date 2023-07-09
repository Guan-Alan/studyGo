package main

import "fmt"

func main() {
	node := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 3}
	node.Left = node1
	node.Right = node2

	ints := _inorderTraversalV2(node)
	for _, item := range ints {
		fmt.Println(item)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return Traversal(root)
}

// Traversal 关耀文写的
func Traversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	res = append(res, Traversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, Traversal(root.Right)...)
	return res
}

// V2

func _inorderTraversalV2(root *TreeNode) []int {
	var res []int
	var f func(*TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}

		f(node.Left)
		res = append(res, node.Val)
		f(node.Right)

	}
	f(root)
	return res
}

/*
*

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/binary-tree-inorder-traversal/solution/er-cha-shu-de-zhong-xu-bian-li-by-leetcode-solutio/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func _inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}
