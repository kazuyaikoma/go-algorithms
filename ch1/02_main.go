package main

import "fmt"

func connect(s1, s2 string) string {
	t := make([]rune, len(s1)*2)

	r1, r2 := []rune(s1), []rune(s2)

	for i := 0; i < len(r1); i++ {
		t[i*2] = r1[i]
		t[i*2+1] = r2[i]
	}

	return string(t)
}

func main() {
	fmt.Println(connect("パトカー", "タクシー"))
}
