package main

import "fmt"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++{
		for j := i + 1; j < len(nums); j++{
			for z := j + 1; z < len(nums); z++{
				if nums[i] + nums[j] + nums[z] == 0 {
					triple := []int{nums[i], nums[j], nums[z]}
					result = append(result, triple)
				}
			}
		}
	}
	removeRepeats(result)

	return result
}

func removeRepeats(arrs [][]int) [][]int{
	result := make([][]int, 0)
	for i := 0; i < len(arrs); i++{
		for j := 0; j < len(result); j++{
			if compare(arrs[i], result[j]){
				result = append(result, arrs[i])
			}
		}
	}
	return result
}

func compare(first, second []int)  bool {

	for i := 0; i < len(first); i++ {
		if !inside(first[i], second){
			return false
		}
	}
	return true
}

func inside(n int, arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if n == arr[i] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
