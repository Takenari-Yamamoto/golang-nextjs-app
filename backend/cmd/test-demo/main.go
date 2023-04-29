package main

import "fmt"

type Fethcer interface {
	FetchData() (string, error)
}

func GetData(f Fethcer) (string, error) {
	return f.FetchData()
}

func main() {
	fmt.Println("Hello, world!")
}
