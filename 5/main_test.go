package main

import (
	"slices"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name   string
		s1     []int
		s2     []int
		wantOK bool
		want   []int
	}{
		{
			name:   "task example",
			s1:     []int{65, 3, 58, 678, 64},
			s2:     []int{64, 2, 3, 43},
			wantOK: true,
			want:   []int{3, 64},
		},
		{
			name:   "nil slices",
			wantOK: false,
			want:   []int{},
		},
		{
			name:   "nil first slice",
			s2:     []int{1, 2},
			wantOK: false,
			want:   []int{},
		},
		{
			name:   "nil second slice",
			s1:     []int{1, 2},
			wantOK: false,
			want:   []int{},
		},
		{
			name:   "empty first slice",
			s1:     []int{},
			s2:     []int{1, 2},
			wantOK: false,
			want:   []int{},
		},
		{
			name:   "empty second slice",
			s1:     []int{1, 2},
			s2:     []int{},
			wantOK: false,
			want:   []int{},
		},
		{
			name:   "no common elements",
			s1:     []int{1, 2},
			s2:     []int{3, 4},
			wantOK: false,
			want:   []int{},
		},
		{
			name:   "same elements in different order",
			s1:     []int{3, 1, 2},
			s2:     []int{2, 3, 1},
			wantOK: true,
			want:   []int{3, 1, 2},
		},
		{
			name:   "duplicates in both slices",
			s1:     []int{3, 3, 64, 3, 64},
			s2:     []int{64, 64, 3, 3},
			wantOK: true,
			want:   []int{3, 64},
		},
		{
			name:   "zero is the only common element",
			s1:     []int{0, 1},
			s2:     []int{0, 2},
			wantOK: true,
			want:   []int{0},
		},
		{
			name:   "negative values",
			s1:     []int{-5, 0, -1, 7},
			s2:     []int{-1, -5, 2},
			wantOK: true,
			want:   []int{-5, -1},
		},
		{
			name:   "different lengths",
			s1:     []int{8},
			s2:     []int{1, 2, 8, 3, 4},
			wantOK: true,
			want:   []int{8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOK, got := intersection(tt.s1, tt.s2)

			if gotOK != tt.wantOK {
				t.Errorf("intersection(%v, %v): bool = %t, want %t", tt.s1, tt.s2, gotOK, tt.wantOK)
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("intersection(%v, %v): elements = %v, want %v (in any order)", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}
