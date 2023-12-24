package kit

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Range 根据给定长度返回一个 int 数组。
func Range(num int) []int {
	length := If(num < 0, -num).Else(num)
	result := make([]int, length)
	step := If(num < 0, -1).Else(1)

	for i, j := 0, 0; i < length; i, j = i+1, j+step {
		result[i] = j
	}

	return result
}

// RangeFrom 根据给定的起始位置和长度返回一个 int 数组。
func RangeFrom[T constraints.Integer | constraints.Float](start T, elementNum int) []T {
	length := If(elementNum < 0, -elementNum).Else(elementNum)
	result := make([]T, length)
	step := If(elementNum < 0, -1).Else(1)

	for i, j := 0, start; i < length; i, j = i+1, j+T(step) {
		result[i] = j
	}

	return result
}

// RangeWithSteps 根据给定的起止位置和步长返回一个数组。
func RangeWithSteps[T constraints.Integer | constraints.Float](start, end, step T) []T {
	result := []T{}
	if start == end || step == 0 {
		return result
	}

	if start < end {
		if step < 0 {
			return result
		}
		for i := start; i < end; i += step {
			result = append(result, i)
		}
		return result
	}

	if step > 0 {
		return result
	}

	for i := start; i > end; i += step {
		result = append(result, i)
	}
	return result
}

// Clamp 将给定值限制在一个区间内。
func Clamp[T constraints.Ordered](value T, min T, max T) T {
	return If(value < min, min).ElseIf(value > max, max).Else(value)
}

// Sum 对集合进行求和。
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T {
	var sum T = 0
	for _, val := range collection {
		sum += val
	}
	return sum
}

// SumBy 根据 iteratee 函数对集合进行求和。
func SumBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection []T, iteratee func(item T) R) R {
	var sum R = 0
	for _, item := range collection {
		sum = sum + iteratee(item)
	}
	return sum
}

// Round 对浮点数保留 precision 位小数。
func Round[T constraints.Float](value T, precision int) T {
	ratio := math.Pow(10, float64(precision))
	res := math.Round(float64(value)*ratio) / ratio
	return T(res)
}
