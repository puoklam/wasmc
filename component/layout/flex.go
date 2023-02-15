package layout

import . "github.com/puoklam/wasmc/component"

type Flex struct {
	*Element
}

// func (f Flex) AppendChildren(children ...Component) {
// 	for _, c := range children {
// 		f.Element.Value().Call("appendChild", c.Value())
// 	}
// }

func NewFlex(dir string, opts ...ElementOption) Flex {
	display, mainAlign, crossAlign := "flex", "start", "start"
	ds := map[string]string{
		"display":         display,
		"flex-direction":  dir,
		"align-items":     crossAlign,
		"justify-content": mainAlign,
	}
	opts = append([]ElementOption{WithStyle(ds)}, opts...)
	e := NewElement(opts...)
	return Flex{e}
}
