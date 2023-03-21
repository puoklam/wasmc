package component

import (
	"fmt"
	"math/rand"
	"strings"
	"syscall/js"
	"time"
)

// global stylesheet
var doc js.Value
var ss js.Value

func init() {
	// class name generator
	rand.Seed(time.Now().UnixNano())
	// create global css stylesheet
	doc = js.Global().Get("document")
	ssEle := doc.Call("createElement", "style")
	ssEle.Set("type", "text/css")
	doc.Call("getElementsByTagName", "head").Index(0).Call("appendChild", ssEle)
	ss = ssEle.Get("sheet")
}

type elementOptions struct {
	tag     string             // initial tag
	id      string             // initial id
	classes []string           // initial classes
	text    string             // initial text
	style   styleOption        // initial style
	attr    map[string]any     // initial attributes
	lsnr    map[string]js.Func // initial event listeners
}

func newEleOpts() elementOptions {
	return elementOptions{
		tag:     "div",
		classes: []string{},
		text:    "",
		style:   map[string]string{},
	}
}

type Element struct {
	value js.Value
}

func (e Element) JS() js.Value {
	return e.value
}

func (e Element) Children() *children {
	c := e.value.Get("children")
	cLen := c.Get("length").Int()
	ch := make([]Component, 0, cLen)
	for i := 0; i < cLen; i++ {
		ch = append(ch, ElementOf(c.Index(i)))
	}
	return &children{ch, e}
}

type ElementOption interface {
	apply(*elementOptions)
}

type styleOption map[string]string

func (s styleOption) String() string {
	str := "{ "
	for k, v := range s {
		str = fmt.Sprintf("%s %s: %s;", str, k, v)
	}
	return str + " }"
}

type optionFunc func(*elementOptions)

func (f optionFunc) apply(opts *elementOptions) {
	f(opts)
}

func WithTag(t string) ElementOption {
	var fn optionFunc = func(opts *elementOptions) {
		opts.tag = t
	}
	return fn
}

func WithId(id string) ElementOption {
	var fn optionFunc = func(opts *elementOptions) {
		opts.id = id
	}
	return fn
}

func WithClasses(cs []string) ElementOption {
	var fn optionFunc = func(opts *elementOptions) {
		opts.classes = append(opts.classes, cs...)
	}
	return fn
}

func WithText(t string) ElementOption {
	var fn optionFunc = func(opts *elementOptions) {
		opts.text = t
	}
	return fn
}

func WithStyle(s map[string]string) ElementOption {
	var fn optionFunc = func(opts *elementOptions) {
		if opts.style == nil {
			opts.style = make(styleOption)
		}
		for k, v := range s {
			opts.style[k] = v
		}
	}
	return fn
}

func WithAttr(a map[string]any) ElementOption {
	var fn optionFunc = func(opts *elementOptions) {
		if opts.attr == nil {
			opts.attr = make(map[string]any)
		}
		for k, v := range a {
			opts.attr[k] = v
		}
	}
	return fn
}

func WithListener(l map[string]js.Func) ElementOption {
	var fn optionFunc = func(opts *elementOptions) {
		if opts.lsnr == nil {
			opts.lsnr = make(map[string]js.Func)
		}
		for k, v := range l {
			opts.lsnr[k] = v
		}
	}
	return fn
}

func NewElement(opts ...ElementOption) *Element {
	eleOpts := newEleOpts()
	for _, opt := range opts {
		opt.apply(&eleOpts)
	}

	ele := doc.Call("createElement", string(eleOpts.tag))

	ele.Set("id", string(eleOpts.id))
	ele.Set("innerText", string(eleOpts.text))
	ele.Set("className", strings.Join(eleOpts.classes, " "))

	// create new class with random name and append to element's class list
	cl := ele.Get("classList")
	cn := randStringRunes(5)
	cssText := fmt.Sprintf(".%s %s", cn, eleOpts.style)
	ss.Call("insertRule", cssText)
	cl.Call("add", cn)
	// style := ele.Get("style")
	// for k, v := range eleOpts.style {
	// 	style.Call("setProperty", k, v)
	// }
	for k, v := range eleOpts.attr {
		ele.Call("setAttribute", k, v)
	}
	for k, v := range eleOpts.lsnr {
		ele.Call("addEventListener", k, v)
	}

	return &Element{ele}
}

func ElementOf(v js.Value) *Element {
	return &Element{v}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
