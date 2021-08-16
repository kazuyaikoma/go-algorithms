package main

import "fmt"

func arrange(s string) string {
	t := []rune(s)
	return string(t[0]) + string(t[2]) + string(t[4]) + string(t[6])
}

func main() {
	fmt.Println(arrange("パタトクカシーー"))
}
