package goutils

import "testing"

func TestInSliceInt(t *testing.T) {
	t.Log(InSliceInt(int(1), []int{1, 2, 3}))
}

func TestInSliceUint(t *testing.T) {
	t.Log(InSliceUint(uint(1), []uint{1, 2, 3}))
}

func TestInSliceInt8(t *testing.T) {
	t.Log(InSliceInt8(int8(1), []int8{1, 2, 3}))
}

func TestInSliceUint8(t *testing.T) {
	t.Log(InSliceUint8(uint8(1), []uint8{1, 2, 3}))
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

func TestInSliceStr(t *testing.T) {
	t.Log(InSliceStr(string("a"), []string{"a", "b", "c"}))
}

func TestSliceUniqueInt(t *testing.T) {
	s := []int{1, 2, 3, 1}

	SliceUniqueInt(&s)

	t.Log(s)
}

func TestSliceUniqueUint(t *testing.T) {
	s := []uint{1, 2, 3, 1}

	SliceUniqueUint(&s)

	t.Log(s)
}

func TestSliceUniqueStr(t *testing.T) {
	s := []string{"a", "a", "b", "c"}

	SliceUniqueStr(&s)

	t.Log(s)
}

func TestSliceSplitStr(t *testing.T) {
	t.Log(SliceSplitStr([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}, 3))
}
