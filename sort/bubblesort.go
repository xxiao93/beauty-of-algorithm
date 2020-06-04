package sort

import "fmt"

func BubbleSort(a []int, n int) {
	if n <= 1 {
		return 
	}

	for i:=0; i<n; i++ {
		flag := false

		for j:=i; j<n-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}

		}

		if flag == false {
			break
		}

	}
}

func StartBubbleSort() {
	a := []int{3, 2, 5, 7, 9}
	BubbleSort(a, 5)
	for _, v := range(a) {
		fmt.Println(v)
	}
}
