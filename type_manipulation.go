package kit

// Empty 返回一个零值。
func Empty[T any]() T {
	var zero T
	return zero
}
