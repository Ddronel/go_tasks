package main

import "fmt"

func main() {
	unums := []uint8{1, 2, 3}
	first := make(chan uint8)
	second := make(chan float64)

	go func() {
		for _, n := range unums {
			first <- n
		}
		close(first)
	}()

	go func() {
		for unum := range first {
			second <- cube(unum)
		}

		close(second)
	}()

	for unum := range second {
		fmt.Println(unum)
	}
}

func cube(u uint8) float64 {
	v := float64(u)

	return v * v * v
}
