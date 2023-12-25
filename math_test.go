package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Range(4)
	result2 := Range(-4)
	result3 := Range(0)
	is.Equal(result1, []int{0, 1, 2, 3})
	is.Equal(result2, []int{0, -1, -2, -3})
	is.Equal(result3, []int{})
}

func TestRangeFrom(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := RangeFrom(1, 5)
	result2 := RangeFrom(-1, -5)
	result3 := RangeFrom(10, 0)
	result4 := RangeFrom(2.0, 3)
	result5 := RangeFrom(-2.0, -3)
	is.Equal(result1, []int{1, 2, 3, 4, 5})
	is.Equal(result2, []int{-1, -2, -3, -4, -5})
	is.Equal(result3, []int{})
	is.Equal(result4, []float64{2.0, 3.0, 4.0})
	is.Equal(result5, []float64{-2.0, -3.0, -4.0})
}

func TestRangeClose(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := RangeWithSteps(0, 20, 6)
	result2 := RangeWithSteps(0, 3, -5)
	result3 := RangeWithSteps(1, 1, 0)
	result4 := RangeWithSteps(3, 2, 1)
	result5 := RangeWithSteps(1.0, 4.0, 2.0)
	result6 := RangeWithSteps[float32](-1.0, -4.0, -1.0)
	is.Equal([]int{0, 6, 12, 18}, result1)
	is.Equal([]int{}, result2)
	is.Equal([]int{}, result3)
	is.Equal([]int{}, result4)
	is.Equal([]float64{1.0, 3.0}, result5)
	is.Equal([]float32{-1.0, -2.0, -3.0}, result6)
}

func TestClamp(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Clamp(0, -10, 10)
	result2 := Clamp(-42, -10, 10)
	result3 := Clamp(42, -10, 10)

	is.Equal(result1, 0)
	is.Equal(result2, -10)
	is.Equal(result3, 10)
}

func TestSum(t *testing.T) {
	is := assert.New(t)

	result1 := Sum([]float32{2.3, 3.3, 4, 5.3})
	result2 := Sum([]int32{2, 3, 4, 5})
	result3 := Sum([]uint32{2, 3, 4, 5})
	result4 := Sum([]uint32{})
	result5 := Sum([]complex128{4_4, 2_2})

	is.Equal(result1, float32(14.900001))
	is.Equal(result2, int32(14))
	is.Equal(result3, uint32(14))
	is.Equal(result4, uint32(0))
	is.Equal(result5, complex128(6_6))
}

func TestSumBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := SumBy([]float32{2.3, 3.3, 4, 5.3}, func(n float32) float32 { return n })
	result2 := SumBy([]int32{2, 3, 4, 5}, func(n int32) int32 { return n })
	result3 := SumBy([]uint32{2, 3, 4, 5}, func(n uint32) uint32 { return n })
	result4 := SumBy([]uint32{}, func(n uint32) uint32 { return n })
	result5 := SumBy([]complex128{4_4, 2_2}, func(n complex128) complex128 { return n })

	is.Equal(result1, float32(14.900001))
	is.Equal(result2, int32(14))
	is.Equal(result3, uint32(14))
	is.Equal(result4, uint32(0))
	is.Equal(result5, complex128(6_6))
}

func TestRound(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal(Round(float32(55.555), -3), float32(0))
	is.Equal(Round(float32(55.555), -2), float32(100))
	is.Equal(Round(float32(55.555), -1), float32(60))
	is.Equal(Round(float32(55.555), 0), float32(56))
	is.Equal(Round(float32(55.555), 1), float32(55.6))
	is.Equal(Round(float32(55.555), 2), float32(55.56))
	is.Equal(Round(float32(55.555), 3), float32(55.555))
	is.Equal(Round(float32(55.555), 4), float32(55.5550))

	is.Equal(Round(float64(55.555), -3), float64(0))
	is.Equal(Round(float64(55.555), -2), float64(100))
	is.Equal(Round(float64(55.555), -1), float64(60))
	is.Equal(Round(float64(55.555), 0), float64(56))
	is.Equal(Round(float64(55.555), 1), float64(55.6))
	is.Equal(Round(float64(55.555), 2), float64(55.56))
	is.Equal(Round(float64(55.555), 3), float64(55.555))
	is.Equal(Round(float64(55.555), 4), float64(55.5550))
}

func TestMax(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Max(1, 2, 3)
	result2 := Max(3, 2, 1)
	result3 := Max[int]()

	is.Equal(result1, 3)
	is.Equal(result2, 3)
	is.Equal(result3, 0)
}

func TestMaxBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MaxBy(func(item string, max string) bool {
		return len(item) > len(max)
	}, "s1", "string2", "s3")
	result2 := MaxBy(func(item string, max string) bool {
		return len(item) > len(max)
	}, "string1", "string2", "s3")
	result3 := MaxBy(func(item string, max string) bool {
		return len(item) > len(max)
	})

	is.Equal(result1, "string2")
	is.Equal(result2, "string1")
	is.Equal(result3, "")
}

func TestMin(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Min(1, 2, 3)
	result2 := Min(3, 2, 1)
	result3 := Min[int]()

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, 0)
}

func TestMinBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MinBy(func(item string, min string) bool {
		return len(item) < len(min)
	}, "s1", "string2", "s3")
	result2 := MinBy(func(item string, min string) bool {
		return len(item) < len(min)
	}, "string1", "string2", "s3")
	result3 := MinBy(func(item string, min string) bool {
		return len(item) < len(min)
	})

	is.Equal(result1, "s1")
	is.Equal(result2, "s3")
	is.Equal(result3, "")
}

func TestAverage(t *testing.T) {
	t.Parallel()

	is := assert.New(t)

	is.Equal(0, Average([]int{0, 0}))
	is.Equal(1, Average([]int{1, 1}))

	avg := Average([]float64{1.2, 1.4})
	is.Equal(1.3, Round(avg, 1))
}

func TestAbs(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	assert.Equal(0, Abs(0))
	assert.Equal(1, Abs(-1))

	assert.Equal(0.1, Abs(-0.1))

	assert.Equal(int64(1), Abs(int64(-1)))
	assert.Equal(float32(1), Abs(float32(-1)))
}
