package exer8

// TODO: The Point struct, NewPoint function, .String and .Norm methods
import "fmt"
import "math"

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	point := Point{x, y}
	return point
}

func (p Point) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func (p Point) Norm() float64 {
	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

func (p Point) Scale(t float64) {
	p.x = p.x * t
	p.y = p.y * t
}

func (p Point) Rotate(r float64) {
	p.x = p.x*math.Cos(r) - p.y*math.Sin(r)
	p.y = p.x*math.Sin(r) + p.y*math.Cos(r)
}
