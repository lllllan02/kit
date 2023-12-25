package kit

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Abs 取绝对值。
func Abs[T constraints.Float | constraints.Integer](value T) T {
	if value < 0 {
		return -value
	}
	return value
}

// Average 求集合平均值。
func Average[T constraints.Float | constraints.Integer](slice []T) T {
	var sum T

	for _, item := range slice {
		sum += item
	}

	return sum / T(len(slice))
}

// AverageBy 根据 iteratee 函数对集合求平均值。
func AverageBy[T any, U constraints.Float | constraints.Integer](slice []T, iteratee func(T) U) U {
	var sum U

	for _, item := range slice {
		sum += iteratee(item)
	}

	return sum / U(len(slice))
}

// Clamp 将给定值限制在一个区间内。
func Clamp[T constraints.Ordered](value T, min T, max T) T {
	return If(value < min, min).ElseIf(value > max, max).Else(value)
}

// Max 搜索集合的最大值。当集合为空时返回零值。
func Max[T constraints.Ordered](slice ...T) T {
	if len(slice) == 0 {
		return Empty[T]()
	}

	max := slice[0]

	for i := 1; i < len(slice); i++ {
		item := slice[i]

		max = Ternary(item > max, item, max)
	}

	return max
}

// MaxBy 使用给定的比较函数搜索集合的最大值。
// 如果集合的几个值等于最大值，则返回第一个这样的值。当集合为空时返回零值。
//
// 比较函数，当传入的第一个参数大于第二个参数时返回 true。
func MaxBy[T any](comparison func(T, T) bool, slice ...T) T {
	if len(slice) == 0 {
		return Empty[T]()
	}

	max := slice[0]

	for i := 1; i < len(slice); i++ {
		item := slice[i]

		max = Ternary(comparison(item, max), item, max)
	}

	return max
}

// Min 搜索集合的最小值。当集合为空时返回零值。
func Min[T constraints.Ordered](slice ...T) T {
	if len(slice) == 0 {
		return Empty[T]()
	}

	min := slice[0]

	for i := 1; i < len(slice); i++ {
		item := slice[i]

		min = Ternary(item < min, item, min)
	}

	return min
}

// MinBy 使用给定的比较函数搜索集合的最小值。
// 如果集合的几个值等于最小值，则返回第一个这样的值。当集合为空时返回零值。
//
// 比较函数，当传入的第一个参数小于第二个参数时返回 true。
func MinBy[T any](comparison func(T, T) bool, slice ...T) T {
	if len(slice) == 0 {
		return Empty[T]()
	}

	min := slice[0]

	for i := 1; i < len(slice); i++ {
		item := slice[i]

		min = Ternary(comparison(item, min), item, min)
	}

	return min
}

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

// Round 对浮点数保留 precision 位小数。
func Round[T constraints.Float](value T, precision int) T {
	ratio := math.Pow(10, float64(precision))
	res := math.Round(float64(value)*ratio) / ratio
	return T(res)
}

// Sum 对集合进行求和。
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](slice []T) T {
	var sum T = 0
	for _, val := range slice {
		sum += val
	}
	return sum
}

// SumBy 根据 iteratee 函数对集合进行求和。
func SumBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](slice []T, iteratee func(item T) R) R {
	var sum R = 0
	for _, item := range slice {
		sum = sum + iteratee(item)
	}
	return sum
}
