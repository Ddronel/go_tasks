package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestPrintTypes(t *testing.T) {
	tests := []struct {
		name   string
		values []any
		want   string
	}{
		{
			name: "assignment_example",
			values: []any{
				int(42), int(052), int(0x2A),
				float64(3.14), "Golang", true, complex64(1 + 2i),
			},
			want: "int\nint\nint\nfloat64\nstring\nbool\ncomplex64\n",
		},
		{
			name:   "nil_input",
			values: nil,
			want:   "",
		},
		{
			name:   "empty_input",
			values: []any{},
			want:   "",
		},
		{
			name:   "nil_interface_and_typed_nil",
			values: []any{nil, (*int)(nil)},
			want:   "<nil>\n*int\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer
			if err := printTypes(&output, tt.values); err != nil {
				t.Fatalf("printTypes() returned an unexpected error: %v", err)
			}
			if got := output.String(); got != tt.want {
				t.Errorf("printTypes(%v) output = %q, want %q", tt.values, got, tt.want)
			}
		})
	}
}

type errorWriter struct {
	err error
}

func (w errorWriter) Write(_ []byte) (int, error) {
	return 0, w.err
}

func TestPrintTypesWriterError(t *testing.T) {
	wantErr := errors.New("write failed")
	gotErr := printTypes(errorWriter{err: wantErr}, []any{42})
	if !errors.Is(gotErr, wantErr) {
		t.Errorf("printTypes() error = %v, want %v", gotErr, wantErr)
	}
}

func TestJoinValues(t *testing.T) {
	tests := []struct {
		name   string
		values []any
		want   string
	}{
		{
			name: "assignment_example",
			values: []any{
				int(42), int(052), int(0x2A),
				float64(3.14), "Golang", true, complex64(1 + 2i),
			},
			want: "4242423.14Golangtrue(1+2i)",
		},
		{
			name:   "nil_input",
			values: nil,
			want:   "",
		},
		{
			name:   "zero_values",
			values: []any{0, float64(0), "", false, complex64(0)},
			want:   "00false(0+0i)",
		},
		{
			name:   "order_and_duplicates_without_spaces",
			values: []any{3, 1, 3},
			want:   "313",
		},
		{
			name:   "negative_values",
			values: []any{-42, float64(-3.5), false, complex64(1 - 2i)},
			want:   "-42-3.5false(1-2i)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := joinValues(tt.values); got != tt.want {
				t.Errorf("joinValues(%v) = %q, want %q", tt.values, got, tt.want)
			}
		})
	}
}

func TestHashRunes(t *testing.T) {
	tests := []struct {
		name  string
		input []rune
		salt  string
		want  string
	}{
		{
			name:  "assignment_example",
			input: []rune("4242423.14Golangtrue(1+2i)"),
			salt:  "go-2024",
			want:  "53f2f60ac6c41389d3ed3d84d88d8c2860bf8981c677be18243a6f35a6b6a1b3",
		},
		{
			name:  "nil_input",
			input: nil,
			salt:  "go-2024",
			want:  "66802df107aace17871a5b610ff9eb11706e13477bb24e93966ca80671c0fac6",
		},
		{
			name:  "even_rune_count",
			input: []rune("abcd"),
			salt:  "go-2024",
			want:  "eef2e40f2b77de8474f4dec20ca90bd2e84b9eb3ddb332bd53f57ce69bb784ae",
		},
		{
			name:  "odd_rune_count",
			input: []rune("abcde"),
			salt:  "go-2024",
			want:  "1b612f3b227f152c2297f870c5ab43bc4a36bba3e2cb8a89d0993b9a88555967",
		},
		{
			name:  "single_multibyte_rune",
			input: []rune("я"),
			salt:  "go-2024",
			want: "b11ddc04d514226c29d631ffef8c38436c377767a7284523ef2dadb829195a7b",
		},
		{
			name:  "empty_salt",
			input: []rune("abc"),
			salt:  "",
			want:  "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashRunes(tt.input, tt.salt); got != tt.want {
				t.Errorf("hashRunes(%q, %q) = %q, want %q", string(tt.input), tt.salt, got, tt.want)
			}
		})
	}
}
