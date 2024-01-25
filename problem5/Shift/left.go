package Shift

import "fmt"

func Positions(slice []int, num int) []int {
	length := len(slice)
	if length == 0 {
		fmt.Println("Rotation is not needed")
	}
	if num > length {
		num = num % length
	}
	if num == 0 {
		fmt.Println("Rotation is not needed")
	}
	return append(slice[num:], slice[:num]...)

}
