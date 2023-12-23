package kit

// Contains 判断元素是否在集合中。
func Contains[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}
	return false
}

// ContainsBy 根据 predicate 判断元素是否在集合中。
func ContainsBy[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}
	return false
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
