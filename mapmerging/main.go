package main

import (
	"fmt"
	"mapmerging/Ascend"
)

func main() {
	var map1 = map[int]int{1: 10, 2: 20, 3: 30, 4: 40}
	var map2 = map[int]int{1: 50, 6: 60, 7: 70, 8: 80}
	var map3 = make(map[int]int)
	for key1, value1 := range map1 {
		map3[key1] = value1
	}
	for key2, value2 := range map2 {
		if value, ok := map3[key2]; ok {
			map3[key2] = value2 + value
		} else {
			map3[key2] = value2
		}

	}
	fmt.Println(map3)
	highkey := Ascend.HighestValue(map3)
	fmt.Println(highkey)

}
