package util

func ToInt64(num int64) *int64 {
	return &num
}

func FindAndDeleteInt64(s []int64, item int64) []int64 {
	index := 0
	for _, i := range s {
		if i != item {
			s[index] = i
			index++
		}
	}
	return s[:index]
}
