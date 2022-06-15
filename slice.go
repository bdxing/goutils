package goutils

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
