package main

import "fmt"

func main() {
	fmt.Println(_numTreesV1(3))
}

// _numTreesV1 正确的
func _numTreesV1(n int) int {
	nums := make([]int, n+1)
	nums[0], nums[1] = 1, 1
	// 递推n个节点，每个节点有多少种解法
	for i := 2; i <= n; i++ {
		// j代表当前是哪个值作为根节点，用这个j的左子树可能性数量乘以右子树的可能性数量，就是j作为根节点时的bst树数量
		// i表示当前一共有几个元素，eg：3
		// 当i=3时，可用的元素就是[1,2,3]
		// 循环累加 左边两个元素，右边零个元素 + 左边1个元素，右边1个元素 + 左边零个元素，右边两个元素的情况，就是三个元素能构建的所有二叉树的情况
		// 实际也是 1，2，3分别作为根节点时能够成的二叉树数量
		for j := 1; j <= i; j++ {
			nums[i] += nums[j-1] * nums[i-j]
		}
	}

	return nums[n]
}

// numTrees 解法超时
var remark [][]int

func numTrees(n int) int {
	if n <= 1 {
		return 1
	}
	remark = make([][]int, n+1, n+1)
	for index, _ := range remark {
		remark[index] = make([]int, n+1, n+1)
	}
	return Trees(1, n)
}

func Trees(start, end int) int {
	if start >= end {
		return 1
	}

	res := 0
	for i := start; i <= end; i++ {
		lefts := Trees(start, i-1)
		rights := Trees(i+1, end)
		res += lefts * rights
	}
	return res
}
