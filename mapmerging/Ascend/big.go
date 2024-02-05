package Ascend

func HighestValue(map3 map[int]int) int {
	highValue := 0
	highkey := 0
	for key := range map3 {
		if map3[key] > highValue {
			highValue = map3[key]
			highkey = key
		}
	}
	return highkey

}
