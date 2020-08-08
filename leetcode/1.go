package leetcode

import "fmt"

func minArray(numbers []int) int {
	length := len(numbers)

	if length == 1 {
		return numbers[0]
	}

	for i:=0; i<length - 1; i++ {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
	}

	return numbers[0]
}

func Test_1() {
	numbers := []int{2, 2, 2, 0, 1}
	res := minArray(numbers)
	fmt.Println(res)
}