package main

import (
	"assignment4/pack9"
	"fmt"
)

func main() {
	map1 := pack9.SetItem()
	message := pack9.Retrieve(map1, "Item1")
	fmt.Println(message)
	deletemsg := pack9.DeleteKey(map1, "Item4")
	fmt.Println(deletemsg)
	fmt.Println(map1)

}
