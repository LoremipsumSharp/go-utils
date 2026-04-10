package collection

import (
	"slices"
	"testing"
)

func TestSortSubset(t *testing.T) {
	tests := []struct {
		name         string
		super        []string
		sub          []string
		expected     []string
	}{
		{
			name:     "sort subset by super order",
			super:    []string{"apple", "banana", "cherry", "date"},
			sub:      []string{"date", "apple", "cherry"},
			expected: []string{"apple", "cherry", "date"},
		},
		{
			name:     "empty subset",
			super:    []string{"a", "b", "c"},
			sub:      []string{},
			expected: []string{},
		},
		{
			name:     "single element",
			super:    []string{"x", "y", "z"},
			sub:      []string{"y"},
			expected: []string{"y"},
		},
		{
			name:     "subset already sorted",
			super:    []string{"1", "2", "3", "4", "5"},
			sub:      []string{"2", "4"},
			expected: []string{"2", "4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortSubset(tt.super, tt.sub, func(s string) string { return s }, func(s string) string { return s })
			if !slices.Equal(tt.sub, tt.expected) {
				t.Errorf("got %v, want %v", tt.sub, tt.expected)
			}
		})
	}
}

func TestSortSubsetWithStructs(t *testing.T) {
	type Person struct {
		ID   int
		Name string
	}

	super := []Person{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
		{ID: 4, Name: "Diana"},
	}

	sub := []Person{
		{ID: 3, Name: "Charlie"},
		{ID: 1, Name: "Alice"},
		{ID: 4, Name: "Diana"},
	}

	expected := []Person{
		{ID: 1, Name: "Alice"},
		{ID: 3, Name: "Charlie"},
		{ID: 4, Name: "Diana"},
	}

	SortSubset(super, sub, func(p Person) int { return p.ID }, func(p Person) int { return p.ID })

	if !slices.Equal(sub, expected) {
		t.Errorf("got %v, want %v", sub, expected)
	}
}

func TestSortSubsetWithInts(t *testing.T) {
	super := []int{10, 20, 30, 40, 50}
	sub := []int{50, 20, 40}
	expected := []int{20, 40, 50}

	SortSubset(super, sub, func(i int) int { return i }, func(i int) int { return i })

	if !slices.Equal(sub, expected) {
		t.Errorf("got %v, want %v", sub, expected)
	}
}
