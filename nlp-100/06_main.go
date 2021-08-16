package main

import (
	"fmt"
	"reflect"
)

func charNgram(s []rune, n int) [][]rune {
	t := make([][]rune, len(s)-n+1)
	for i := 0; i < len(s)-n+1; i++ {
		t[i] = s[i : i+n]
	}

	return t
}

func uniqueRuneSlice(s [][]rune) [][]rune {
	ret := [][]rune{}
	for _, v := range s {
		if !exists(ret, v) {
			ret = append(ret, v)
		}
	}

	return ret
}

func exists(set [][]rune, word []rune) bool {
	ex := false
	for _, v := range set {
		if reflect.DeepEqual(v, word) {
			ex = true
			break
		}
	}

	return ex
}

func union(X, Y [][]rune) [][]rune {
	for _, y := range Y {
		if !exists(X, y) {
			X = append(X, y)
		}
	}
	return X
}

func intersection(X, Y [][]rune) [][]rune {
	ret := [][]rune{}
	for _, v := range X {
		if exists(Y, v) {
			ret = append(ret, v)
		}
	}

	return ret
}

// 差集合(X - Y)
func difference(X, Y [][]rune) [][]rune {
	ret := [][]rune{}
	for _, v := range X {
		if !exists(Y, v) {
			ret = append(ret, v)
		}
	}

	return ret
}

func runeToString(r [][]rune) (s []string) {
	for _, r := range r {
		s = append(s, string(r))
	}
	return
}

func main() {
	s1 := "paraparaparadise"
	s2 := "paragraph"
	biGram := "se"

	X := charNgram([]rune(s1), 2)
	Y := charNgram([]rune(s2), 2)
	uniqueX := uniqueRuneSlice(X)
	unioned := union(uniqueX, Y)

	fmt.Println("-----集合X-----")
	fmt.Println(runeToString(X))
	fmt.Println("-----集合Y-----")
	fmt.Println(runeToString(Y))

	fmt.Println("-----和集合-----")
	fmt.Println(runeToString(unioned))
	fmt.Println("-----積集合-----")
	fmt.Println(runeToString(intersection(uniqueX, Y)))
	fmt.Println("-----差集合(X-Y)-----")
	fmt.Println(runeToString(difference(uniqueX, Y)))

	fmt.Println("-----集合Xにbi-gram 'se'を含むかどうか-----")
	fmt.Println(exists(X, []rune(biGram)))
	fmt.Println("-----集合Yにbi-gram 'se'を含むかどうか-----")
	fmt.Println(exists(Y, []rune(biGram)))
}
