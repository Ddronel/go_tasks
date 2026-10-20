package main

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	first := makeChannel(0, 5, 6)
	second := makeChannel(1, 2, 3, 4)
	third := makeChannel(-1, 5)

	var result []int
	for value := range merge(first, second, third) {
		result = append(result, value)
	}

	slices.Sort(result)
	want := []int{-1, 0, 1, 2, 3, 4, 5, 5, 6}
	if !slices.Equal(result, want) {
		t.Errorf("merge result = %v, want %v", result, want)
	}
}

func TestMergeEmptyChannel(t *testing.T) {
	empty := makeChannel()
	numbers := makeChannel(1, 2)

	var result []int
	for value := range merge(empty, numbers) {
		result = append(result, value)
	}

	want := []int{1, 2}
	if !slices.Equal(result, want) {
		t.Errorf("merge result = %v, want %v", result, want)
	}
}

func TestMergeNoChannels(t *testing.T) {
	for value := range merge() {
		t.Fatalf("expected no numbers, got %d", value)
	}
}

func TestMakeChannel(t *testing.T) {
	var result []int
	for value := range makeChannel(0, -2, 5, 5) {
		result = append(result, value)
	}

	want := []int{0, -2, 5, 5}
	if !slices.Equal(result, want) {
		t.Errorf("makeChannel result = %v, want %v", result, want)
	}
}
