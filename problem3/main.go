package main

import (
	"fmt"
	"problem3/pack3"
)

func main() {
	var st string = "aaaabbaaccde"
	res := pack3.MaxChar(st)
	fmt.Printf("The maximum number of characters repeated in a given string is %s", res)
}
