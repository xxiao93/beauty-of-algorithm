package sort

import "fmt"

/*
搜索旋转排序数组 leetcode 33 题目

list = [4, 5, 6, 7, 0, 1, 2, 3]   target = 4
return 0 
else
没有找到 
return -1

时间复杂度 O(logN) 二分查找
*/

func BinarySearchForCircularArray(array []int, target int) int {
	l, r := 0, len(array) - 1

	if len(array) == 0 {
		return -1
	}

	for l <= r {
		mid := (l + r) / 2
		if array[mid] == target {
			return mid
		}

		// 左半边升序
		if array[l] < array[mid] {
			if array[l] <= target && target < array[mid] {
				r = mid - 1
			}

			if target < array[l] || target > array[mid] {
				l = mid + 1
			}
		} else {
		// 右半边升序
			if array[mid] < target && target <= array[r] {
				l = mid + 1
			}

			if target > array[r] || target < array[mid] {
				r = mid - 1
			}
		}
	}
	
	return -1
}

func StartBinarySearchForCircularArray() {
	array := []int{4, 5, 6, 7, 0, 1, 2, 3}
	target := 2

	res := BinarySearchForCircularArray(array, target)
	fmt.Println(res)
}