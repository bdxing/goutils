package goutils

// StrHumpToUnder 驼峰转下划线
func StrHumpToUnder(s string) string {
	sr := []rune(s)

	var upperIndex []int
	for i := range sr {
		if sr[i] < 0x41 || sr[i] > 0x5A {
			continue
		}
		if i == 0x00 {
			sr[i] += 0x20
		} else {
			upperIndex = append(upperIndex, i)
		}
	}

	sr = append(sr, make([]rune, len(upperIndex))...)

	for k, v := range upperIndex {
		v += k
		sr[v] += 0x20
		copy(sr[v+1:], sr[v:])
		sr[v] = 0x5F
	}

	return string(sr)
}

// StrUnderToHump 下划线转驼峰
func StrUnderToHump(s string) string {
	sr := []rune(s)

	var humpIndex []int
	for i := range sr {
		if sr[i] < 0x41 || sr[i] > 0x5A {
			if i == 0x00 {
				sr[i] -= 0x20
			}
		}
		if sr[i] == 0x5F {
			humpIndex = append(humpIndex, i)
		}
	}

	for k, v := range humpIndex {
		v -= k
		copy(sr[v:], sr[v+1:])
		sr[v] -= 0x20
	}

	return string(sr[:len(sr)-len(humpIndex)])
}
