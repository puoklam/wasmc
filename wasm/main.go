package main

import (
	"fmt"
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

	mainEle := ElementOf(doc.Call("getElementById", "main"))

	grid := NewGrid()
	gi1 := NewGridItem(6, WithText("HI"))
	gi2 := NewGridItem(6, WithText("Ho"))
	gi3 := NewGridItem(8, WithText("888"))
	gi4 := NewGridItem(2, WithText("5"))
	gi5 := NewGridItem(1, WithText("6"))
	gi6 := NewGridItem(2, WithText("7"))
	grid.Children().Append(gi1, gi2, gi3, gi4, gi5, gi6)

	fn := js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println(args[0].Get("target").Get("value"))
		return nil
	})
	input := NewInput[float64](
		"number",
		WithId("inp"),
		WithAttr(map[string]any{"value": 1.2}),
		WithListener(map[string]js.Func{"change": fn}),
	)
	// var st Setter[float64] = input
	// var vr Valuer[float64] = input

	mainEle.Children().Append(div, h3, box, circle, flex, grid, input)
	<-make(chan bool)
}
