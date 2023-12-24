package kit

// Keys 返回 map 中所有键的切片。
func Keys[K comparable, V any](in map[K]V) []K {
	keys := make([]K, 0, len(in))

	for k := range in {
		keys = append(keys, k)
	}

	return keys
}

// Values 返回 map 中所有值的切片。
func Values[K comparable, V any](in map[K]V) []V {
	vals := make([]V, 0, len(in))

	for _, v := range in {
		vals = append(vals, v)
	}

	return vals
}

// ValueOr 返回给定键的值，如果键不存在则返回回退值。
func ValueOr[K comparable, V any](in map[K]V, key K, fallback V) V {
	if v, ok := in[key]; ok {
		return v
	}

	return fallback
}

// PickBy 使用 predicate 函数对映射进行过滤，保留返回值为 true 的键值对。
func PickBy[K comparable, V any](in map[K]V, predicate func(K, V) bool) map[K]V {
	res := make(map[K]V)

	for k, v := range in {
		if predicate(k, v) {
			res[k] = v
		}
	}

	return res
}

// PickByKeys 使用键数组对映射进行过滤，保留数组中存在的键。
func PickByKeys[K comparable, V any](in map[K]V, keys []K) map[K]V {
	res := make(map[K]V)

	for _, k := range keys {
		if v, ok := in[k]; ok {
			res[k] = v
		}
	}

	return res
}

// PickByValues 使用值数组对映射进行过滤，保留数组中存在的值。
func PickByValues[K comparable, V comparable](in map[K]V, values []V) map[K]V {
	seen := make(map[V]bool, len(values))

	for _, v := range values {
		seen[v] = true
	}

	res := make(map[K]V)

	for k, v := range in {
		if seen[v] {
			res[k] = v
		}
	}

	return res
}

// OmitBy 使用 predicate 函数对映射进行过滤，删除返回值为 true 的键值对。
func OmitBy[K comparable, V any](in map[K]V, predicate func(K, V) bool) map[K]V {
	res := make(map[K]V)

	for k, v := range in {
		if !predicate(k, v) {
			res[k] = v
		}
	}

	return res
}

// OmitByKeys 使用键数组对映射进行过滤，删除数组中存在的键。
func OmitByKeys[K comparable, V any](in map[K]V, keys []K) map[K]V {
	seen := make(map[K]bool, len(keys))

	for _, k := range keys {
		seen[k] = true
	}

	res := make(map[K]V)

	for k, v := range in {
		if !seen[k] {
			res[k] = v
		}
	}

	return res
}

// OmitByValues 使用值数组对映射进行过滤，删除数组中存在的值。
func OmitByValues[K comparable, V comparable](in map[K]V, values []V) map[K]V {
	seen := make(map[V]bool, len(values))

	for _, v := range values {
		seen[v] = true
	}

	res := make(map[K]V)

	for k, v := range in {
		if !seen[v] {
			res[k] = v
		}
	}

	return res
}

// Entries 将 map 转换为包含键值对的数组。
func Entries[K comparable, V any](in map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(in))

	for k, v := range in {
		entries = append(entries, Entry[K, V]{
			Key:   k,
			Value: v,
		})
	}

	return entries
}

// ToPairs 将 map 转换为包含键值对的数组，等价于 Entries。
func ToPairs[K comparable, V any](in map[K]V) []Entry[K, V] {
	return Entries(in)
}

// FromEntries 将包含键值对的数组转化为 map。
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	res := make(map[K]V, len(entries))

	for _, entry := range entries {
		res[entry.Key] = entry.Value
	}

	return res
}

// FromPairs 将包含键值对的数组转化为 map，等价于 FromEntries。
func FromPairs[K comparable, V any](entries []Entry[K, V]) map[K]V {
	return FromEntries(entries)
}

// Invert 反转映射中键值关系。如果映射中出现重复的值，后遍历的值则会覆盖先前的值。
func Invert[K comparable, V comparable](in map[K]V) map[V]K {
	res := make(map[V]K, len(in))

	for k, v := range in {
		res[v] = k
	}

	return res
}

// Assign 从左到右合并多个映射关系。
func Assign[K comparable, V any](maps ...map[K]V) map[K]V {
	res := make(map[K]V)

	for _, m := range maps {
		for k, v := range m {
			res[k] = v
		}
	}

	return res
}

// MapToSlice 根据提供的 iteratee 函数将 map 转换为数组。
func MapToSlice[K comparable, V any, R any](in map[K]V, iteratee func(K, V) R) []R {
	res := make([]R, 0, len(in))

	for k, v := range in {
		res = append(res, iteratee(k, v))
	}

	return res
}
