package pack1

func Duplicate(sl []int) []int {
	var final_slice []int
	map1 := make(map[int]int)
	for _, num := range sl {
		map1[num] = map1[num] + 1
	}
	for num, count := range map1 {
		if count > 1 {
			final_slice = append(final_slice, num)
		}
	}

	return final_slice
}
