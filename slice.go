package kit

import (
	"fmt"
	"math/rand"

	"golang.org/x/exp/constraints"
)

// Contain 判断元素是否在集合中。
func Contain[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}
	return false
}

// ContainBy 根据 predicate 判断元素是否在集合中。
func ContainBy[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}
	return false
}

// Count 统计集合中 value 出现的次数。
func Count[T comparable](collection []T, value T) int {
	var count int

	for _, item := range collection {
		if item == value {
			count++
		}
	}

	return count
}

// CountValues 统计集合中元素的出现次数。
func CountValues[T comparable](collection []T) map[T]int {
	result := make(map[T]int)

	for _, item := range collection {
		result[item]++
	}

	return result
}

// Difference 返回两个集合之间的差异。
// 第一个值是 list2 中不存在的元素的集合。
// 第二个值是 list1 中不存在的元素的集合。
func Difference[T comparable](list1 []T, list2 []T) ([]T, []T) {
	left := []T{}
	right := []T{}

	leftSeen := make(map[T]bool, len(list1))
	rightSeen := make(map[T]bool, len(list2))

	for _, item := range list1 {
		leftSeen[item] = true
	}

	for _, item := range list2 {
		rightSeen[item] = true
		if !leftSeen[item] {
			right = append(right, item)
		}
	}

	for _, item := range list1 {
		if !rightSeen[item] {
			left = append(left, item)
		}
	}

	return left, right
}

// Every 判断集合 subset 中的全部元素是否都存在于另一个集合 collection 中。
// 如果 subset 为空，则返回 true。
func Every[T comparable](collection []T, subset []T) bool {
	contain := make(map[T]bool, len(collection))

	for _, item := range collection {
		contain[item] = true
	}

	for _, item := range subset {
		if !contain[item] {
			return false
		}
	}

	return true
}

// EveryBy 如果 predicate 为集合中的所有元素返回 true 或集合 subset 为空，则返回 true。
func EveryBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, v := range collection {
		if !predicate(v) {
			return false
		}
	}

	return true
}

// Fill	使用 initial 填充集合。
func Fill[T Clonable[T]](collection []T, initial T) []T {
	result := make([]T, 0, len(collection))

	for range collection {
		result = append(result, initial.Clone())
	}

	return result
}

// Filter 使用 predicate 函数对集合进行过滤，保留返回值为 true 的所有元素。
func Filter[T any](collection []T, predicate func(int, T) bool) []T {
	result := make([]T, 0, len(collection))

	for i, item := range collection {
		if predicate(i, item) {
			result = append(result, item)
		}
	}

	return result
}

// Find 根据 predicate 在切片中搜索元素。如果找到元素，则返回元素和 true。
func Find[T any](collection []T, predicate func(T) bool) (T, bool) {
	result, _, exist := FindIndexOf(collection, predicate)
	return result, exist
}

// FindDuplicates 返回集合中重复出现的元素去重后的切片。
// 结果值的顺序由它们在集合中出现的顺序决定。
func FindDuplicates[T comparable](collection []T) []T {
	isDupl := make(map[T]bool, len(collection))

	for _, item := range collection {
		_, ok := isDupl[item]
		isDupl[item] = ok
	}

	result := make([]T, 0, len(collection)-len(isDupl))

	for _, item := range collection {
		if isDupl[item] {
			isDupl[item] = false
			result = append(result, item)
		}
	}

	return result
}

// FindDuplicatesBy 返回集合中重复出现的元素去重后的切片。
// 结果值的顺序由它们在数组中出现的顺序决定。
// 它接受 'iteratee'，即为数组中的每个元素调用以生成计算唯一性的标准。
func FindDuplicatesBy[T any, U comparable](collection []T, iteratee func(T) U) []T {
	isDupl := make(map[U]bool, len(collection))

	for _, item := range collection {
		key := iteratee(item)

		_, ok := isDupl[key]
		isDupl[key] = ok
	}

	result := make([]T, 0, len(collection)-len(isDupl))

	for _, item := range collection {
		key := iteratee(item)

		if isDupl[key] {
			isDupl[key] = false
			result = append(result, item)
		}
	}

	return result
}

// FindIndexOf 根据 predicate 搜索切片中的元素并返回索引和 true。
// 如果未找到该元素，则返回 -1 和 false。
func FindIndexOf[T any](collection []T, predicate func(T) bool) (T, int, bool) {
	for i, item := range collection {
		if predicate(item) {
			return item, i, true
		}
	}

	return Empty[T](), -1, false
}

// FindLastIndexOf 根据 predicate 搜索切片中的最后一个元素并返回索引和 true。
// 如果未找到该元素，则返回 -1 和 false。
func FindLastIndexOf[T any](collection []T, predicate func(T) bool) (T, int, bool) {
	for i := len(collection) - 1; i >= 0; i-- {
		if predicate(collection[i]) {
			return collection[i], i, true
		}
	}

	return Empty[T](), -1, false
}

// FindOrElse 根据 predicate 搜索切片中的元素。如果找到则返回元素，否则返回给定的回退值。
func FindOrElse[T any](collection []T, fallback T, predicate func(T) bool) T {
	for _, item := range collection {
		if predicate(item) {
			return item
		}
	}

	return fallback
}

// FindUniques 返回包含集合所有唯一元素的切片。
// 结果值的顺序由它们在集合中出现的顺序决定。
func FindUniques[T comparable](collection []T) []T {
	isDupl := make(map[T]bool, len(collection))

	for _, item := range collection {
		_, ok := isDupl[item]
		isDupl[item] = ok // 出现一次为 false，重复出现为 true
	}

	result := make([]T, 0, len(collection)-len(isDupl))

	for _, item := range collection {
		if !isDupl[item] {
			result = append(result, item)
		}
	}

	return result
}

// FindUniquesBy 返回包含集合所有唯一元素的切片。
// 结果值的顺序由它们在数组中出现的顺序决定。
// 它接受 'iteratee'，即为数组中的每个元素调用以生成计算唯一性的标准。
func FindUniquesBy[T any, U comparable](collecton []T, iteratee func(T) U) []T {
	isDul := make(map[U]bool, len(collecton))

	for _, item := range collecton {
		key := iteratee(item)

		_, ok := isDul[key]
		isDul[key] = ok
	}

	result := make([]T, 0, len(collecton)-len(isDul))

	for _, item := range collecton {
		key := iteratee(item)

		if !isDul[key] {
			result = append(result, item)
		}
	}

	return result
}

// ForEach 遍历集合中的所有元素并执行 iteratee 函数。
func ForEach[T any](collection []T, iteratee func(int, T)) {
	for i, item := range collection {
		iteratee(i, item)
	}
}

// GroupBy 根据提供的 iteratee 函数对集合进行分组。
func GroupBy[T any, U comparable](collection []T, iteratee func(T) U) map[U][]T {
	result := make(map[U][]T)

	for _, item := range collection {
		key := iteratee(item)
		result[key] = append(result[key], item)
	}

	return result
}

// IndexOf 返回在数组中找到值的第一次出现的索引。如果找不到该值则返回-1。
func IndexOf[T comparable](collection []T, element T) int {
	for i, item := range collection {
		if item == element {
			return i
		}
	}
	return -1
}

// Intersect 返回两个集合相交的部分。
func Intersect[T comparable](list1 []T, list2 []T) []T {
	seen := make(map[T]bool, len(list1))

	for _, item := range list1 {
		seen[item] = true
	}

	result := make([]T, 0, len(list2))

	for _, item := range list2 {
		if seen[item] {
			result = append(result, item)
		}
	}

	return result
}

// Last 返回集合中最后一个元素，如果为空则返回错误。
func Last[T any](collection []T) (T, error) {
	length := len(collection)

	if length == 0 {
		var t T
		return t, fmt.Errorf("last: cannot extract the last element of an empty slice")
	}

	return collection[length-1], nil
}

// LastIndexOf 返回在数组中找到值的最后一次出现的索引。如果找不到该值则返回-1。
func LastIndexOf[T comparable](collection []T, element T) int {
	for i := len(collection) - 1; i >= 0; i-- {
		if collection[i] == element {
			return i
		}
	}

	return -1
}

// Max 搜索集合的最大值。当集合为空时返回零值。
func Max[T constraints.Ordered](collection []T) T {
	if len(collection) == 0 {
		return Empty[T]()
	}

	max := collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		max = Ternary(item > max, item, max)
	}

	return max
}

// MaxBy 使用给定的比较函数搜索集合的最大值。
// 如果集合的几个值等于最大值，则返回第一个这样的值。当集合为空时返回零值。
//
// 比较函数，当传入的第一个参数大于第二个参数时返回 true。
func MaxBy[T any](collection []T, comparison func(T, T) bool) T {
	if len(collection) == 0 {
		return Empty[T]()
	}

	max := collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		max = Ternary(comparison(item, max), item, max)
	}

	return max
}

// Min 搜索集合的最小值。当集合为空时返回零值。
func Min[T constraints.Ordered](collection []T) T {
	if len(collection) == 0 {
		return Empty[T]()
	}

	min := collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		min = Ternary(item < min, item, min)
	}

	return min
}

// MinBy 使用给定的比较函数搜索集合的最小值。
// 如果集合的几个值等于最小值，则返回第一个这样的值。当集合为空时返回零值。
//
// 比较函数，当传入的第一个参数小于第二个参数时返回 true。
func MinBy[T any](collection []T, comparison func(T, T) bool) T {
	if len(collection) == 0 {
		return Empty[T]()
	}

	min := collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		min = Ternary(comparison(item, min), item, min)
	}

	return min
}

// Nth 返回集合索引 “n” 处的元素。
// 如果 “n” 为负数，则返回末尾的第 n 个元素。当第 n 个超出切片边界时返回错误。
func Nth[T any, N constraints.Integer](collection []T, nth N) (T, error) {
	n := int(nth)
	length := len(collection)

	if n >= length || -n > length {
		var t T
		return t, fmt.Errorf("nth: %d out of slice bounds", n)
	}

	if n >= 0 {
		return collection[n], nil
	}
	return collection[length+n], nil
}

// None 判断集合 subset 中的全部元素是否都不存在于集合 collection 中。
// 如果 subset 为空，则返回 true。
func None[T comparable](collection []T, subset []T) bool {
	contain := make(map[T]bool, len(collection))

	for _, item := range collection {
		contain[item] = true
	}

	for _, item := range subset {
		if contain[item] {
			return false
		}
	}

	return true
}

// NoneBy 如果 predicate 对集合中的任何元素都不返回 true 或集合为空，则返回 true。
func NoneBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, v := range collection {
		if predicate(v) {
			return false
		}
	}

	return true
}

// Repeat 创建一个长度为 count 所有元素为 initial 的切片。
func Repeat[T Clonable[T]](count int, initial T) []T {
	result := make([]T, 0, count)

	for i := 0; i < count; i++ {
		result = append(result, initial.Clone())
	}

	return result
}

// RepeatBy 创建一个长度为 count 元素为 predicate 函数提供的切片。
func RepeatBy[T any](count int, predicate func(int) T) []T {
	result := make([]T, 0, count)

	for i := 0; i < count; i++ {
		result = append(result, predicate(i))
	}

	return result
}

// Replace 将集合中 n 个目标元素进行替换。
func Replace[T comparable](collection []T, old T, new T, n int) []T {
	result := make([]T, len(collection))
	copy(result, collection)

	for i := range result {
		if result[i] == old && n != 0 {
			result[i] = new
			n--
		}
	}

	return result
}

// ReplaceAll 将集合中所有目标元素进行替换。
func ReplaceAll[T comparable](collection []T, old T, new T) []T {
	return Replace(collection, old, new, -1)
}

// Reverse 对切片中的元素进行翻转。
func Reverse[T any](collection []T) []T {
	length := len(collection)
	half := length / 2

	for i := 0; i < half; i++ {
		j := length - i - 1
		collection[i], collection[j] = collection[j], collection[i]
	}

	return collection
}

// Sample 返回集合中的随机一个元素。
func Sample[T any](collection []T) T {
	size := len(collection)
	if size == 0 {
		return Empty[T]()
	}

	return collection[rand.Intn(size)]
}

// Samples 返回集合中 n 个随机唯一元素。
func Samples[T any](collection []T, count int) []T {
	size := len(collection)

	copy := append([]T{}, collection...)

	results := []T{}

	for i := 0; i < size && i < count; i++ {
		copyLength := size - i

		index := rand.Intn(size - i)
		results = append(results, copy[index])

		// 删除元素。与最后一个元素交换并删除它更快。
		copy[index] = copy[copyLength-1]
		copy = copy[:copyLength-1]
	}

	return results
}

// Slice 对集合截取切片，能够处理数组越界的问题而不 panic。
func Slice[T any](collection []T, start int, end int) []T {
	size := len(collection)

	if start >= end {
		return []T{}
	}

	start = If(start > size, size).ElseIf(start < 0, 0).Else(start)
	end = If(end > size, size).ElseIf(end < 0, 0).Else(end)

	return collection[start:end]
}

// SliceToMap 根据 transform 将数组转换为键值对。
func SliceToMap[T any, K comparable, V any](collection []T, transform func(T) (K, V)) map[K]V {
	result := make(map[K]V, len(collection))

	for _, item := range collection {
		k, v := transform(item)
		result[k] = v
	}

	return result
}

// Some 判断集合 subset 中的是否至少有一个元素存在于另一个集合 collection 中。
// 如果 subset 为空，则返回 false。
func Some[T comparable](collection []T, subset []T) bool {
	contain := make(map[T]bool, len(collection))

	for _, item := range collection {
		contain[item] = true
	}

	for _, item := range subset {
		if contain[item] {
			return true
		}
	}

	return false
}

// SomeBy 如果 predicate 为集合中的至少一个元素返回 true 则返回 true。
// 如果集合为空则返回 false。
func SomeBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, v := range collection {
		if predicate(v) {
			return true
		}
	}

	return false
}

// Unique 对集合进行去重，不改变所有元素在原集合中的顺序。
func Unique[T comparable](collection []T) []T {
	seen := make(map[T]bool, len(collection))
	result := make([]T, 0, len(collection))

	for _, item := range collection {
		if seen[item] {
			continue
		}

		seen[item] = true
		result = append(result, item)
	}

	return result
}

// UniqueBy 根据提供的 iteratee 函数对集合进行去重，不改变元素在原集合中的顺序。
func UniqueBy[T any, U comparable](collection []T, iteratee func(T) U) []T {
	seen := make(map[U]bool, len(collection))
	result := make([]T, 0, len(collection))

	for _, item := range collection {
		key := iteratee(item)

		if seen[key] {
			continue
		}

		seen[key] = true
		result = append(result, item)
	}

	return result
}

// Union 返回给定集合中的所有不同元素。
// 结果返回不会相对改变元素的顺序。
func Union[T comparable](lists ...[]T) []T {
	seen := make(map[T]bool)
	result := []T{}

	for _, list := range lists {
		for _, item := range list {
			if !seen[item] {
				result = append(result, item)
			}
			seen[item] = true
		}
	}

	return result
}

// Without 返回不包括所有给定值的切片。
func Without[T comparable](collection []T, exclude ...T) []T {
	seen := make(map[T]bool, len(exclude))

	for _, item := range exclude {
		seen[item] = true
	}

	result := make([]T, 0, len(collection))

	for _, item := range collection {
		if !seen[item] {
			result = append(result, item)
		}
	}

	return result
}

func WithoutEmpty[T comparable](collection []T) []T {
	var empty T

	result := make([]T, 0, len(collection))
	for _, item := range collection {
		if item != empty {
			result = append(result, item)
		}
	}

	return result
}
