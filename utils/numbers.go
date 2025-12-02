package utils

type RealNumber interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int |
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint |
	~float32 | ~float64
}

func Abs[T RealNumber](num T) T {
	if num < 0 {
		return -num
	}

	return num
}
