package goutils

// SliceInInt slice 里是否存在 element
func SliceInInt(element int, slice []int) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInUint slice 里是否存在 element
func SliceInUint(element uint, slice []uint) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInInt8 slice 里是否存在 element
func SliceInInt8(element int8, slice []int8) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInUint8 slice 里是否存在 element
func SliceInUint8(element uint8, slice []uint8) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInInt32 slice 里是否存在 element
func SliceInInt32(element int32, slice []int32) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInUint32 slice 里是否存在 element
func SliceInUint32(element uint32, slice []uint32) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInInt64 slice 里是否存在 element
func SliceInInt64(element int64, slice []int64) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInUint64 slice 里是否存在 element
func SliceInUint64(element uint64, slice []uint64) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInFloat32 slice 里是否存在 element
func SliceInFloat32(element float32, slice []float32) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInFloat64 slice 里是否存在 element
func SliceInFloat64(element float64, slice []float64) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInStr slice 里是否存在 element
func SliceInStr(element string, slice []string) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceUniqueInt 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueInt(slice *[]int) {
	var (
		mg = map[int]struct{}{}
		i  int
		k  int
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueUint 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueUint(slice *[]uint) {
	var (
		mg = map[uint]struct{}{}
		i  int
		k  uint
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueInt8 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueInt8(slice *[]int8) {
	var (
		mg = map[int8]struct{}{}
		i  int
		k  int8
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueUint8 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueUint8(slice *[]uint8) {
	var (
		mg = map[uint8]struct{}{}
		i  int
		k  uint8
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueInt32 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueInt32(slice *[]int32) {
	var (
		mg = map[int32]struct{}{}
		i  int
		k  int32
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueUint32 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueUint32(slice *[]uint32) {
	var (
		mg = map[uint32]struct{}{}
		i  int
		k  uint32
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueInt64 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueInt64(slice *[]int64) {
	var (
		mg = map[int64]struct{}{}
		i  int
		k  int64
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueUint64 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueUint64(slice *[]uint64) {
	var (
		mg = map[uint64]struct{}{}
		i  int
		k  uint64
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueStr 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueStr(slice *[]string) {
	var (
		mg = map[string]struct{}{}
		i  int
		k  string
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceSplitInt 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitInt(slice []int, n int) (as [][]int) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		as = append(as, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				as = append(as, slice)
			}
			break
		}
		as = append(as, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitUint 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitUint(slice []uint, n int) (as [][]uint) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		as = append(as, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				as = append(as, slice)
			}
			break
		}
		as = append(as, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitInt8 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitInt8(slice []int8, n int) (as [][]int8) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		as = append(as, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				as = append(as, slice)
			}
			break
		}
		as = append(as, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitUint8 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitUint8(slice []uint8, n int) (as [][]uint8) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		as = append(as, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				as = append(as, slice)
			}
			break
		}
		as = append(as, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitInt32 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitInt32(slice []int32, n int) (as [][]int32) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		as = append(as, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				as = append(as, slice)
			}
			break
		}
		as = append(as, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitUint32 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitUint32(slice []uint32, n int) (as [][]uint32) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		as = append(as, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				as = append(as, slice)
			}
			break
		}
		as = append(as, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitInt64 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitInt64(slice []int64, n int) (as [][]int64) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		as = append(as, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				as = append(as, slice)
			}
			break
		}
		as = append(as, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitUint64 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitUint64(slice []uint64, n int) (as [][]uint64) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		as = append(as, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				as = append(as, slice)
			}
			break
		}
		as = append(as, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitStr 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitStr(slice []string, n int) (as [][]string) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		as = append(as, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				as = append(as, slice)
			}
			break
		}
		as = append(as, slice[:n])
		slice = slice[n:]
	}
	return
}
