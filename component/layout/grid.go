package layout

import . "github.com/puoklam/wasmc/component"

type Grid struct {
	*Element
}

func NewGrid(opts ...ElementOption) Grid {
	display, tmpCol := "grid", "repeat(12, [col-start] 1fr)"
	ds := map[string]string{
		"display":               display,
		"grid-template-columns": tmpCol,
	}
	opts = append([]ElementOption{WithStyle(ds)}, opts...)
	e := NewElement(opts...)
	return Grid{e}
}
