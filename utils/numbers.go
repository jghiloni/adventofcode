package utils

import "math"

type RealNumber interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int |
		~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint |
		~float32 | ~float64
}

type NaturalNumber interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint
}

func Abs[T RealNumber](num T) T {
	if num < 0 {
		return -num
	}

	return num
}

func Base10LeftShift[T NaturalNumber](num T, pow10 int) uint64 {
	a := uint64(Abs(num))
	for range pow10 {
		a *= 10
	}

	return a
}

func GetNumDigits[T NaturalNumber](num T) int {
	return int(math.Ceil(math.Log10(float64(num + 1))))
}
