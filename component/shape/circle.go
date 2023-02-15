package shape

import (
	"fmt"

	. "github.com/puoklam/wasmc/component"
)

type Circle struct {
	*Element
}

func (c Circle) Area() float64 {
	// return math.Pi * c.Radius * c.Radius
	return 0
}

func NewCircle(r float64, opts ...ElementOption) Circle {
	dia := fmt.Sprintf("%fpx", r*2)
	// default style
	ds := map[string]string{
		"width":         dia,
		"height":        dia,
		"border-radius": "50%",
	}
	opts = append([]ElementOption{WithStyle(ds)}, opts...)
	e := NewElement(opts...)
	return Circle{e}
}
