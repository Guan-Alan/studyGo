package main

func main() {
	generateTrees(3)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var remark [][][]*TreeNode

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	remark = make([][][]*TreeNode, n+1, n+1)
	for index, _ := range remark {
		remark[index] = make([][]*TreeNode, n+1, n+1)
	}
	return generateTree(1, n)
}

func generateTree(start, n int) []*TreeNode {
	if start > n {
		return []*TreeNode{nil}
	}

	if len(remark[start][n]) > 0 {
		return remark[start][n]
	}

	var res []*TreeNode
	for i := start; i <= n; i++ {
		leftTress := generateTree(start, i-1)
		rightTress := generateTree(i+1, n)
		for _, left := range leftTress {
			for _, right := range rightTress {
				node := &TreeNode{Val: i, Left: left, Right: right}
				res = append(res, node)
			}
		}
	}

	remark[start][n] = res
	return res
}
