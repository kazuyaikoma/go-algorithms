package main

import "fmt"

func insertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := i - 1; 0 <= j; j-- {
			if a[j+1] < a[j] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func main() {
	a := []int{2, 3, 4, 1, 5}
	fmt.Println(insertionSort(a))
}
