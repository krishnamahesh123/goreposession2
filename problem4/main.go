package main

import (
	"fmt"
	"problem4/pack4"
)

func main() {
	var slice = []int{10, 20, 30, 40, 5}
	number := 20
	slice = pack4.Rem(slice, number)
	fmt.Println(slice)
}
