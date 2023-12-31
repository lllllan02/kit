package kit

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Keys(map[string]int{"foo": 1, "bar": 2})
	sort.Strings(r1)

	is.Equal(r1, []string{"bar", "foo"})
}

func TestValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Values(map[string]int{"foo": 1, "bar": 2})
	sort.Ints(r1)

	is.Equal(r1, []int{1, 2})
}

func TestValueOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := ValueOr(map[string]int{"foo": 1}, "bar", 2)
	is.Equal(r1, 2)

	r2 := ValueOr(map[string]int{"foo": 1}, "foo", 2)
	is.Equal(r2, 1)
}

func TestPickBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := PickBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(key string, value int) bool {
		return value%2 == 1
	})

	is.Equal(r1, map[string]int{"foo": 1, "baz": 3})
}

func TestPickByKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := PickByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz"})

	is.Equal(r1, map[string]int{"foo": 1, "baz": 3})
}

func TestPickByValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := PickByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{1, 3})

	is.Equal(r1, map[string]int{"foo": 1, "baz": 3})
}

func TestOmitBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := OmitBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(key string, value int) bool {
		return value%2 == 1
	})

	is.Equal(r1, map[string]int{"bar": 2})
}

func TestOmitByKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := OmitByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz"})

	is.Equal(r1, map[string]int{"bar": 2})
}

func TestOmitByValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := OmitByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{1, 3})

	is.Equal(r1, map[string]int{"bar": 2})
}

func TestEntries(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Entries(map[string]int{"foo": 1, "bar": 2})

	sort.Slice(r1, func(i, j int) bool {
		return r1[i].Value < r1[j].Value
	})
	is.EqualValues(r1, []Entry[string, int]{
		{
			Key:   "foo",
			Value: 1,
		},
		{
			Key:   "bar",
			Value: 2,
		},
	})
}

func TestToPairs(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := ToPairs(map[string]int{"baz": 3, "qux": 4})

	sort.Slice(r1, func(i, j int) bool {
		return r1[i].Value < r1[j].Value
	})
	is.EqualValues(r1, []Entry[string, int]{
		{
			Key:   "baz",
			Value: 3,
		},
		{
			Key:   "qux",
			Value: 4,
		},
	})
}

func TestFromEntries(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FromEntries([]Entry[string, int]{
		{
			Key:   "foo",
			Value: 1,
		},
		{
			Key:   "bar",
			Value: 2,
		},
	})

	is.Len(r1, 2)
	is.Equal(r1["foo"], 1)
	is.Equal(r1["bar"], 2)
}

func TestFromPairs(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FromPairs([]Entry[string, int]{
		{
			Key:   "baz",
			Value: 3,
		},
		{
			Key:   "qux",
			Value: 4,
		},
	})

	is.Len(r1, 2)
	is.Equal(r1["baz"], 3)
	is.Equal(r1["qux"], 4)
}

func TestInvert(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Invert(map[string]int{"a": 1, "b": 2})
	r2 := Invert(map[string]int{"a": 1, "b": 2, "c": 1})

	is.Len(r1, 2)
	is.EqualValues(map[int]string{1: "a", 2: "b"}, r1)
	is.Len(r2, 2)
}

func TestAssign(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Assign(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4})

	is.Len(result1, 3)
	is.Equal(result1, map[string]int{"a": 1, "b": 3, "c": 4})
}

func TestMapToSlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MapToSlice(map[int]int{1: 5, 2: 6, 3: 7, 4: 8}, func(k int, v int) string {
		return fmt.Sprintf("%d_%d", k, v)
	})
	result2 := MapToSlice(map[int]int{1: 5, 2: 6, 3: 7, 4: 8}, func(k int, _ int) string {
		return strconv.FormatInt(int64(k), 10)
	})

	is.Equal(len(result1), 4)
	is.Equal(len(result2), 4)
	is.ElementsMatch(result1, []string{"1_5", "2_6", "3_7", "4_8"})
	is.ElementsMatch(result2, []string{"1", "2", "3", "4"})
}

func TestFindKey(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := FindKey(map[string]int{"foo": 1, "bar": 2, "baz": 3}, 2)
	is.Equal("bar", result1)
	is.True(ok1)

	result2, ok2 := FindKey(map[string]int{"foo": 1, "bar": 2, "baz": 3}, 42)
	is.Equal("", result2)
	is.False(ok2)

	type test struct {
		foobar string
	}

	result3, ok3 := FindKey(map[string]test{"foo": {"foo"}, "bar": {"bar"}, "baz": {"baz"}}, test{"foo"})
	is.Equal("foo", result3)
	is.True(ok3)

	result4, ok4 := FindKey(map[string]test{"foo": {"foo"}, "bar": {"bar"}, "baz": {"baz"}}, test{"hello world"})
	is.Equal("", result4)
	is.False(ok4)
}

func TestFindKeyBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := FindKeyBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return k == "foo"
	})
	is.Equal("foo", result1)
	is.True(ok1)

	result2, ok2 := FindKeyBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return false
	})
	is.Equal("", result2)
	is.False(ok2)
}
