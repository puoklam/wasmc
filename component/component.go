package component

import "syscall/js"

type Shape interface {
	Area() float64
}

type Component interface {
	// Classes()
	// Render()
	Value() js.Value
	Children() *children
}

type children struct {
	s []Component // list of children
	p Component   // parent of children
}

func (cl *children) Append(cs ...Component) {
	for _, c := range cs {
		cl.p.Value().Call("appendChild", c.Value())
	}
	cl.s = cl.p.Children().s
}

func (cl *children) Slice() []Component {
	return cl.s
}
