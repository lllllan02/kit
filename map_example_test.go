package kit

import (
	"fmt"
	"sort"
	"strings"
)

func ExampleKeys() {
	kv := map[string]int{"foo": 1, "bar": 2}

	result := Keys(kv)

	sort.StringSlice(result).Sort()
	fmt.Printf("%v", result)
	// Output: [bar foo]
}

func ExampleValues() {
	kv := map[string]int{"foo": 1, "bar": 2}

	result := Values(kv)

	sort.IntSlice(result).Sort()
	fmt.Printf("%v", result)
	// Output: [1 2]
}

func ExampleValueOr() {
	kv := map[string]int{"foo": 1, "bar": 2}

	result1 := ValueOr(kv, "foo", 42)
	result2 := ValueOr(kv, "baz", 42)

	fmt.Printf("%v %v", result1, result2)
	// Output: 1 42
}

func ExamplePickBy() {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := PickBy(kv, func(key string, value int) bool {
		return value%2 == 1
	})

	fmt.Printf("%v %v %v", len(result), result["foo"], result["baz"])
	// Output: 2 1 3
}

func ExamplePickByKeys() {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := PickByKeys(kv, []string{"foo", "baz"})

	fmt.Printf("%v %v %v", len(result), result["foo"], result["baz"])
	// Output: 2 1 3
}

func ExamplePickByValues() {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := PickByValues(kv, []int{1, 3})

	fmt.Printf("%v %v %v", len(result), result["foo"], result["baz"])
	// Output: 2 1 3
}

func ExampleOmitBy() {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := OmitBy(kv, func(key string, value int) bool {
		return value%2 == 1
	})

	fmt.Printf("%v", result)
	// Output: map[bar:2]
}

func ExampleOmitByKeys() {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := OmitByKeys(kv, []string{"foo", "baz"})

	fmt.Printf("%v", result)
	// Output: map[bar:2]
}

func ExampleOmitByValues() {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := OmitByValues(kv, []int{1, 3})

	fmt.Printf("%v", result)
	// Output: map[bar:2]
}

func ExampleEntries() {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := Entries(kv)

	sort.Slice(result, func(i, j int) bool {
		return strings.Compare(result[i].Key, result[j].Key) < 0
	})
	fmt.Printf("%v", result)
	// Output: [{bar 2} {baz 3} {foo 1}]
}

func ExampleFromEntries() {
	result := FromEntries([]Entry[string, int]{
		{
			Key:   "foo",
			Value: 1,
		},
		{
			Key:   "bar",
			Value: 2,
		},
		{
			Key:   "baz",
			Value: 3,
		},
	})

	fmt.Printf("%v %v %v %v", len(result), result["foo"], result["bar"], result["baz"])
	// Output: 3 1 2 3
}

func ExampleInvert() {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := Invert(kv)

	fmt.Printf("%v %v %v %v", len(result), result[1], result[2], result[3])
	// Output: 3 foo bar baz
}

func ExampleAssign() {
	result := Assign(
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 3, "c": 4},
	)

	fmt.Printf("%v %v %v %v", len(result), result["a"], result["b"], result["c"])
	// Output: 3 1 3 4
}

func ExampleMapToSlice() {
	kv := map[int]int64{1: 1, 2: 2, 3: 3, 4: 4}

	result := MapToSlice(kv, func(k int, v int64) string {
		return fmt.Sprintf("%d_%d", k, v)
	})

	sort.StringSlice(result).Sort()
	fmt.Printf("%v", result)
	// Output: [1_1 2_2 3_3 4_4]
}
