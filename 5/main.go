package main

import (
	"slices"
	"fmt"
)

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	ok, res := intersection(a, b)

	fmt.Printf("Slices has common elements: %t\n", ok)
	fmt.Println(res)
}

func intersection(s1 []int, s2 []int) (bool, []int){
	minLen := min(len(s1), len(s2))
	res := make([]int, 0, minLen)

	for _,v := range s1 {
		if slices.Contains(s2, v) && !slices.Contains(res, v){
			res = append(res, v)
		}
	}

	return len(res) > 0, res
}
