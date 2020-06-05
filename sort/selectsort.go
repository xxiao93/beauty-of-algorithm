package sort

import "fmt"

func SelectSort(a []int, n int)  {
	if n <= 1 {
		return
	}

	for i:=0; i<n-1; i++ {
		minIndex := i
		for j:=i+1; j<n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}

func StartSelectSort()  {
	a := []int{5, 3, 7, 8, 2, 9, 4, 11, 6, 1}
	SelectSort(a, 10)

	for _, v := range(a) {
		fmt.Println(v)
	}
}
