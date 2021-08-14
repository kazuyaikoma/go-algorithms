package main

import (
	"fmt"
	"math"
)

func charNgram(s []rune, n int) [][]rune {
	t := make([][]rune, len(s)-n+1)
	for i := 0; i < len(s)-n+1; i++ {
		t[i] = s[i : i+n]
	}

	return t
}

func union(a, b [][]rune) []string {
	maxL := math.Max(float64(len(a)), float64(len(b)))
	m := make(map[string]interface{}, int(maxL))

	for _, s := range a {
		m[string(s)] = struct{}{}
	}
	for _, s := range b {
		m[string(s)] = struct{}{}
	}

	ret := make([]string, len(m))
	for key := range m {
		ret = append(ret, key)
	}

	return ret
}

func main() {
	s1 := "paraparaparadise"
	s2 := "paragraph"

	X := charNgram([]rune(s1), 2)
	Y := charNgram([]rune(s2), 2)

	unioned := union(X, Y)
	for _, u := range unioned {
		fmt.Println(u)
	}
	// fmt.Println(union(X, Y))
}
