package exer10

import "fmt"

type Triangle struct {
	A, B, C Point
}

type Transformable interface {
	Scale(s float64)
	Rotate(r float64)
}

func TurnDouble(t Transformable, angle float64) {
	t.Scale(2)
	t.Rotate(angle)
}

func (t Triangle) String() string {
	return fmt.Sprintf("[%s %s %s]", t.A, t.B, t.C)
}

func (t *Triangle) Scale(s float64) {
	t.A.Scale(s)
	t.B.Scale(s)
	t.C.Scale(s)
}

func (t *Triangle) Rotate(r float64) {
	t.A.Rotate(r)
	t.B.Rotate(r)
	t.C.Rotate(r)
}
