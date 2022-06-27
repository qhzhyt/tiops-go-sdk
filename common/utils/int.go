package utils

func Int64Map2List(m map[int64]bool) []int64 {
	var l []int64
	for s, b := range m {
		if b {
			l = append(l, s)
		}
	}
	return l
}

func Int64List2Map(l []int64) map[int64]bool {
	m := map[int64]bool{}
	for _, i := range l {
		m[i] = true
	}
	return m
}

func Int64ListDeduplicate(l []int64) []int64 {
	return Int64Map2List(Int64List2Map(l))
}

func Int64ListSum(l []int64) int64 {
	var s int64

	for _, i := range l {
		s += i
	}

	return s
}

func Int64ListHash(l []int64) int64 {
	return Int64ListSum(l)
}
