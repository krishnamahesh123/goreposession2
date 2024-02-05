package pack9

import (
	"fmt"
	"time"
)

type Item struct {
	Name      string
	Expire_at time.Time
}

func SetItem() map[string]Item {
	var item1 = Item{
		Name:      "Milk",
		Expire_at: time.Date(2024, time.February, 6, 12, 0, 0, 0, time.UTC),
	}
	var item2 = Item{
		Name:      "Sugar",
		Expire_at: time.Date(2024, time.February, 6, 12, 0, 0, 0, time.UTC),
	}
	var item3 = Item{
		Name:      "Panner",
		Expire_at: time.Date(2024, time.February, 6, 12, 0, 0, 0, time.UTC),
	}
	var item4 = Item{
		Name:      "Bread",
		Expire_at: time.Date(2024, time.February, 6, 12, 0, 0, 0, time.UTC),
	}
	map1 := make(map[string]Item)
	map1["Item1"] = item1
	map1["Item2"] = item2
	map1["Item3"] = item3
	map1["Item4"] = item4
	fmt.Println(map1)
	return map1

}
func Retrieve(map1 map[string]Item, item1 string) string {
	for key, value := range map1 {
		if key == item1 {
			fmt.Println("The given Item is present in the store")
			if time.Now().After(value.Expire_at) {
				return "key is expired"
			} else {
				return "Key is not expired"
			}

		}

	}
	return "key is not present in the store"
}

func DeleteKey(map1 map[string]Item, givenKey string) string {
	for key := range map1 {
		if key == givenKey {
			delete(map1, key)
			return "key is deleted"
		}
	}
	return "key is not present in the store to delete"
}
