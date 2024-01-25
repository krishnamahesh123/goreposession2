package main

import (
	"fmt"
	"problem5/Shift"
)

func main() {
	sl := []int{10, 20, 30, 40, 50}
	num := 2
	sl = Shift.Positions(sl, num)
	fmt.Println(sl)

}
