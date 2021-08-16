package main

import "fmt"

func bubbleSort(a []int) []int {
	var sorted bool
	for {
		sorted = true
		for i := 0; i < len(a)-1; i++ {
			j := i + 1
			if a[j] < a[i] {
				sorted = false
				a[i], a[j] = a[j], a[i]
			}
		}
		if sorted {
			return a
		}
	}
}

func main() {
	a := []int{2, 4, 5, 1, 3}
	fmt.Println(bubbleSort(a))
}
