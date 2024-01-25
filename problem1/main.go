package main

import (
	"assignment2/pack1"
	"fmt"
)

func main() {
	var sl = []int{10, 20, 20, 20, 20, 10, 10}
	final_slice := pack1.Duplicate(sl)
	fmt.Println(final_slice)
}
