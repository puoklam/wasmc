package component

import (
	"syscall/js"

	"golang.org/x/exp/constraints"
)

type Shape interface {
	Area() float64
}

type Component interface {
	// Classes()
	// Render()
	JS() js.Value
	Children() *children
}

type children struct {
	s []Component // list of children
	p Component   // parent of children
}

func (cl *children) Append(cs ...Component) {
	for _, c := range cs {
		cl.p.JS().Call("appendChild", c.JS())
	}
	cl.s = cl.p.Children().s
}

func (cl *children) Slice() []Component {
	return cl.s
}

type ValueType interface {
	constraints.Ordered
}
type Valuer[T ValueType] interface {
	Value() T
}

type Setter[T ValueType] interface {
	Set(T)
}
