package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func fisherYetesShuffle(s string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	t := []rune(s)

	var j, k int
	for i := len(s) - 1; 0 < i; i-- {
		k = i - 1
		if k <= 0 {
			continue
		}
		j = r.Intn(i)

		t[i], t[j] = t[j], t[i]
	}

	return string(t)
}

func main() {
	s := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

	var u string
	var tmp []rune
	for _, v := range strings.Split(s, " ") {
		if 4 < len(v) {
			tmp = []rune(v)
			v = string(tmp[0])
			v += fisherYetesShuffle(string(tmp[1 : len(tmp)-1]))
			v += string(tmp[len(tmp)-1])
		}

		u += v + " "
	}

	fmt.Println(u)
}
