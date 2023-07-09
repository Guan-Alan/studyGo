package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tnode1 := &TreeNode{Val: 1}
	tnode2 := &TreeNode{Val: 2}
	tnode3 := &TreeNode{Val: 3}
	tnode4 := &TreeNode{Val: 4}

	tnode3.Right = tnode4
	tnode3.Left = tnode1
	tnode4.Left = tnode2
	recoverTree(tnode3)
	fmt.Println("========")
	tree(tnode3)
}

// recoverTree 中序遍历, 正常结果应该为升序
// 如果发生前一个节点的值大于了当前节点:
// 1.如果node1为nil，说明第一次发生，找到第一个错误的节点位置就是pre，node1=pre
// 但是有可能是相邻的两个节点被交换，所以node2保存为当前节点位置
// 2.如果node1不为nil，说明异常情况第二次发生，此时找到第二个错误节点位置，就是node本身，node2=node（
// 这种情况说明是间隔的节点发生了交换）
// 找到这两个节点后，交换这两个节点val，就完成了bst树的恢复
// 如果这两个节点都是nil，说明没有发生异常节点被交换，直接返回原bst树
func recoverTree(root *TreeNode) {
	var pre *TreeNode
	var node1 *TreeNode
	var node2 *TreeNode
	var findTree func(node *TreeNode)

	findTree = func(node *TreeNode) {
		if node == nil {
			return
		}

		findTree(node.Left)
		if pre != nil && pre.Val > node.Val {
			if node1 == nil {
				node1 = pre
				node2 = node
			} else {
				node2 = node
				return
			}
		}
		pre = node
		findTree(node.Right)
	}

	findTree(root)
	if node1 == nil && node2 == nil {
		return
	}

	node1.Val, node2.Val = node2.Val, node1.Val
}

func tree(node *TreeNode) {
	if node == nil {
		return
	}
	tree(node.Left)
	fmt.Println(node.Val)
	tree(node.Right)
}
