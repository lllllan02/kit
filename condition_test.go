package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Ternary(true, "if", "else")
	is.Equal(result1, "if")

	result2 := Ternary(false, "if", "else")
	is.Equal(result2, "else")
}

func TestIfElse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := If(true, "if").ElseIf(false, "else if").Else("else")
	is.Equal(result1, "if")

	result2 := If(true, "if").ElseIf(true, "else if").Else("else")
	is.Equal(result2, "if")

	result3 := If(false, "if").ElseIf(true, "else if").Else("else")
	is.Equal(result3, "else if")

	result4 := If(false, "if").ElseIf(false, "else if").Else("else")
	is.Equal(result4, "else")
}

func TestSwitchCase(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Switch[int, int](42).Case(42, 1).Case(1, 2).Default(3)
	is.Equal(result1, 1)

	result2 := Switch[int, int](42).Case(42, 1).Case(42, 2).Default(3)
	is.Equal(result2, 1)

	result3 := Switch[int, int](42).Case(1, 1).Case(42, 2).Default(3)
	is.Equal(result3, 2)

	result4 := Switch[int, int](42).Case(1, 1).Case(1, 2).Default(3)
	is.Equal(result4, 3)
}
