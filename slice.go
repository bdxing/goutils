package goutils

// InSlice group 里是否存在 element。
// 注意：只支持的 Slice 基础类型，其他类型直接返回 false
func InSlice(element, group interface{}) bool {
	switch group.(type) {
	case []int:
		if e, s := element.(int); s {
			return InSliceInt(e, group.([]int))
		}
	case []uint:
		if e, s := element.(uint); s {
			return InSliceUint(e, group.([]uint))
		}
	case []int32:
		if e, s := element.(int32); s {
			return InSliceInt32(e, group.([]int32))
		}
	case []uint32:
		if e, s := element.(uint32); s {
			return InSliceUint32(e, group.([]uint32))
		}
	case []int64:
		if e, s := element.(int64); s {
			return InSliceInt64(e, group.([]int64))
		}
	case []uint64:
		if e, s := element.(uint64); s {
			return InSliceUint64(e, group.([]uint64))
		}
	case []float32:
		if e, s := element.(float32); s {
			return InSliceFloat32(e, group.([]float32))
		}
	case []float64:
		if e, s := element.(float64); s {
			return InSliceFloat64(e, group.([]float64))
		}
	case []string:
		if e, s := element.(string); s {
			return InSliceString(e, group.([]string))
		}
	}
	return false
}

// InSliceInt group 里是否存在 element
func InSliceInt(element int, group []int) bool {
	for k := range group {
		if group[k] == element {
			return true
		}
	}
	return false
}

// InSliceUint group 里是否存在 element
func InSliceUint(element uint, group []uint) bool {
	for k := range group {
		if group[k] == element {
			return true
		}
	}
	return false
}

// InSliceInt32 group 里是否存在 element
func InSliceInt32(element int32, group []int32) bool {
	for k := range group {
		if group[k] == element {
			return true
		}
	}
	return false
}

// InSliceUint32 group 里是否存在 element
func InSliceUint32(element uint32, group []uint32) bool {
	for k := range group {
		if group[k] == element {
			return true
		}
	}
	return false
}

// InSliceInt64 group 里是否存在 element
func InSliceInt64(element int64, group []int64) bool {
	for k := range group {
		if group[k] == element {
			return true
		}
	}
	return false
}

// InSliceUint64 group 里是否存在 element
func InSliceUint64(element uint64, group []uint64) bool {
	for k := range group {
		if group[k] == element {
			return true
		}
	}
	return false
}

// InSliceFloat32 group 里是否存在 element
func InSliceFloat32(element float32, group []float32) bool {
	for k := range group {
		if group[k] == element {
			return true
		}
	}
	return false
}

// InSliceFloat64 group 里是否存在 element
func InSliceFloat64(element float64, group []float64) bool {
	for k := range group {
		if group[k] == element {
			return true
		}
	}
	return false
}

// InSliceString group 里是否存在 element
func InSliceString(element string, group []string) bool {
	for k := range group {
		if group[k] == element {
			return true
		}
	}
	return false
}
