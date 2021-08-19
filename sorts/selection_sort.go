package main

import "fmt"

func min(a []int) (idx, n int) {
	minI, minV := 0, a[0]
	for i, v := range a {
		if v < minV {
			minI, minV = i, v
		}
	}

	return minI, minV
}

func selectionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		minI, minV := min(a[i:len(a)])
		if minV < a[i] {
			a[i], a[i+minI] = minV, a[i]
		}
	}

	return a
}

func main() {
	a := []int{2, 4, 5, 1, 3}
	fmt.Println(selectionSort(a))
}
