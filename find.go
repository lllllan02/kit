package kit

import (
	"fmt"
	"math/rand"

	"golang.org/x/exp/constraints"
)

// IndexOf 返回在数组中找到值的第一次出现的索引。如果找不到该值则返回-1。
func IndexOf[T comparable](collection []T, element T) int {
	for i, item := range collection {
		if item == element {
			return i
		}
	}
	return -1
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

// Find 根据 predicate 在切片中搜索元素。如果找到元素，则返回元素和 true。
func Find[T any](collection []T, predicate func(T) bool) (T, bool) {
	result, _, exist := FindIndexOf(collection, predicate)
	return result, exist
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

// FindKey 返回第一个值匹配的键。
func FindKey[K comparable, V comparable](object map[K]V, value V) (K, bool) {
	for k, v := range object {
		if v == value {
			return k, true
		}
	}

	return Empty[K](), false
}

// FindKeyBy 返回第一个与 predicate 返回值匹配的键。
func FindKeyBy[K comparable, V any](object map[K]V, predicate func(key K, value V) bool) (K, bool) {
	for k, v := range object {
		if predicate(k, v) {
			return k, true
		}
	}

	return Empty[K](), false
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

// Last 返回集合中最后一个元素，如果为空则返回错误。
func Last[T any](collection []T) (T, error) {
	length := len(collection)

	if length == 0 {
		var t T
		return t, fmt.Errorf("last: cannot extract the last element of an empty slice")
	}

	return collection[length-1], nil
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
