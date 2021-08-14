package main

import "fmt"

func reverse(s string) string {
	t := make([]rune, len(s))
	for i, c := range s {
		t[len(s)-i-1] = c
	}
	return string(t)
}

func main() {
	fmt.Println(reverse("stressed"))
}
