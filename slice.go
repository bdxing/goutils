package goutils

// InSliceInt array 里是否存在 element
func InSliceInt(element int, array []int) bool {
	for k := range array {
		if array[k] == element {
			return true
		}
	}
	return false
}

// InSliceUint array 里是否存在 element
func InSliceUint(element uint, array []uint) bool {
	for k := range array {
		if array[k] == element {
			return true
		}
	}
	return false
}

// InSliceInt8 array 里是否存在 element
func InSliceInt8(element int8, array []int8) bool {
	for k := range array {
		if array[k] == element {
			return true
		}
	}
	return false
}

// InSliceUint8 array 里是否存在 element
func InSliceUint8(element uint8, array []uint8) bool {
	for k := range array {
		if array[k] == element {
			return true
		}
	}
	return false
}

// InSliceInt32 array 里是否存在 element
func InSliceInt32(element int32, array []int32) bool {
	for k := range array {
		if array[k] == element {
			return true
		}
	}
	return false
}

// InSliceUint32 array 里是否存在 element
func InSliceUint32(element uint32, array []uint32) bool {
	for k := range array {
		if array[k] == element {
			return true
		}
	}
	return false
}

// InSliceInt64 array 里是否存在 element
func InSliceInt64(element int64, array []int64) bool {
	for k := range array {
		if array[k] == element {
			return true
		}
	}
	return false
}

// InSliceUint64 array 里是否存在 element
func InSliceUint64(element uint64, array []uint64) bool {
	for k := range array {
		if array[k] == element {
			return true
		}
	}
	return false
}

// InSliceFloat32 array 里是否存在 element
func InSliceFloat32(element float32, array []float32) bool {
	for k := range array {
		if array[k] == element {
			return true
		}
	}
	return false
}

// InSliceFloat64 array 里是否存在 element
func InSliceFloat64(element float64, array []float64) bool {
	for k := range array {
		if array[k] == element {
			return true
		}
	}
	return false
}

// InSliceStr array 里是否存在 element
func InSliceStr(element string, array []string) bool {
	for k := range array {
		if array[k] == element {
			return true
		}
	}
	return false
}

// SliceUniqueInt 去除 array 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueInt(array *[]int) {
	var (
		mg = map[int]struct{}{}
		i  int
		k  int
	)
	for i = range *array {
		mg[(*array)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*array)[i] = k
		i++
	}
	*array = (*array)[:i]
}

// SliceUniqueUint 去除 array 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueUint(array *[]uint) {
	var (
		mg = map[uint]struct{}{}
		i  int
		k  uint
	)
	for i = range *array {
		mg[(*array)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*array)[i] = k
		i++
	}
	*array = (*array)[:i]
}

// SliceUniqueStr 去除 array 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueStr(array *[]string) {
	var (
		mg = map[string]struct{}{}
		i  int
		k  string
	)
	for i = range *array {
		mg[(*array)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*array)[i] = k
		i++
	}
	*array = (*array)[:i]
}

// SliceSplitStr 把 array 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitStr(array []string, n int) (as [][]string) {
	if len(array) == 0 {
		return
	}
	if n == 0 {
		as = append(as, array)
		return
	}
	for {
		if len(array) < n {
			if len(array) != 0 {
				as = append(as, array)
			}
			break
		}
		as = append(as, array[:n])
		array = array[n:]
	}
	return
}
