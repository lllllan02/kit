package kit

import "fmt"

func ExampleTernary() {
	result := Ternary(true, "a", "b")
	fmt.Printf("%v", result)

	// Output: a
}

func ExampleIf() {
	result1 := If(true, 1).
		ElseIf(false, 2).
		Else(3)
	fmt.Printf("%v\n", result1)

	result2 := If(false, 1).
		ElseIf(true, 2).
		Else(3)
	fmt.Printf("%v\n", result2)

	result3 := If(false, 1).
		ElseIf(false, 2).
		Else(3)
	fmt.Printf("%v\n", result3)

	// Output:
	// 1
	// 2
	// 3
}

func ExampleSwitch() {
	result1 := Switch[int, string](1).
		Case(1, "1").
		Case(2, "2").
		Default("3")
	fmt.Printf("%v\n", result1)

	result2 := Switch[int, string](2).
		Case(1, "1").
		Case(2, "2").
		Default("3")
	fmt.Printf("%v\n", result2)

	result3 := Switch[int, string](42).
		Case(1, "1").
		Case(2, "2").
		Default("3")
	fmt.Printf("%v\n", result3)

	// Output:
	// 1
	// 2
	// 3
}
