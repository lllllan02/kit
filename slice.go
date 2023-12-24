package kit

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

// ForEach 遍历集合中的所有元素并执行 iteratee 函数。
func ForEach[T any](collection []T, iteratee func(int, T)) {
	for i, item := range collection {
		iteratee(i, item)
	}
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

// GroupBy 根据提供的 iteratee 函数对集合进行分组。
func GroupBy[T any, U comparable](collection []T, iteratee func(T) U) map[U][]T {
	result := make(map[U][]T)

	for _, item := range collection {
		key := iteratee(item)
		result[key] = append(result[key], item)
	}

	return result
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

// Fill	使用 initial 填充集合。
func Fill[T Clonable[T]](collection []T, initial T) []T {
	result := make([]T, 0, len(collection))

	for range collection {
		result = append(result, initial.Clone())
	}

	return result
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

// SliceToMap 根据 transform 将数组转换为键值对。
func SliceToMap[T any, K comparable, V any](collection []T, transform func(T) (K, V)) map[K]V {
	result := make(map[K]V, len(collection))

	for _, item := range collection {
		k, v := transform(item)
		result[k] = v
	}

	return result
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
