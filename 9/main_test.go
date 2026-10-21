package main

import "testing"

func TestCube(t *testing.T) {
	tests := []struct {
		name  string
		input uint8
		want  float64
	}{
		{name: "zero", input: 0, want: 0},
		{name: "one", input: 1, want: 1},
		{name: "two", input: 2, want: 8},
		{name: "three", input: 3, want: 27},
		{name: "max uint8", input: 255, want: 16581375},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cube(tt.input)

			if got != tt.want {
				t.Errorf("cube(%d) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
