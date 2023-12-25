package kit

import (
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/exp/constraints"
)

// AppendIfAbsent 当元素不存在于集合中时将元素添加到集合末尾。
func AppendIfAbsent[T comparable](slice []T, item T) []T {
	if !Contain(slice, item) {
		slice = append(slice, item)
	}
	return slice
}

// Contain 判断元素是否在集合中。
func Contain[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

// ContainBy 根据 predicate 判断元素是否在集合中。
func ContainBy[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}
	return false
}

// Concat 合并所有数组。
func Concat[T any](slice []T, slices ...[]T) []T {
	result := append([]T{}, slice...)

	for _, item := range slices {
		result = append(result, item...)
	}

	return result
}

// Compact 去除数组中的零值元素。
func Compact[T comparable](slice []T) []T {
	var zero T

	result := make([]T, 0, len(slice))

	for _, item := range slice {
		if item != zero {
			result = append(result, item)
		}
	}

	return result
}

// Count 统计集合中 value 出现的次数。
func Count[T comparable](slice []T, value T) int {
	var count int

	for _, item := range slice {
		if item == value {
			count++
		}
	}

	return count
}

// CountBy 统计数组中满足 predicate 函数的元素个数。
func CountBy[T any](slice []T, predicate func(int, T) bool) int {
	var count int

	for i, item := range slice {
		if predicate(i, item) {
			count++
		}
	}

	return count
}

// CountValues 统计集合中元素的出现次数。
func CountValues[T comparable](slice []T) map[T]int {
	result := make(map[T]int)

	for _, item := range slice {
		result[item]++
	}

	return result
}

// DeleteAt 删除数组的指定区间。
func DeleteAt[T any](slice []T, start int, end ...int) []T {
	size := len(slice)

	if start < 0 || start >= size {
		return slice
	}

	if len(end) > 0 {
		end := end[0]

		if end <= start {
			return slice
		}

		end = Ternary(end > size, size, end)
		return append(slice[:start], slice[end:]...)
	}

	if start == size-1 {
		return slice[:start]
	}
	return append(slice[:start], slice[start+1:]...)
}

// Difference 返回两个集合之间的差异。
// 第一个值是 slice 中不存在的元素的集合。
// 第二个值是 compare 中不存在的元素的集合。
func Difference[T comparable](slice []T, compare []T) ([]T, []T) {
	left := []T{}
	right := []T{}

	leftSeen := make(map[T]bool, len(slice))
	rightSeen := make(map[T]bool, len(compare))

	for _, item := range slice {
		leftSeen[item] = true
	}

	for _, item := range compare {
		rightSeen[item] = true
		if !leftSeen[item] {
			right = append(right, item)
		}
	}

	for _, item := range slice {
		if !rightSeen[item] {
			left = append(left, item)
		}
	}

	return left, right
}

// Equal 比较两个数组是否相等，即拥有相同长度并且每个位置上的元素相同。
func Equal[T comparable](slice1 []T, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

// Every 判断集合 subset 中的全部元素是否都存在于另一个集合 slice 中。
// 如果 subset 为空，则返回 true。
func Every[T comparable](slice []T, subset []T) bool {
	contain := make(map[T]bool, len(slice))

	for _, item := range slice {
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
func EveryBy[T any](slice []T, predicate func(item T) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}

	return true
}

// Fill	使用 initial 填充集合。
func Fill[T Clonable[T]](slice []T, initial T) []T {
	result := make([]T, 0, len(slice))

	for range slice {
		result = append(result, initial.Clone())
	}

	return result
}

// Filter 使用 predicate 函数对集合进行过滤，保留返回值为 true 的所有元素。
func Filter[T any](slice []T, predicate func(int, T) bool) []T {
	result := make([]T, 0, len(slice))

	for i, item := range slice {
		if predicate(i, item) {
			result = append(result, item)
		}
	}

	return result
}

// Find 根据 predicate 在切片中搜索元素。如果找到元素，则返回元素和 true。
func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
	result, _, exist := FindIndexOf(slice, predicate)
	return result, exist
}

// FindDuplicates 返回集合中重复出现的元素去重后的切片。
// 结果值的顺序由它们在集合中出现的顺序决定。
func FindDuplicates[T comparable](slice []T) []T {
	isDupl := make(map[T]bool, len(slice))

	for _, item := range slice {
		_, ok := isDupl[item]
		isDupl[item] = ok
	}

	result := make([]T, 0, len(slice)-len(isDupl))

	for _, item := range slice {
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
func FindDuplicatesBy[T any, U comparable](slice []T, iteratee func(T) U) []T {
	isDupl := make(map[U]bool, len(slice))

	for _, item := range slice {
		key := iteratee(item)

		_, ok := isDupl[key]
		isDupl[key] = ok
	}

	result := make([]T, 0, len(slice)-len(isDupl))

	for _, item := range slice {
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
func FindIndexOf[T any](slice []T, predicate func(T) bool) (T, int, bool) {
	for i, item := range slice {
		if predicate(item) {
			return item, i, true
		}
	}

	return Empty[T](), -1, false
}

// FindLastIndexOf 根据 predicate 搜索切片中的最后一个元素并返回索引和 true。
// 如果未找到该元素，则返回 -1 和 false。
func FindLastIndexOf[T any](slice []T, predicate func(T) bool) (T, int, bool) {
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(slice[i]) {
			return slice[i], i, true
		}
	}

	return Empty[T](), -1, false
}

// FindOrElse 根据 predicate 搜索切片中的元素。如果找到则返回元素，否则返回给定的回退值。
func FindOrElse[T any](slice []T, fallback T, predicate func(T) bool) T {
	for _, item := range slice {
		if predicate(item) {
			return item
		}
	}

	return fallback
}

// FindUniques 返回包含集合所有唯一元素的切片。
// 结果值的顺序由它们在集合中出现的顺序决定。
func FindUniques[T comparable](slice []T) []T {
	isDupl := make(map[T]bool, len(slice))

	for _, item := range slice {
		_, ok := isDupl[item]
		isDupl[item] = ok // 出现一次为 false，重复出现为 true
	}

	result := make([]T, 0, len(slice)-len(isDupl))

	for _, item := range slice {
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
func ForEach[T any](slice []T, iteratee func(int, T)) {
	for i, item := range slice {
		iteratee(i, item)
	}
}

// GroupBy 根据提供的 iteratee 函数对集合进行分组。
func GroupBy[T any, U comparable](slice []T, iteratee func(T) U) map[U][]T {
	result := make(map[U][]T)

	for _, item := range slice {
		key := iteratee(item)
		result[key] = append(result[key], item)
	}

	return result
}

// IndexOf 返回在数组中找到值的第一次出现的索引。如果找不到该值则返回-1。
func IndexOf[T comparable](slice []T, element T) int {
	for i, item := range slice {
		if item == element {
			return i
		}
	}
	return -1
}

// InsertAt 将目标元素或集合插入到另一个集合指定下标处。
// 如果下标小于零，默认插入到集合前方。
// 如果下标大于集合的长度，默认插入到集合后方。
func InsertAt[T any](slice []T, index int, value ...T) []T {
	if len(value) == 0 {
		return slice
	}

	if index < 0 {
		return append(value, slice...)
	}

	if index > len(slice) {
		return append(slice, value...)
	}

	return append(slice[:index], append(value, slice[index:]...)...)
}

// Intersect 返回所有集合中都存在的元素切片。
func Intersect[T comparable](slices ...[]T) []T {
	size := len(slices)
	if size == 0 {
		return []T{}
	}
	if size == 1 {
		return Unique(slices[0])
	}

	reducer := func(slice1 []T, slice2 []T) []T {
		seen := make(map[T]bool, len(slice1))
		result := make([]T, 0, len(slice1))

		for _, item := range slice1 {
			seen[item] = true
		}

		for _, item := range slice2 {
			if seen[item] {
				seen[item] = false
				result = append(result, item)
			}
		}

		return result
	}

	result := reducer(slices[0], slices[1])
	for i := 2; i < size; i++ {
		result = reducer(result, slices[i])
	}
	return result
}

// IsAscending 判断数组是否从小到大有序。
func IsAscending[T constraints.Ordered](slice []T) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
		}
	}
	return true
}

// IsDescending 判断数组是否从大到小有序。
func IsDescending[T constraints.Ordered](slice []T) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] < slice[i] {
			return false
		}
	}
	return true
}

// IsSorted 判断数组是否有序，从小到大或从大到小。
func IsSorted[T constraints.Ordered](slice []T) bool {
	return IsAscending(slice) || IsDescending(slice)
}

// Join 使用 sepratator 连接数组中所有元素返回一个字符串。
func Join[T any](slice []T, separator string) string {
	str := Map(slice, func(i int, item T) string {
		return fmt.Sprint(item)
	})
	return strings.Join(str, separator)
}

// Last 返回集合中最后一个元素，如果为空则返回错误。
func Last[T any](slice []T) (T, error) {
	length := len(slice)

	if length == 0 {
		var t T
		return t, fmt.Errorf("last: cannot extract the last element of an empty slice")
	}

	return slice[length-1], nil
}

// LastIndexOf 返回在数组中找到值的最后一次出现的索引。如果找不到该值则返回-1。
func LastIndexOf[T comparable](slice []T, element T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == element {
			return i
		}
	}

	return -1
}

// Map 将数组转换成另一个类型的数组。
func Map[T any, U any](slice []T, iteratee func(int, T) U) []U {
	result := make([]U, 0, len(slice))

	for i, item := range slice {
		result = append(result, iteratee(i, item))
	}

	return result
}

// Merge 合并所有数组。
func Merge[T any](slices ...[]T) []T {
	result := []T{}

	for _, slice := range slices {
		result = append(result, slice...)
	}

	return result
}

// Nth 返回集合索引 “n” 处的元素。
// 如果 “n” 为负数，则返回末尾的第 n 个元素。当第 n 个超出切片边界时返回错误。
func Nth[T any, N constraints.Integer](slice []T, nth N) (T, error) {
	n := int(nth)
	length := len(slice)

	if n >= length || -n > length {
		var t T
		return t, fmt.Errorf("nth: %d out of slice bounds", n)
	}

	if n >= 0 {
		return slice[n], nil
	}
	return slice[length+n], nil
}

// None 判断集合 subset 中的全部元素是否都不存在于集合 slice 中。
// 如果 subset 为空，则返回 true。
func None[T comparable](slice []T, subset []T) bool {
	contain := make(map[T]bool, len(slice))

	for _, item := range slice {
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
func NoneBy[T any](slice []T, predicate func(item T) bool) bool {
	for _, v := range slice {
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
func Replace[T comparable](slice []T, old T, new T, n int) []T {
	result := make([]T, len(slice))
	copy(result, slice)

	for i := range result {
		if result[i] == old && n != 0 {
			result[i] = new
			n--
		}
	}

	return result
}

// ReplaceAll 将集合中所有目标元素进行替换。
func ReplaceAll[T comparable](slice []T, old T, new T) []T {
	return Replace(slice, old, new, -1)
}

// Reverse 对切片中的元素进行翻转。
func Reverse[T any](slice []T) []T {
	length := len(slice)
	half := length / 2

	for i := 0; i < half; i++ {
		j := length - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}

	return slice
}

// Sample 返回集合中的随机一个元素。
func Sample[T any](slice []T) T {
	size := len(slice)
	if size == 0 {
		return Empty[T]()
	}

	return slice[rand.Intn(size)]
}

// Samples 返回集合中 n 个随机唯一元素。
func Samples[T any](slice []T, count int) []T {
	size := len(slice)

	copy := append([]T{}, slice...)

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
func Slice[T any](slice []T, start int, end int) []T {
	size := len(slice)

	if start >= end {
		return []T{}
	}

	start = If(start > size, size).ElseIf(start < 0, 0).Else(start)
	end = If(end > size, size).ElseIf(end < 0, 0).Else(end)

	return slice[start:end]
}

// SliceToMap 根据 transform 将数组转换为键值对。
func SliceToMap[T any, K comparable, V any](slice []T, transform func(T) (K, V)) map[K]V {
	result := make(map[K]V, len(slice))

	for _, item := range slice {
		k, v := transform(item)
		result[k] = v
	}

	return result
}

// Some 判断集合 subset 中的是否至少有一个元素存在于另一个集合 slice 中。
// 如果 subset 为空，则返回 false。
func Some[T comparable](slice []T, subset []T) bool {
	contain := make(map[T]bool, len(slice))

	for _, item := range slice {
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
func SomeBy[T any](slice []T, predicate func(item T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}

	return false
}

// Shuffle 打乱集合中的元素顺序。
func Shuffle[T any](slice []T) []T {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}

// UpdateAt 更新数组指定位置的元素。如果下标越界，则返回原数组。
func UpdateAt[T any](slice []T, index int, value T) []T {
	size := len(slice)

	if index < 0 || index >= size {
		return slice
	}
	return append(slice[:index], append([]T{value}, slice[index+1:]...)...)
}

// Unique 对集合进行去重，不改变所有元素在原集合中的顺序。
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool, len(slice))
	result := make([]T, 0, len(slice))

	for _, item := range slice {
		if seen[item] {
			continue
		}

		seen[item] = true
		result = append(result, item)
	}

	return result
}

// UniqueBy 根据提供的 iteratee 函数对集合进行去重，不改变元素在原集合中的顺序。
func UniqueBy[T any, U comparable](slice []T, iteratee func(T) U) []T {
	seen := make(map[U]bool, len(slice))
	result := make([]T, 0, len(slice))

	for _, item := range slice {
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
func Union[T comparable](slices ...[]T) []T {
	seen := make(map[T]bool)
	result := []T{}

	for _, slice := range slices {
		for _, item := range slice {
			if seen[item] {
				continue
			}
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// UnionBy 根据 predicate 函数返回所有集合中的不同元素。
func UnionBy[T any, U comparable](predicate func(T) U, slices ...[]T) []T {
	seen := make(map[U]bool)
	result := []T{}

	for _, slice := range slices {
		for _, item := range slice {
			key := predicate(item)
			if seen[key] {
				continue
			}
			seen[key] = true
			result = append(result, item)
		}
	}

	return result
}

// Without 返回不包括所有给定值的切片。
func Without[T comparable](slice []T, exclude ...T) []T {
	seen := make(map[T]bool, len(exclude))

	for _, item := range exclude {
		seen[item] = true
	}

	result := make([]T, 0, len(slice))

	for _, item := range slice {
		if !seen[item] {
			result = append(result, item)
		}
	}

	return result
}

func WithoutEmpty[T comparable](slice []T) []T {
	var empty T

	result := make([]T, 0, len(slice))
	for _, item := range slice {
		if item != empty {
			result = append(result, item)
		}
	}

	return result
}
