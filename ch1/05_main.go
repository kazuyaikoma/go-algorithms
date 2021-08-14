package main

import (
	"fmt"
	"strings"
)

func charNgram(s []rune, n int) [][]rune {
	t := make([][]rune, len(s)-n+1)
	for i := 0; i < len(s)-n+1; i++ {
		t[i] = s[i : i+n]
	}
	return t
}

func wordNgram(s [][]rune, n int) [][][]rune {
	t := make([][][]rune, len(s)-n+1)
	for i := 0; i < len(s)-n+1; i++ {
		t[i] = s[i : i+n]
	}

	return t
}

func main() {
	s := "I am an NLPer"

	fmt.Println("-----文字bi-gram-----")
	t := charNgram([]rune(s), 2)
	for _, v := range t {
		fmt.Println(string(v))
	}

	fmt.Println("-----単語bi-gram-----")
	words := strings.Split(s, " ")
	rWords := make([][]rune, len(words))
	for i, v := range words {
		rWords[i] = []rune(v)
	}

	wbg := wordNgram(rWords, 2)
	for _, v := range wbg {
		for _, w := range v {
			fmt.Printf(string(w) + " ")
		}
		fmt.Println("")
	}
}
