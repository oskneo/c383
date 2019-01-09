package main

import (
	"exer10"
	"fmt"
)

func main() {
	fmt.Println(exer10.Fib(40, 34))

}

func printSlice(s []uint) {
	fmt.Printf("len=%d %v\n", len(s), s)
}
