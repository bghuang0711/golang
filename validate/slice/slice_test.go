package slice

import "testing"

func TestSlice(t *testing.T) {
	s0 := []int{0, 1, 2, 3, 4}
	Append(s0)
	t.Log(s0)
	Modify(s0)
	t.Log(s0)
}
