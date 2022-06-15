package goutils

import "testing"

func TestInSlice(t *testing.T) {
	t.Log(InSlice(int(11), []int{1, 2, 3}))
}

func TestInSliceInt(t *testing.T) {
	t.Log(InSliceInt(int(1), []int{1, 2, 3}))
}

func TestInSliceUint(t *testing.T) {
	t.Log(InSliceUint(uint(1), []uint{1, 2, 3}))
}

func TestInSliceInt32(t *testing.T) {
	t.Log(InSliceInt32(int32(1), []int32{1, 2, 3}))
}

func TestInSliceUint32(t *testing.T) {
	t.Log(InSliceUint32(uint32(1), []uint32{1, 2, 3}))
}

func TestInSliceInt64(t *testing.T) {
	t.Log(InSliceInt64(int64(1), []int64{1, 2, 3}))
}

func TestInSliceUint64(t *testing.T) {
	t.Log(InSliceUint64(uint64(1), []uint64{1, 2, 3}))
}

func TestInSliceFloat32(t *testing.T) {
	t.Log(InSliceFloat32(float32(1), []float32{1, 2, 3}))
}

func TestInSliceFloat64(t *testing.T) {
	t.Log(InSliceFloat64(float64(1), []float64{1, 2, 3}))
}

func TestInSliceString(t *testing.T) {
	t.Log(InSliceString(string("a"), []string{"a", "b", "c"}))
}
