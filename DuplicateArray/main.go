package main

import "fmt"

func main() {
	var arr = [10]int{10, 20, 30, 40, 50, 60, 10, 20, 40, 50}
	len := duplicate(arr)
	fmt.Println(len)

}
func duplicate(arr [10]int) int {
	var new_arr []int
	map1 := make(map[int]int)
	for _, num := range arr {
		map1[num] = map1[num] + 1
	}
	fmt.Println(map1)
	for key := range map1 {
		new_arr = append(new_arr, key)
	}
	return len(new_arr)

}
