package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	originalSlice := generateRandomSlice(10)
	fmt.Printf("Origin slice: %v\n", originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Printf("Slice with even number: %v\n", evenSlice)

	addedSlice := addElements(evenSlice, 3)
	fmt.Printf("Slice after added element 3: %v\n", addedSlice)

	fmt.Printf("Pointer to main slice: %p\n", &addedSlice)
	copiedSlice := copySlice(addedSlice)
	fmt.Printf("Pointer to copied slice: %p\n", &copiedSlice)
	fmt.Printf("Slices value after copy: %v\n", copiedSlice)

	addedSlice[0] = -1
	fmt.Printf("Source after change: %v\n", addedSlice)
	fmt.Printf("Copy: %v\n", copiedSlice)

	removedSlice := removeElement(copiedSlice, 0)
	fmt.Printf("Slice after remove first element: %v\n", removedSlice)
}

func generateRandomSlice(count int) []int {
	slice := make([]int, count)

	for i := range slice {
		slice[i] = rand.IntN(1000)
	}

	return slice
}

func sliceExample(slice []int) []int {
	newSlice := make([]int, 0, len(slice))

	for _, v := range slice {
		if v%2 == 0 {
			newSlice = append(newSlice, v)
		}
	}

	return newSlice
}

func addElements(slice []int, el int) []int {
	return append(slice, el)
}

func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	result := make([]int, len(slice)-1)
	copy(result, slice[:index])
	copy(result[index:], slice[index+1:])

	return result
}
