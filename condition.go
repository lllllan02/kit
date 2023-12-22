package kit

// Ternary 三元运算
func Ternary[T any](condition bool, ifOutput T, elseOutput T) T {
	if condition {
		return ifOutput
	}

	return elseOutput
}

// If·ElseIf·Else
type ifElse[T any] struct {
	result T
	done   bool
}

func If[T any](condition bool, result T) *ifElse[T] {
	if condition {
		return &ifElse[T]{result, true}
	}

	var t T
	return &ifElse[T]{t, false}
}

func (i *ifElse[T]) ElseIf(condition bool, result T) *ifElse[T] {
	if !i.done && condition {
		i.result = result
		i.done = true
	}

	return i
}

func (i *ifElse[T]) Else(result T) T {
	if i.done {
		return i.result
	}

	return result
}

// Switch·Case·Default
type switchCase[T comparable, R any] struct {
	predicate T
	result    R
	done      bool
}

func Switch[T comparable, R any](predicate T) *switchCase[T, R] {
	var result R

	return &switchCase[T, R]{
		predicate,
		result,
		false,
	}
}

func (s *switchCase[T, R]) Case(val T, result R) *switchCase[T, R] {
	if !s.done && s.predicate == val {
		s.result = result
		s.done = true
	}

	return s
}

func (s *switchCase[T, R]) Default(result R) R {
	if !s.done {
		s.result = result
	}

	return s.result
}
