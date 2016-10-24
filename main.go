package main

import "fmt"

func ExampleReadClosedChannel() {
	ch := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)

	for i := 0; i < 10; i++ {
		v := <-ch
		fmt.Println("Read", v)
	}
	v := <-ch
	fmt.Println("Read", v)
}

func main() {
	ExampleReadClosedChannel()
}
