package main

import (
	"slices"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name string
		s1   []string
		s2   []string
		want []string
	}{
		{
			name: "nil slices",
			want: []string{},
		},
		{
			name: "nil first slice",
			s2:   []string{"apple"},
			want: []string{},
		},
		{
			name: "nil second slice",
			s1:   []string{"apple", "banana"},
			want: []string{"apple", "banana"},
		},
		{
			name: "empty first slice",
			s1:   []string{},
			s2:   []string{"apple"},
			want: []string{},
		},
		{
			name: "empty second slice",
			s1:   []string{"apple", "banana"},
			s2:   []string{},
			want: []string{"apple", "banana"},
		},
		{
			name: "no common elements",
			s1:   []string{"apple", "banana"},
			s2:   []string{"cherry", "date"},
			want: []string{"apple", "banana"},
		},
		{
			name: "all elements excluded",
			s1:   []string{"apple", "banana"},
			s2:   []string{"banana", "apple"},
			want: []string{},
		},
		{
			name: "preserve order and duplicates",
			s1:   []string{"cherry", "banana", "apple", "cherry"},
			s2:   []string{"banana", "banana"},
			want: []string{"cherry", "apple", "cherry"},
		},
		{
			name: "exclude empty string",
			s1:   []string{"", "go", ""},
			s2:   []string{""},
			want: []string{"go"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := difference(tt.s1, tt.s2)

			if !slices.Equal(got, tt.want) {
				t.Errorf("difference(%v, %v) = %v, want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}
