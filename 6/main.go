package main

import(
	"fmt"
	"math/rand"
)

func main() {
	numbers := generateRandom(4)

	for v := range numbers {
		fmt.Println(v)
	}
}

func generateRandom(count int) <- chan int {
	numbers := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			numbers <- rand.Intn(50)
		}
		close(numbers)
	}()

	return numbers
}
