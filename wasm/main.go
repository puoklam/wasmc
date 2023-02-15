package main

import (
	"syscall/js"

	. "github.com/puoklam/wasmc/component"
	. "github.com/puoklam/wasmc/component/layout"
	. "github.com/puoklam/wasmc/component/shape"
)

// func CL() js.Func {
// 	return js.FuncOf(func(this js.Value, args []js.Value) any {
// 		doc := js.Global().Get("document")
// 		h := doc.Call("getElementById", "h")
// 		c := h.Get("classList")
// 		return c
// 	})
// }

// func main() {
// 	ch := make(chan struct{}, 0)
// 	js.Global().Set("cl", CL())
// 	<-ch
// }

func main() {
	doc := js.Global().Get("document")

	div := NewElement(WithText("hello"))
	h3 := NewElement(WithTag("h3"), WithClasses([]string{"p", "o"}))
	box := NewBox(
		300,
		200,
		WithText("hello"),
		WithStyle(map[string]string{"background-color": "blue"}),
	)
	circle := NewCircle(
		100,
		WithStyle(map[string]string{"background-color": "green"}),
	)
	flex := NewFlex("row")
	item1 := NewBox(
		200,
		100,
		WithStyle(map[string]string{"background-color": "red"}),
	)
	item2 := NewBox(
		300,
		100,
		WithStyle(map[string]string{"background-color": "yellow"}),
	)
	flex.Children().Append(item1, item2)

	mainEle := NewElementFromJS(doc.Call("getElementById", "main"))
	mainEle.Children().Append(div, h3, box, circle, flex)
}
