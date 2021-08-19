package main

import "fmt"

func insertionSort(a []int) []int {
	sorted := true
	for {
		for i := len(a) - 1; 1 <= i; i++ {
			if a[i] < a[i-1] {
				a[i], a[i-1] = a[i-1], a[i]
				sorted = false
			}
		}
		if sorted {
			return a
		}
	}
}

func main() {
	a := []int{2, 3, 4, 1, 5}
	fmt.Println(insertionSort(a))
}
