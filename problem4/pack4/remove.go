package pack4

func Rem(slice []int, number int) []int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == number {
			copy(slice[i:], slice[i+1:])
		}
	}
	return slice[:len(slice)-1]

}
