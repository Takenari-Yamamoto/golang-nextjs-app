package main

import (
	"fmt"
	"time"
)

func square(n int) {
	fmt.Printf("Square of %d is %d\n", n, n*n)
}

func main() {
	// シーケンシャルに関数を実行
	go square(2)
	go square(4)
	go square(6)

	// シーケンシャルな実行が終わるのを待つ
	time.Sleep(1 * time.Second)
}
