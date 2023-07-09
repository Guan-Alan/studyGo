package main

import "fmt"

func main() {
	split := maximumEvenSplit(28)
	fmt.Println(split)
}

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 == 1 {
		return []int64{}
	}

	var i int64 = 2
	var sum int64 = 0
	res := make([]int64, 0, 0)
	for ; sum < finalSum; i += 2 {
		sum += i
		res = append(res, i)
		if sum == finalSum {
			return res
		}

		if sum > finalSum {
			for index, item := range res {
				if item == sum-finalSum {
					res = append(res[:index], res[index+1:]...)
					return res
				}
			}
		}
	}
	return []int64{}
}
