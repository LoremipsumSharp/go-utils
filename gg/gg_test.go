package gg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsZero(t *testing.T) {
	var foo Foo
	assert.True(t, IsZero(foo))

	foo.Bar = "bar"
	assert.False(t, IsZero(foo))
}

func TestToType(t *testing.T) {
	// Test converting to slice
	originalSlice := []int{1, 2, 3}
	convertedSlice, ok := ToType[[]int](originalSlice)
	assert.True(t, ok)
	assert.Equal(t, originalSlice, convertedSlice)

	// Test converting to pointer
	originalPtr := &Foo{Bar: "test"}
	convertedPtr, ok := ToType[*Foo](originalPtr)
	assert.True(t, ok)
	assert.Equal(t, originalPtr, convertedPtr)

	// Test converting to struct
	originalStruct := Foo{Bar: "hello"}
	convertedStruct, ok := ToType[Foo](originalStruct)
	assert.True(t, ok)
	assert.Equal(t, originalStruct, convertedStruct)
}

func TestMustToType(t *testing.T) {
	// Test successful conversion
	originalSlice := []int{1, 2, 3}
	result := MustToType[[]int](originalSlice)
	assert.Equal(t, originalSlice, result)

	// Test panic on failure
	assert.Panics(t, func() {
		MustToType[int]("not an int")
	})
}

type Foo struct {
	Bar string
}
