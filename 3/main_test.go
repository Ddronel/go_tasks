package main

import (
	"maps"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name    string
		initial map[string]int
		key     string
		value   int
		want    map[string]int
	}{
		{
			name:  "nil map",
			key:   "go",
			value: 10,
			want:  map[string]int{"go": 10},
		},
		{
			name:    "empty map",
			initial: map[string]int{},
			key:     "go",
			value:   10,
			want:    map[string]int{"go": 10},
		},
		{
			name:    "new key preserves existing entries",
			initial: map[string]int{"go": 10},
			key:     "rust",
			value:   20,
			want:    map[string]int{"go": 10, "rust": 20},
		},
		{
			name:    "overwrite existing key",
			initial: map[string]int{"go": 10, "rust": 20},
			key:     "go",
			value:   30,
			want:    map[string]int{"go": 30, "rust": 20},
		},
		{
			name:  "zero value",
			key:   "zero",
			value: 0,
			want:  map[string]int{"zero": 0},
		},
		{
			name:  "empty key",
			key:   "",
			value: 10,
			want:  map[string]int{"": 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers := StringIntMap{data: tt.initial}

			numbers.Add(tt.key, tt.value)

			if !maps.Equal(numbers.data, tt.want) {
				t.Errorf("Add(%q, %d): got %v, want %v", tt.key, tt.value, numbers.data, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name    string
		initial map[string]int
		key     string
		want    map[string]int
	}{
		{
			name: "nil map",
			key:  "go",
			want: nil,
		},
		{
			name:    "empty map",
			initial: map[string]int{},
			key:     "go",
			want:    map[string]int{},
		},
		{
			name:    "remove existing key",
			initial: map[string]int{"zero": 0, "one": 1},
			key:     "one",
			want:    map[string]int{"zero": 0},
		},
		{
			name:    "remove missing key",
			initial: map[string]int{"zero": 0, "one": 1},
			key:     "two",
			want:    map[string]int{"zero": 0, "one": 1},
		},
		{
			name:    "remove only key",
			initial: map[string]int{"go": 10},
			key:     "go",
			want:    map[string]int{},
		},
		{
			name:    "remove key with zero value",
			initial: map[string]int{"zero": 0, "one": 1},
			key:     "zero",
			want:    map[string]int{"one": 1},
		},
		{
			name:    "remove empty key",
			initial: map[string]int{"": 10, "go": 20},
			key:     "",
			want:    map[string]int{"go": 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers := StringIntMap{data: tt.initial}

			numbers.Remove(tt.key)

			if !maps.Equal(numbers.data, tt.want) {
				t.Errorf("Remove(%q): got %v, want %v", tt.key, numbers.data, tt.want)
			}
		})
	}

	t.Run("remove same key twice", func(t *testing.T) {
		numbers := StringIntMap{data: map[string]int{"go": 10, "zero": 0}}
		want := map[string]int{"zero": 0}

		numbers.Remove("go")
		numbers.Remove("go")

		if !maps.Equal(numbers.data, want) {
			t.Errorf("Remove(%q) twice: got %v, want %v", "go", numbers.data, want)
		}
	})
}

func TestGet(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]int
		key       string
		wantValue int
		wantOK    bool
	}{
		{
			name: "nil map",
			key:  "go",
		},
		{
			name:  "empty map",
			input: map[string]int{},
			key:   "go",
		},
		{
			name:      "existing key",
			input:     map[string]int{"go": 10, "rust": 20},
			key:       "go",
			wantValue: 10,
			wantOK:    true,
		},
		{
			name:      "zero value",
			input:     map[string]int{"zero": 0},
			key:       "zero",
			wantValue: 0,
			wantOK:    true,
		},
		{
			name:      "negative value",
			input:     map[string]int{"negative": -10},
			key:       "negative",
			wantValue: -10,
			wantOK:    true,
		},
		{
			name:  "missing key",
			input: map[string]int{"zero": 0, "go": 10},
			key:   "missing",
		},
		{
			name:      "empty key",
			input:     map[string]int{"": 42, "go": 10},
			key:       "",
			wantValue: 42,
			wantOK:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers := StringIntMap{data: tt.input}

			gotV, gotOk := numbers.Get(tt.key)

			if gotV != tt.wantValue || gotOk != tt.wantOK {
				t.Errorf("Get(%q) = (%d, %t), want (%d, %t)", tt.key, gotV, gotOk, tt.wantValue, tt.wantOK)
			}
		})
	}
}

func TestExists(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]int
		key   string
		want  bool
	}{
		{
			name: "nil map",
			key:  "go",
		},
		{
			name:  "empty map",
			input: map[string]int{},
			key:   "go",
		},
		{
			name:  "existing key",
			input: map[string]int{"go": 10, "rust": 20},
			key:   "go",
			want:  true,
		},
		{
			name:  "key with zero value",
			input: map[string]int{"zero": 0},
			key:   "zero",
			want:  true,
		},
		{
			name:  "missing key",
			input: map[string]int{"zero": 0, "go": 10},
			key:   "missing",
		},
		{
			name:  "empty key",
			input: map[string]int{"": 0, "go": 10},
			key:   "",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers := StringIntMap{data: tt.input}

			got := numbers.Exists(tt.key)

			if got != tt.want {
				t.Errorf("Exists(%q) = %t, want %t", tt.key, got, tt.want)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]int
		want  map[string]int
	}{
		{
			name:  "usual",
			input: map[string]int{"zero": 0, "one": 1, "two": 2},
			want:  map[string]int{"zero": 0, "one": 1, "two": 2},
		},
		{
			name:  "empty",
			input: map[string]int{},
			want:  map[string]int{},
		},
		{
			name:  "nil",
			input: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers := StringIntMap{data: tt.input}

			got := numbers.Copy()

			if !maps.Equal(got, tt.want) {
				t.Fatalf("Copy(%v) = %v, want %v", tt.input, got, tt.want)
			}
			if !maps.Equal(numbers.data, tt.want) {
				t.Fatalf("Copy changed its input: got %v, want %v", tt.input, tt.want)
			}

			for i := range tt.input {
				tt.input[i]++
			}
			if !maps.Equal(got, tt.want) {
				t.Errorf("copy changed after modifying the original: got %v, want %v", got, tt.want)
			}
		})
	}
}
