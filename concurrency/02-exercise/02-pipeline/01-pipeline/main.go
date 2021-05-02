package main

import "fmt"

// TODO: Build a pipeline
// generator() -> squarer() -> print

// generator - converts a list of integers to a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func main() {
	// set up the pipeline

	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.

	for n := range square(square(generator(2, 3, 4, 5, 6))) {
		fmt.Println(n)
	}
}
