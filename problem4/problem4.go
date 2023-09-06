package main

import (
	"fmt"
	"math/rand"
)

func sum_to_n_a(n int) int {
	// sums numbers 0 to n recursively.
	if n <= 0 {
		return 0
	}
	return n + sum_to_n_a(n-1)
}

func sum_to_n_b(n int, numWorkers int) int {
	resultChan := make(chan int)
	chunkSize := n / numWorkers

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if i == numWorkers-1 {
			end = n
		}

		go func(start, end int) {
			sum := 0
			for j := start; j <= end; j++ {
				sum += j
			}
			resultChan <- sum
		}(start+1, end)
	}

	totalSum := 0
	for i := 0; i < numWorkers; i++ {
		totalSum += <-resultChan
	}

	return totalSum
}

func sum_to_n_c(n int) int {
	// functional progamming?
	print(n)
	return 0
}

func main() {
	n := 100

	fmt.Println("func A: ", sum_to_n_a(n))
	fmt.Println("func B: ", sum_to_n_b(n, rand.Intn(12)))
	sum_to_n_c(n)
}
