package main

import "fmt"

type Element struct {
	Value     int
	Frequency int
}

func main() {
	var arr = []int{2, 3, 2, 4, 5, 12, 2, 3, 3, 3, 12}
	res := CheckFrequency(arr)
	fmt.Println("The resulting array after sorted is")
	fmt.Println(res)
}
func CheckFrequency(arr []int) []int {
	map1 := make(map[int]int)
	for _, freq := range arr {
		map1[freq] = map1[freq] + 1
	}
	fmt.Println(map1)
	sl := make([]Element, 0, len(map1))
	res := make([]int, 0, len(arr))
	for val, freq := range map1 {
		sl = append(sl, Element{Value: val, Frequency: freq})
	}
	fmt.Println(sl)
	var temp Element
	//var temp2 int

	for i := 0; i < len(sl)-1; i++ {
		for j := i + 1; j < len(sl); j++ {
			if sl[i].Frequency < sl[j].Frequency || (sl[i].Frequency == sl[j].Frequency && sl[i].Value > sl[j].Value) {
				temp = sl[i]
				sl[i] = sl[j]
				sl[j] = temp

			}

		}
	}
	fmt.Println(sl)
	for _, num := range sl {
		for i := 0; i < num.Frequency; i++ {
			res = append(res, num.Value)
		}
	}
	return res
}
