package pack2

func Test(sl []int) []int {
	var final_slice []int
	for i := 0; i < len(sl)-1; i++ {
		for j := i + 1; j < len(sl); j++ {
			if sl[i]+sl[j] == 10 {
				final_slice = append(final_slice, i)
				final_slice = append(final_slice, j)
			}
		}
	}
	if len(final_slice) == 0 {
		final_slice = append(final_slice, -1)
	}
	return final_slice
}
