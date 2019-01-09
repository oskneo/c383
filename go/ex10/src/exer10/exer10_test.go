package exer10

import (
	"fmt"
	"math"
	"testing"
)

func TestFib(t *testing.T) {
	n := Fib(10, 11)
	fmt.Println(n)
	// if n > 0 {
	// 	t.Error("It is wrong.")
	// }
	pt := Point{3, 4}
	TurnDouble(&pt, 3*math.Pi/2)
	fmt.Println(pt)
	tri := Triangle{Point{1, 2}, Point{-3, 4}, Point{5, -6}}
	TurnDouble(&tri, math.Pi)
	fmt.Println(tri)

	DrawCircle(40, 20, "out.png")
}
