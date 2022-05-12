package slice

func Append(s []int) {
	s = append(s, 5)
}

func Modify(s []int) {
	for i := range s {
		s[i] += 5
	}
}
