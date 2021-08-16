package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 12
	y := "気温"
	z := 22.4

	message := strconv.Itoa(x) + "時の" + y + "は" + strconv.FormatFloat(z, 'f', -1, 64)
	fmt.Println(message)
}
