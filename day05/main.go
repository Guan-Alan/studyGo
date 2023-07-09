package main

import (
	"fmt"
	"sort"
)

var nums = [][]int{
	{1, 2, 3},
	{4, 5, 6},
}

var nums1 = [][]int{
	{1},
}

func main() {
	fmt.Println(matrixSum(nums))
	fmt.Println(matrixSum1(nums))
}

func matrixSum1(nums [][]int) int {
	res := 0
	m := len(nums)
	n := len(nums[0])
	for i := 0; i < m; i++ {
		sort.Ints(nums[i])
	}

	for j := 0; j < n; j++ {
		maxVal := 0
		for i := 0; i < m; i++ {
			if nums[i][j] > maxVal {
				maxVal = nums[i][j]
			}
		}
		res += maxVal
	}
	return res
}

func matrixSum(nums [][]int) int {
	return GetNumsMax(nums, 0)
}

func GetNumsMax(nums [][]int, curScore int) int {
	var newNums [][]int
	var score int
	for _, num := range nums {
		line, m := GetLineMax(num)
		if m > score {
			score = m
		}
		if len(line) == 0 {
			continue
		}
		newNums = append(newNums, line)
	}
	if len(newNums) != 0 {
		return GetNumsMax(newNums, curScore+score)
	}
	return curScore + score
}

func GetLineMax(line []int) (newLine []int, max int) {
	maxIndex := -1
	for index, i := range line {
		if i > max {
			max = i
			maxIndex = index
		}
	}
	if maxIndex == -1 {
		return
	}

	newLine = append(line[:maxIndex], line[maxIndex+1:]...)
	return
}
