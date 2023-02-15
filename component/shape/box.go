package shape

import (
	"fmt"

	. "github.com/puoklam/wasmc/component"
)

type Box struct {
	*Element
}

func (b Box) Area() float64 {
	// return b.Width * b.Height
	return 0
}

func (b Box) Diagonal() float64 {
	// return math.Sqrt(b.Width*b.Width + b.Height*b.Height)
	return 0
}

func NewBox(w, h float64, opts ...ElementOption) Box {
	width := fmt.Sprintf("%fpx", w)
	height := fmt.Sprintf("%fpx", h)
	ds := map[string]string{
		"width":  width,
		"height": height,
	}
	opts = append([]ElementOption{WithStyle(ds)}, opts...)
	e := NewElement(opts...)
	return Box{e}
}
