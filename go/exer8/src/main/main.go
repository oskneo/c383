package main

import (
	"exer8"
	"fmt"
)

func main() {
	fmt.Printf("%d\n", exer8.Hailstone(5))
	printSlice(exer8.HailstoneSequenceAppend(5))
	printSlice(exer8.HailstoneSequenceAllocate(5))
	pt := exer8.NewPoint(3, 4)
	fmt.Println(pt)
	fmt.Println(pt.String() == "(3, 4)")
	fmt.Println(pt.Norm())
	fmt.Println(pt.Norm() == 5.0)
}

func printSlice(s []uint) {
	fmt.Printf("len=%d %v\n", len(s), s)
}
