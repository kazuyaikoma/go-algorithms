package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ".", "")

	for _, sp := range strings.Split(s, " ") {
		fmt.Println(len(sp))
	}
}
