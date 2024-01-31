package main

import "fmt"

func main() {
	numberCh := make(chan int)
	massageCh := make(chan string)

	go func(numberCh chan<- int) {
		numberCh <- 100
	}(numberCh)

	go func(massageCh chan<- string) {
		massageCh <- "One Hundread"
	}(massageCh)

	number := <-numberCh
	massage := <-massageCh

	fmt.Println(number)
	fmt.Println(massage)
}
