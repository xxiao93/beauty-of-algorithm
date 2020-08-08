package leetcode

import "fmt"

func twoSum(nums []int, target int) []int {
	length := len(nums)
	target_map := make(map[int]int, length)
    for index, value := range(nums) {
		i, ok := target_map[target - value]
		if !ok {
			target_map[value] = index
		} else {
			return []int{index, i}
		}
	}
	return nil
}

func Test_2() {
	nums := []int{2, 4, 5, 8}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}