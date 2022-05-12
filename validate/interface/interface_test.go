package _interface

import "testing"

func TestInterface(t *testing.T) {
	var (
		t0 T
		t1 *T
		t2 T1
		i0 interface{} = t0
		i1 interface{} = t1
		i2 interface{} = t2
	)
	t.Log(i0 == t0, i0 == nil)
	t.Log(i1 == t1, i1 == nil)
	t.Log(i2 == t2, i2 == nil)
}
