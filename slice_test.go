package kit

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestContain(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Contain([]int{0, 1, 2, 3, 4, 5}, 5)
	result2 := Contain([]int{0, 1, 2, 3, 4, 5}, 6)

	is.Equal(result1, true)
	is.Equal(result2, false)
}

func TestContainBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type a struct {
		A int
		B string
	}

	a1 := []a{{A: 1, B: "1"}, {A: 2, B: "2"}, {A: 3, B: "3"}}
	result1 := ContainBy(a1, func(t a) bool { return t.A == 1 && t.B == "2" })
	result2 := ContainBy(a1, func(t a) bool { return t.A == 2 && t.B == "2" })

	a2 := []string{"aaa", "bbb", "ccc"}
	result3 := ContainBy(a2, func(t string) bool { return t == "ccc" })
	result4 := ContainBy(a2, func(t string) bool { return t == "ddd" })

	is.Equal(result1, false)
	is.Equal(result2, true)
	is.Equal(result3, true)
	is.Equal(result4, false)
}

func TestEvery(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Every([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := Every([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := Every([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	result4 := Every([]int{0, 1, 2, 3, 4, 5}, []int{})

	is.True(result1)
	is.False(result2)
	is.False(result3)
	is.True(result4)
}

func TestEveryBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := EveryBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 5
	})

	is.True(result1)

	result2 := EveryBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 3
	})

	is.False(result2)

	result3 := EveryBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 0
	})

	is.False(result3)

	result4 := EveryBy([]int{}, func(x int) bool {
		return x < 5
	})

	is.True(result4)
}

func TestSome(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Some([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := Some([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := Some([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	result4 := Some([]int{0, 1, 2, 3, 4, 5}, []int{})

	is.True(result1)
	is.True(result2)
	is.False(result3)
	is.False(result4)
}

func TestSomeBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := SomeBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 5
	})

	is.True(result1)

	result2 := SomeBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 3
	})

	is.True(result2)

	result3 := SomeBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 0
	})

	is.False(result3)

	result4 := SomeBy([]int{}, func(x int) bool {
		return x < 5
	})

	is.False(result4)
}

func TestNone(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := None([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := None([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := None([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	result4 := None([]int{0, 1, 2, 3, 4, 5}, []int{})

	is.False(result1)
	is.False(result2)
	is.True(result3)
	is.True(result4)
}

func TestNoneBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := NoneBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 5
	})

	is.False(result1)

	result2 := NoneBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 3
	})

	is.False(result2)

	result3 := NoneBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 0
	})

	is.True(result3)

	result4 := NoneBy([]int{}, func(x int) bool {
		return x < 5
	})

	is.True(result4)
}

func TestIntersect(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	s1 := []int{1, 2, 2, 3}
	s2 := []int{1, 2, 3, 4}
	s3 := []int{0, 2, 3, 5, 6}
	s4 := []int{0, 5, 6}

	expected := [][]int{
		{2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{},
	}
	result := []any{
		Intersect(s1, s2, s3),
		Intersect(s1, s2),
		Intersect(s1),
		Intersect(s1, s4),
	}

	for i := 0; i < len(result); i++ {
		assert.Equal(expected[i], result[i])
	}
}

func TestDifference(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	left1, right1 := Difference([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 6})
	is.Equal(left1, []int{1, 3, 4, 5})
	is.Equal(right1, []int{6})

	left2, right2 := Difference([]int{1, 2, 3, 4, 5}, []int{0, 6})
	is.Equal(left2, []int{1, 2, 3, 4, 5})
	is.Equal(right2, []int{0, 6})

	left3, right3 := Difference([]int{0, 1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4, 5})
	is.Equal(left3, []int{})
	is.Equal(right3, []int{})
}

func TestUnion(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Union([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 10})
	result2 := Union([]int{0, 1, 2, 3, 4, 5}, []int{6, 7})
	result3 := Union([]int{0, 1, 2, 3, 4, 5}, []int{})
	result4 := Union([]int{0, 1, 2}, []int{0, 1, 2})
	result5 := Union([]int{}, []int{})
	is.Equal(result1, []int{0, 1, 2, 3, 4, 5, 10})
	is.Equal(result2, []int{0, 1, 2, 3, 4, 5, 6, 7})
	is.Equal(result3, []int{0, 1, 2, 3, 4, 5})
	is.Equal(result4, []int{0, 1, 2})
	is.Equal(result5, []int{})

	result11 := Union([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 10}, []int{0, 1, 11})
	result12 := Union([]int{0, 1, 2, 3, 4, 5}, []int{6, 7}, []int{8, 9})
	result13 := Union([]int{0, 1, 2, 3, 4, 5}, []int{}, []int{})
	result14 := Union([]int{0, 1, 2}, []int{0, 1, 2}, []int{0, 1, 2})
	result15 := Union([]int{}, []int{}, []int{})
	is.Equal(result11, []int{0, 1, 2, 3, 4, 5, 10, 11})
	is.Equal(result12, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	is.Equal(result13, []int{0, 1, 2, 3, 4, 5})
	is.Equal(result14, []int{0, 1, 2})
	is.Equal(result15, []int{})
}

func TestWithout(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Without([]int{0, 2, 10}, 0, 1, 2, 3, 4, 5)
	result2 := Without([]int{0, 7}, 0, 1, 2, 3, 4, 5)
	result3 := Without([]int{}, 0, 1, 2, 3, 4, 5)
	result4 := Without([]int{0, 1, 2}, 0, 1, 2)
	result5 := Without([]int{})
	is.Equal(result1, []int{10})
	is.Equal(result2, []int{7})
	is.Equal(result3, []int{})
	is.Equal(result4, []int{})
	is.Equal(result5, []int{})
}

func TestWithoutEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := WithoutEmpty([]int{0, 1, 2})
	result2 := WithoutEmpty([]int{1, 2})
	result3 := WithoutEmpty([]int{})
	is.Equal(result1, []int{1, 2})
	is.Equal(result2, []int{1, 2})
	is.Equal(result3, []int{})
}

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Filter([]int{1, 2, 3, 4}, func(_ int, x int) bool {
		return x%2 == 0
	})

	is.Equal(r1, []int{2, 4})

	r2 := Filter([]string{"", "foo", "", "bar", ""}, func(_ int, x string) bool {
		return len(x) > 0
	})

	is.Equal(r2, []string{"foo", "bar"})
}

func TestForEach(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	callParams1 := []string{}
	callParams2 := []int{}

	ForEach([]string{"a", "b", "c"}, func(i int, item string) {
		callParams1 = append(callParams1, item)
		callParams2 = append(callParams2, i)
	})

	is.ElementsMatch([]string{"a", "b", "c"}, callParams1)
	is.ElementsMatch([]int{0, 1, 2}, callParams2)
	is.IsIncreasing(callParams2)
}

func TestUniq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Unique([]int{1, 2, 2, 1})

	is.Equal(len(result1), 2)
	is.Equal(result1, []int{1, 2})
}

func TestUniqBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := UniqueBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, []int{0, 1, 2})
}

func TestGroupBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, map[int][]int{
		0: {0, 3},
		1: {1, 4},
		2: {2, 5},
	})
}

func TestReverse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Reverse([]int{0, 1, 2, 3, 4, 5})
	result2 := Reverse([]int{0, 1, 2, 3, 4, 5, 6})
	result3 := Reverse([]int{})

	is.Equal(result1, []int{5, 4, 3, 2, 1, 0})
	is.Equal(result2, []int{6, 5, 4, 3, 2, 1, 0})
	is.Equal(result3, []int{})
}

type foo struct {
	bar string
}

func (f foo) Clone() foo {
	return foo{f.bar}
}

func TestFill(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Fill([]foo{{"a"}, {"a"}}, foo{"b"})
	result2 := Fill([]foo{}, foo{"a"})

	is.Equal(result1, []foo{{"b"}, {"b"}})
	is.Equal(result2, []foo{})
}

func TestRepeat(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Repeat(2, foo{"a"})
	result2 := Repeat(0, foo{"a"})

	is.Equal(result1, []foo{{"a"}, {"a"}})
	is.Equal(result2, []foo{})
}

func TestRepeatBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	cb := func(i int) int {
		return int(math.Pow(float64(i), 2))
	}

	result1 := RepeatBy(0, cb)
	result2 := RepeatBy(2, cb)
	result3 := RepeatBy(5, cb)

	is.Equal([]int{}, result1)
	is.Equal([]int{0, 1}, result2)
	is.Equal([]int{0, 1, 4, 9, 16}, result3)
}

func TestSliceToMap(t *testing.T) {
	t.Parallel()

	type foo struct {
		baz string
		bar int
	}
	transform := func(f *foo) (string, int) {
		return f.baz, f.bar
	}
	testCases := []struct {
		in     []*foo
		expect map[string]int
	}{
		{
			in:     []*foo{{baz: "apple", bar: 1}},
			expect: map[string]int{"apple": 1},
		},
		{
			in:     []*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}},
			expect: map[string]int{"apple": 1, "banana": 2},
		},
		{
			in:     []*foo{{baz: "apple", bar: 1}, {baz: "apple", bar: 2}},
			expect: map[string]int{"apple": 2},
		},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			is := assert.New(t)
			is.Equal(SliceToMap(testCase.in, transform), testCase.expect)
		})
	}
}

func TestCount(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	count1 := Count([]int{1, 2, 1}, 1)
	count2 := Count([]int{1, 2, 1}, 3)
	count3 := Count([]int{}, 1)

	is.Equal(count1, 2)
	is.Equal(count2, 0)
	is.Equal(count3, 0)
}

func TestCountValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal(map[int]int{}, CountValues([]int{}))
	is.Equal(map[int]int{1: 1, 2: 1}, CountValues([]int{1, 2}))
	is.Equal(map[int]int{1: 1, 2: 2}, CountValues([]int{1, 2, 2}))
	is.Equal(map[string]int{"": 1, "foo": 1, "bar": 1}, CountValues([]string{"foo", "bar", ""}))
	is.Equal(map[string]int{"foo": 1, "bar": 2}, CountValues([]string{"foo", "bar", "bar"}))
}

func TestSlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 2, 3, 4}

	out1 := Slice(in, 0, 0)
	out2 := Slice(in, 0, 1)
	out3 := Slice(in, 0, 5)
	out4 := Slice(in, 0, 6)
	out5 := Slice(in, 1, 1)
	out6 := Slice(in, 1, 5)
	out7 := Slice(in, 1, 6)
	out8 := Slice(in, 4, 5)
	out9 := Slice(in, 5, 5)
	out10 := Slice(in, 6, 5)
	out11 := Slice(in, 6, 6)
	out12 := Slice(in, 1, 0)
	out13 := Slice(in, 5, 0)
	out14 := Slice(in, 6, 4)
	out15 := Slice(in, 6, 7)
	out16 := Slice(in, -10, 1)
	out17 := Slice(in, -1, 3)
	out18 := Slice(in, -10, 7)

	is.Equal([]int{}, out1)
	is.Equal([]int{0}, out2)
	is.Equal([]int{0, 1, 2, 3, 4}, out3)
	is.Equal([]int{0, 1, 2, 3, 4}, out4)
	is.Equal([]int{}, out5)
	is.Equal([]int{1, 2, 3, 4}, out6)
	is.Equal([]int{1, 2, 3, 4}, out7)
	is.Equal([]int{4}, out8)
	is.Equal([]int{}, out9)
	is.Equal([]int{}, out10)
	is.Equal([]int{}, out11)
	is.Equal([]int{}, out12)
	is.Equal([]int{}, out13)
	is.Equal([]int{}, out14)
	is.Equal([]int{}, out15)
	is.Equal([]int{0}, out16)
	is.Equal([]int{0, 1, 2}, out17)
	is.Equal([]int{0, 1, 2, 3, 4}, out18)
}

func TestReplace(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 0, 1, 2, 3, 0}

	out1 := Replace(in, 0, 42, 2)
	out2 := Replace(in, 0, 42, 1)
	out3 := Replace(in, 0, 42, 0)
	out4 := Replace(in, 0, 42, -1)
	out5 := Replace(in, 0, 42, -1)
	out6 := Replace(in, -1, 42, 2)
	out7 := Replace(in, -1, 42, 1)
	out8 := Replace(in, -1, 42, 0)
	out9 := Replace(in, -1, 42, -1)
	out10 := Replace(in, -1, 42, -1)

	is.Equal([]int{42, 1, 42, 1, 2, 3, 0}, out1)
	is.Equal([]int{42, 1, 0, 1, 2, 3, 0}, out2)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out3)
	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, out4)
	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, out5)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out6)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out7)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out8)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out9)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out10)
}

func TestReplaceAll(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 0, 1, 2, 3, 0}

	out1 := ReplaceAll(in, 0, 42)
	out2 := ReplaceAll(in, -1, 42)

	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, out1)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out2)
}

func TestIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := IndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := IndexOf([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 2)
	is.Equal(result2, -1)
}

func TestLastIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 4)
	is.Equal(result2, -1)
}

func TestFind(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	result1, ok1 := Find([]string{"a", "b", "c", "d"}, func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d"}[index], item)
		index++
		return item == "b"
	})

	result2, ok2 := Find([]string{"foobar"}, func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(ok1, true)
	is.Equal(result1, "b")
	is.Equal(ok2, false)
	is.Equal(result2, "")
}

func TestFindIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	item1, index1, ok1 := FindIndexOf([]string{"a", "b", "c", "d", "b"}, func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d", "b"}[index], item)
		index++
		return item == "b"
	})
	item2, index2, ok2 := FindIndexOf([]string{"foobar"}, func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(item1, "b")
	is.Equal(ok1, true)
	is.Equal(index1, 1)
	is.Equal(item2, "")
	is.Equal(ok2, false)
	is.Equal(index2, -1)
}

func TestFindLastIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	item1, index1, ok1 := FindLastIndexOf([]string{"a", "b", "c", "d", "b"}, func(item string) bool {
		is.Equal([]string{"b", "d", "c", "b", "a"}[index], item)
		index++
		return item == "b"
	})
	item2, index2, ok2 := FindLastIndexOf([]string{"foobar"}, func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(item1, "b")
	is.Equal(ok1, true)
	is.Equal(index1, 4)
	is.Equal(item2, "")
	is.Equal(ok2, false)
	is.Equal(index2, -1)
}

func TestFindOrElse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	result1 := FindOrElse([]string{"a", "b", "c", "d"}, "x", func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d"}[index], item)
		index++
		return item == "b"
	})
	result2 := FindOrElse([]string{"foobar"}, "x", func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(result1, "b")
	is.Equal(result2, "x")
}

func TestFindUniques(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindUniques([]int{1, 2, 3})

	is.Equal(3, len(result1))
	is.Equal([]int{1, 2, 3}, result1)

	result2 := FindUniques([]int{1, 2, 2, 3, 1, 2})

	is.Equal(1, len(result2))
	is.Equal([]int{3}, result2)

	result3 := FindUniques([]int{1, 2, 2, 1})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)

	result4 := FindUniques([]int{})

	is.Equal(0, len(result4))
	is.Equal([]int{}, result4)
}

func TestFindUniquesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindUniquesBy([]int{0, 1, 2}, func(i int) int {
		return i % 3
	})

	is.Equal(3, len(result1))
	is.Equal([]int{0, 1, 2}, result1)

	result2 := FindUniquesBy([]int{0, 1, 2, 3, 4}, func(i int) int {
		return i % 3
	})

	is.Equal(1, len(result2))
	is.Equal([]int{2}, result2)

	result3 := FindUniquesBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)

	result4 := FindUniquesBy([]int{}, func(i int) int {
		return i % 3
	})

	is.Equal(0, len(result4))
	is.Equal([]int{}, result4)
}

func TestFindDuplicates(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindDuplicates([]int{1, 2, 2, 1, 2, 3})

	is.Equal(2, len(result1))
	is.Equal([]int{1, 2}, result1)

	result2 := FindDuplicates([]int{1, 2, 3})

	is.Equal(0, len(result2))
	is.Equal([]int{}, result2)

	result3 := FindDuplicates([]int{})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)
}

func TestFindDuplicatesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindDuplicatesBy([]int{3, 4, 5, 6, 7}, func(i int) int {
		return i % 3
	})

	is.Equal(2, len(result1))
	is.Equal([]int{3, 4}, result1)

	result2 := FindDuplicatesBy([]int{0, 1, 2, 3, 4}, func(i int) int {
		return i % 5
	})

	is.Equal(0, len(result2))
	is.Equal([]int{}, result2)

	result3 := FindDuplicatesBy([]int{}, func(i int) int {
		return i % 3
	})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)
}

func TestLast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, err1 := Last([]int{1, 2, 3})
	result2, err2 := Last([]int{})

	is.Equal(result1, 3)
	is.Equal(err1, nil)
	is.Equal(result2, 0)
	is.Equal(err2, fmt.Errorf("last: cannot extract the last element of an empty slice"))
}

func TestNth(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, err1 := Nth([]int{0, 1, 2, 3}, 2)
	result2, err2 := Nth([]int{0, 1, 2, 3}, -2)
	result3, err3 := Nth([]int{0, 1, 2, 3}, 42)
	result4, err4 := Nth([]int{}, 0)
	result5, err5 := Nth([]int{42}, 0)
	result6, err6 := Nth([]int{42}, -1)

	is.Equal(result1, 2)
	is.Equal(err1, nil)
	is.Equal(result2, 2)
	is.Equal(err2, nil)
	is.Equal(result3, 0)
	is.Equal(err3, fmt.Errorf("nth: 42 out of slice bounds"))
	is.Equal(result4, 0)
	is.Equal(err4, fmt.Errorf("nth: 0 out of slice bounds"))
	is.Equal(result5, 42)
	is.Equal(err5, nil)
	is.Equal(result6, 42)
	is.Equal(err6, nil)
}

func TestSamples(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	rand.Seed(time.Now().UnixNano())

	result1 := Samples([]string{"a", "b", "c"}, 3)
	result2 := Samples([]string{}, 3)

	sort.Strings(result1)

	is.Equal(result1, []string{"a", "b", "c"})
	is.Equal(result2, []string{})
}

func TestCompact(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{}, Compact([]int{0}))
	is.Equal([]int{1, 2, 3}, Compact([]int{0, 1, 2, 3}))
	is.Equal([]string{}, Compact([]string{""}))
	is.Equal([]string{" "}, Compact([]string{" "}))
	is.Equal([]string{"a", "b", "0"}, Compact([]string{"", "a", "b", "0"}))
	is.Equal([]bool{true, true}, Compact([]bool{false, true, true}))
}

func TestConcat(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{1, 2, 3, 4, 5}, Concat([]int{1, 2, 3}, []int{4, 5}))
	is.Equal([]int{1, 2, 3, 4, 5}, Concat([]int{1, 2, 3}, []int{4}, []int{5}))
}

func TestEqual(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slice3 := []int{3, 2, 1}

	assert.Equal(true, Equal(slice1, slice2))
	assert.Equal(false, Equal(slice1, slice3))
}

func TestCountBy(t *testing.T) {
	t.Parallel()

	nums := []int{1, 2, 3, 4, 5, 6}
	evenFunc := func(i, num int) bool {
		return (num % 2) == 0
	}

	assert := assert.New(t)
	assert.Equal(3, CountBy(nums, evenFunc))
}

func TestDeleteAt(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	assert.Equal([]string{"a", "b", "c"}, DeleteAt([]string{"a", "b", "c"}, -1))
	assert.Equal([]string{"a", "b", "c"}, DeleteAt([]string{"a", "b", "c"}, 3))
	assert.Equal([]string{"b", "c"}, DeleteAt([]string{"a", "b", "c"}, 0))
	assert.Equal([]string{"a", "c"}, DeleteAt([]string{"a", "b", "c"}, 1))
	assert.Equal([]string{"a", "b"}, DeleteAt([]string{"a", "b", "c"}, 2))

	assert.Equal([]string{"b", "c"}, DeleteAt([]string{"a", "b", "c"}, 0, 1))
	assert.Equal([]string{"c"}, DeleteAt([]string{"a", "b", "c"}, 0, 2))
	assert.Equal([]string{}, DeleteAt([]string{"a", "b", "c"}, 0, 3))
	assert.Equal([]string{}, DeleteAt([]string{"a", "b", "c"}, 0, 4))
	assert.Equal([]string{"a"}, DeleteAt([]string{"a", "b", "c"}, 1, 3))
	assert.Equal([]string{"a"}, DeleteAt([]string{"a", "b", "c"}, 1, 4))
}

func TestInsertAt(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	strs := []string{"a", "b", "c"}
	assert.Equal([]string{"1", "a", "b", "c"}, InsertAt(strs, -1, "1"))
	assert.Equal([]string{"1", "a", "b", "c"}, InsertAt(strs, 0, "1"))
	assert.Equal([]string{"a", "1", "b", "c"}, InsertAt(strs, 1, "1"))
	assert.Equal([]string{"a", "b", "1", "c"}, InsertAt(strs, 2, "1"))
	assert.Equal([]string{"a", "b", "c", "1"}, InsertAt(strs, 3, "1"))
	assert.Equal([]string{"a", "b", "c", "1"}, InsertAt(strs, 4, "1"))
	assert.Equal([]string{"1", "2", "3", "a", "b", "c"}, InsertAt(strs, 0, []string{"1", "2", "3"}...))
	assert.Equal([]string{"a", "b", "c", "1", "2", "3"}, InsertAt(strs, 3, []string{"1", "2", "3"}...))
}

func TestUpdateAt(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	assert.Equal([]string{"a", "b", "c"}, UpdateAt([]string{"a", "b", "c"}, -1, "1"))
	assert.Equal([]string{"1", "b", "c"}, UpdateAt([]string{"a", "b", "c"}, 0, "1"))
	assert.Equal([]string{"a", "b", "2"}, UpdateAt([]string{"a", "b", "c"}, 2, "2"))
	assert.Equal([]string{"a", "b", "c"}, UpdateAt([]string{"a", "b", "c"}, 3, "2"))
}

func TestUnionBy(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	testFunc := func(i int) int {
		return i / 2
	}

	result := UnionBy(testFunc, []int{0, 1, 2, 3, 4, 5}, []int{0, 2, 10})
	assert.Equal(result, []int{0, 2, 4, 10})
}

func TestMerge(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	s1 := []int{1, 2, 3, 4}
	s2 := []int{2, 3, 4, 5}
	s3 := []int{4, 5, 6}

	assert.Equal([]int{1, 2, 3, 4, 2, 3, 4, 5, 4, 5, 6}, Merge(s1, s2, s3))
	assert.Equal([]int{1, 2, 3, 4, 2, 3, 4, 5}, Merge(s1, s2))
	assert.Equal([]int{2, 3, 4, 5, 4, 5, 6}, Merge(s2, s3))
}

func TestIsAscending(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	assert.Equal(true, IsAscending([]int{1, 2, 3, 4, 5}))
	assert.Equal(false, IsAscending([]int{5, 4, 3, 2, 1}))
	assert.Equal(false, IsAscending([]int{2, 1, 3, 4, 5}))
}

func TestIsDescending(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	assert.Equal(true, IsDescending([]int{5, 4, 3, 2, 1}))
	assert.Equal(false, IsDescending([]int{1, 2, 3, 4, 5}))
	assert.Equal(false, IsDescending([]int{2, 1, 3, 4, 5}))
}

func TestIsSorted(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	assert.Equal(true, IsSorted([]int{5, 4, 3, 2, 1}))
	assert.Equal(true, IsSorted([]int{1, 2, 3, 4, 5}))
	assert.Equal(false, IsSorted([]int{2, 1, 3, 4, 5}))
}

func TestAppendIfAbsent(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	str1 := []string{"a", "b"}
	assert.Equal([]string{"a", "b"}, AppendIfAbsent(str1, "a"))
	assert.Equal([]string{"a", "b", "c"}, AppendIfAbsent(str1, "c"))
}

func TestMap(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	nums := []int{1, 2, 3, 4}
	multiplyTwo := func(i, num int) int {
		return num * 2
	}

	assert.Equal([]int{2, 4, 6, 8}, Map(nums, multiplyTwo))

	type student struct {
		name string
		age  int
	}
	students := []student{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}
	studentsOfAdd10Aage := []student{
		{"a", 11},
		{"b", 12},
		{"c", 13},
	}
	mapFunc := func(i int, s student) student {
		s.age += 10
		return s
	}

	assert.Equal(studentsOfAdd10Aage, Map(students, mapFunc))
}

func TestJoin(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	nums := []int{1, 2, 3, 4, 5}

	result1 := Join(nums, ",")
	result2 := Join(nums, "-")

	assert.Equal("1,2,3,4,5", result1)
	assert.Equal("1-2-3-4-5", result2)
}
