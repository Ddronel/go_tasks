package main

import (
	"fmt"
	"sync"
)

func main() {
	first := makeChannel(0, 5, 6)
	second := makeChannel(1, 2, 3, 4)

	result := merge(first, second)

	for value := range result {
		fmt.Println(value)
	}
}

func merge(chans ...<-chan int) <-chan int {
	common := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range chans {
		wg.Add(1)

		go func(input <-chan int) {
			defer wg.Done()

			for v := range input {
				common <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(common)
	}()

	return common
}

func makeChannel(values ...int) <-chan int {
	numbers := make(chan int)

	go func() {
		for _, v := range values {
			numbers <- v
		}

		close(numbers)
	}()

	return numbers
}
