package layout

import (
	"fmt"

	. "github.com/puoklam/wasmc/component"
)

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

type GridItem struct {
	*Element
}

func NewGridItem(t int, opts ...ElementOption) GridItem {
	gridCol := fmt.Sprintf("auto / span %d", t)
	ds := map[string]string{
		"grid-column": gridCol,
	}
	opts = append([]ElementOption{WithStyle(ds)}, opts...)
	e := NewElement(opts...)
	return GridItem{e}
}
