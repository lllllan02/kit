package kit

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

type Clonable[T any] interface {
	Clone() T
}
