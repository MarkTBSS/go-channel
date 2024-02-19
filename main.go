package main

import "fmt"

func main() {
	numberCh := make(chan int)
	massageCh := make(chan string, 2)

	go func(numberCh chan<- int) {
		numberCh <- 100
	}(numberCh)

	go func(massageCh chan<- string) {
		massageCh <- "One Hundread"
		massageCh <- "Two Hundread"
	}(massageCh)

	number := <-numberCh
	massage := <-massageCh
	massage2 := <-massageCh

	fmt.Println(number)
	fmt.Println(massage)
	fmt.Println(massage2)
}
