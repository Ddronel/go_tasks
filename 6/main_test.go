package main

import "testing"

func TestGenerateRandom(t *testing.T) {
	const count = 4
	numbers := generateRandom(count)

	if cap(numbers) != 0 {
		t.Error("channel must be unbuffered")
	}

	var result []int
	for value := range numbers {
		result = append(result, value)
		if value < 0 || value >= 50 {
			t.Errorf("number = %d, want a value in [0, 50)", value)
		}
	}

	if len(result) != count {
		t.Errorf("received %d numbers, want %d", len(result), count)
	}
}

func TestGenerateRandomZeroCount(t *testing.T) {
	for value := range generateRandom(0) {
		t.Fatalf("expected no numbers, got %d", value)
	}
}
