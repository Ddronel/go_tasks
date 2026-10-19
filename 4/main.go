package main

import (
	"fmt"
	"slices"
)

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	uniqSlice := difference(slice1, slice2)

	fmt.Println(uniqSlice)
}

func difference(s1 []string, s2 []string) []string {
	slice := make([]string, 0, len(s1))

	for _, v := range s1 {
		if !slices.Contains(s2, v) {
			slice = append(slice, v)
		}
	}

	return slice
}
