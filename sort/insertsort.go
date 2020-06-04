package sort

import "fmt"

func InsertSort(a []int, n int)  {
	if n <=1 {
		return
	}

	for i:=1; i<n; i++ {
		value := a[i]
		j := i-1
		for ; j>=0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
}

func StartInsertSort()  {
	a := []int{3, 5, 2, 6, 4, 9}
	InsertSort(a, 6)

	for _, v := range(a) {
		fmt.Println(v)
	}
}