package main

import "fmt"

func square(numbers []int, ch chan int) {
	for _, n := range numbers {
		ch <- n * n
	}
	close(ch)
}

func main() {
	numbers := []int{2, 4, 6}
	ch := make(chan int)

	go square(numbers, ch)

	// チャネルからデータを受信
	for value := range ch {
		fmt.Println(value)
	}
}
