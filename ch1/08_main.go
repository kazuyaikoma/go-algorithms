package main

import "fmt"

func cipher(r rune) string {
	u := r
	if []rune("a")[0] <= r && r <= []rune("z")[0] {
		u = 219 - r
	}

	return string(u)
}

func main() {
	s := "This is a pen."

	for _, v := range s {
		fmt.Printf(cipher(v))
	}
	fmt.Println("")
}
