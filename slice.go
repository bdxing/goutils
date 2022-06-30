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

// SliceUniqueInt8 去除 array 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueInt8(array *[]int8) {
	var (
		mg = map[int8]struct{}{}
		i  int
		k  int8
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

// SliceUniqueUint8 去除 array 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueUint8(array *[]uint8) {
	var (
		mg = map[uint8]struct{}{}
		i  int
		k  uint8
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

// SliceUniqueInt32 去除 array 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueInt32(array *[]int32) {
	var (
		mg = map[int32]struct{}{}
		i  int
		k  int32
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

// SliceUniqueUint32 去除 array 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueUint32(array *[]uint32) {
	var (
		mg = map[uint32]struct{}{}
		i  int
		k  uint32
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

// SliceUniqueInt64 去除 array 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueInt64(array *[]int64) {
	var (
		mg = map[int64]struct{}{}
		i  int
		k  int64
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

// SliceUniqueUint64 去除 array 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueUint64(array *[]uint64) {
	var (
		mg = map[uint64]struct{}{}
		i  int
		k  uint64
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

// SliceSplitInt 把 array 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitInt(array []int, n int) (as [][]int) {
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

// SliceSplitUint 把 array 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitUint(array []uint, n int) (as [][]uint) {
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

// SliceSplitInt8 把 array 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitInt8(array []int8, n int) (as [][]int8) {
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

// SliceSplitUint8 把 array 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitUint8(array []uint8, n int) (as [][]uint8) {
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

// SliceSplitInt32 把 array 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitInt32(array []int32, n int) (as [][]int32) {
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

// SliceSplitUint32 把 array 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitUint32(array []uint32, n int) (as [][]uint32) {
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

// SliceSplitInt64 把 array 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitInt64(array []int64, n int) (as [][]int64) {
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

// SliceSplitUint64 把 array 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitUint64(array []uint64, n int) (as [][]uint64) {
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
