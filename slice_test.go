package goutils

import "testing"

func TestInSliceInt(t *testing.T) {
	t.Log(InSliceInt(int(1), []int{1, 2, 3}))
}

func TestSliceUniqueInt(t *testing.T) {
	s := []int{1, 2, 3, 1}

	SliceUniqueInt(&s)

	t.Log(s)
}

func TestSliceSplitInt(t *testing.T) {
	t.Log(SliceSplitInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, 3))
}
