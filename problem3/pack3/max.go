package pack3

func MaxChar(st string) string {
	var name = []rune(st)
	var res = string(name[0])
	count := 0
	for i := 0; i < len(name)-1; i++ {
		curr_count := 1
		for j := i + 1; j < len(name); j++ {
			if name[i] != name[j] {
				//break
				curr_count = 1
			}
			curr_count++

		}
		if curr_count > count {
			count = curr_count
			res = string(name[i])

		}

	}
	return res
}
