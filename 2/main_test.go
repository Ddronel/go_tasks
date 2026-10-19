package main

import (
	"slices"
	"testing"
)

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "mixed",
			input: []int{1, 2, 3, 4},
			want:  []int{2, 4},
		},
		{
			name:  "only_even",
			input: []int{2, 4, 64},
			want:  []int{2, 4, 64},
		},
		{
			name:  "only_even",
			input: []int{2, 4, 64},
			want:  []int{2, 4, 64},
		},
		{
			name:  "empty",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "negative_zero",
			input: []int{-3, -2, 0, 1},
			want:  []int{-2, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceExample(tt.input)

			if !slices.Equal(got, tt.want) {
				t.Errorf("sliceExample(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	inputSlice := []int{2, 4, 5}
	tests := []struct {
		name    string
		addedEl int
		want    []int
	}{
		{
			name:    "usual",
			addedEl: 3,
			want:    []int{2, 4, 5, 3},
		},
		{
			name:    "add_negative",
			addedEl: -2,
			want:    []int{2, 4, 5, -2},
		},
		{
			name:    "add_zero",
			addedEl: 0,
			want:    []int{2, 4, 5, -0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addElements(inputSlice, tt.addedEl)

			if !slices.Equal(got, tt.want) {
				t.Errorf(
					"addElements(%v, %d) = %v, want %v",
					inputSlice, tt.addedEl, got, tt.want,
				)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "usual",
			input: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "single_element",
			input: []int{42},
			want:  []int{42},
		},
		{
			name:  "negative_zero_and_duplicates",
			input: []int{-2, 0, -2, 4},
			want:  []int{-2, 0, -2, 4},
		},
		{
			name:  "empty",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "nil",
			input: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := copySlice(tt.input)
			if !slices.Equal(got, tt.want) {
				t.Fatalf("copySlice(%v) = %v, want %v", tt.input, got, tt.want)
			}
			if !slices.Equal(tt.input, tt.want) {
				t.Fatalf("copySlice changed its input: got %v, want %v", tt.input, tt.want)
			}

			for i := range tt.input {
				tt.input[i]++
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("copy changed after modifying the original: got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		index int
		want  []int
	}{
		{
			name:  "first",
			input: []int{10, 20, 30},
			index: 0,
			want:  []int{20, 30},
		},
		{
			name:  "middle",
			input: []int{10, 20, 30},
			index: 1,
			want:  []int{10, 30},
		},
		{
			name:  "last",
			input: []int{10, 20, 30},
			index: 2,
			want:  []int{10, 20},
		},
		{
			name:  "single_element",
			input: []int{42},
			index: 0,
			want:  []int{},
		},
		{
			name:  "remove_by_index_not_value",
			input: []int{1, 2, 1},
			index: 2,
			want:  []int{1, 2},
		},
		{
			name:  "empty",
			input: []int{},
			index: 0,
			want:  []int{},
		},
		{
			name:  "nil",
			input: nil,
			index: 0,
			want:  nil,
		},
		{
			name:  "negative_index",
			input: []int{10, 20, 30},
			index: -1,
			want:  []int{10, 20, 30},
		},
		{
			name:  "index_equal_to_length",
			input: []int{10, 20, 30},
			index: 3,
			want:  []int{10, 20, 30},
		},
		{
			name:  "index_beyond_length",
			input: []int{10, 20, 30},
			index: 10,
			want:  []int{10, 20, 30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := slices.Clone(tt.input)
			got := removeElement(tt.input, tt.index)
			if !slices.Equal(got, tt.want) {
				t.Fatalf("removeElement(%v, %d) = %v, want %v", before, tt.index, got, tt.want)
			}
			if !slices.Equal(tt.input, before) {
				t.Errorf("removeElement changed its input: got %v, want %v", tt.input, before)
			}

			if tt.index >= 0 && tt.index < len(tt.input) && len(got) > 0 {
				got[0]++
				if !slices.Equal(tt.input, before) {
					t.Errorf("modifying the result changed the input: got %v, want %v", tt.input, before)
				}
			}
		})
	}
}
