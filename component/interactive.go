package component

import (
	"syscall/js"
)

type Button struct {
	*Element
	// Disabled bool
}

func NewButton(opts ...ElementOption) Button {
	opts = append([]ElementOption{WithTag("button")}, opts...)
	e := NewElement(opts...)
	return Button{e}
}

type Input[T ValueType] struct {
	*Element
	ref T
}

func NewInput[T ValueType](t string, opts ...ElementOption) Input[T] {
	opts = append([]ElementOption{
		WithTag("input"),
		WithAttr(map[string]any{"type": "number"}),
	}, opts...)
	e := NewElement(opts...)
	var ref T
	return Input[T]{e, ref}
}

func (i Input[T]) Value() T {
	v := i.JS().Get("value")

	switch any(i.ref).(type) {
	case float32, float64:
		v = i.JS().Get("valueAsNumber")
	}

	var r any
	switch v.Type() {
	case js.TypeString:
		r = v.String()
	case js.TypeNumber:
		r = v.Float()
	case js.TypeBoolean:
		r = v.Bool()
	case js.TypeUndefined:
		// TODO: undefined
	}
	return r.(T)
}

func (i Input[T]) Set(v T) {
	i.JS().Set("value", v)
}
