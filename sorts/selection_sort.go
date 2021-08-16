package main

import "fmt"

func min(a []int) (idx, n int) {
	idx, n = 0, a[0]

	for i, v := range a {
		if v < n {
			idx, n = i, v
		}
	}

	return
}

func selectionSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		j, _ := min(a[i:])
		if i != i+j {
			a[i], a[i+j] = a[i+j], a[i]
		}
	}

	return a
}

func main() {
	a := []int{2, 4, 5, 1, 3}
	fmt.Println(selectionSort(a))
}
