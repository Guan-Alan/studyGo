package main

func main() {

}

// twoSum 双指针，适合本题的有序数组情况
func twoSum(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1
	for low < high {
		sum := numbers[low] + numbers[high]
		if sum == target {
			return []int{low + 1, high + 1}
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return []int{-1, -1}
}

// twoSum 利用hash表来找对应的差值，适合无序数组情况
func _twoSum(numbers []int, target int) []int {
	sumMap := make(map[int]int, len(numbers))
	for index1, number := range numbers {
		if index2, ok := sumMap[number]; ok {
			return []int{index2, index1 + 1}
		}

		sumMap[target-number] = index1 + 1
	}
	return []int{}
}
