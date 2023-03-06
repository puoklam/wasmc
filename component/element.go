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
	tag     tagOption     // initial tag
	id      idOption      // initial id
	classes classesOption // initial classes
	text    textOption    // initial text
	style   styleOption   // initial style
	attr    attrOption    // initial attributes
	lsnr    lsnrOption    // initial event listeners
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

type tagOption string
type idOption string
type classesOption []string
type textOption string
type styleOption map[string]string
type attrOption map[string]any
type lsnrOption map[string]js.Func

func (t tagOption) apply(opts *elementOptions) {
	opts.tag = t
}

func (id idOption) apply(opts *elementOptions) {
	opts.id = id
}

func (c classesOption) apply(opts *elementOptions) {
	opts.classes = append(opts.classes, c...)
}

func (t textOption) apply(opts *elementOptions) {
	opts.text = t
}

func (s styleOption) apply(opts *elementOptions) {
	if opts.style == nil {
		opts.style = make(styleOption)
	}
	for k, v := range s {
		opts.style[k] = v
	}
}

func (s styleOption) String() string {
	str := "{ "
	for k, v := range s {
		str = fmt.Sprintf("%s %s: %s;", str, k, v)
	}
	return str + " }"
}

func (a attrOption) apply(opts *elementOptions) {
	if opts.attr == nil {
		opts.attr = make(attrOption)
	}
	for k, v := range a {
		opts.attr[k] = v
	}
}

func (l lsnrOption) apply(opts *elementOptions) {
	if opts.lsnr == nil {
		opts.lsnr = make(lsnrOption)
	}
	for k, v := range l {
		opts.lsnr[k] = v
	}
}

func WithTag(t string) ElementOption {
	return tagOption(t)
}

func WithId(id string) ElementOption {
	return idOption(id)
}

func WithClasses(cs []string) ElementOption {
	return classesOption(cs)
}

func WithText(t string) ElementOption {
	return textOption(t)
}

func WithStyle(s map[string]string) ElementOption {
	style := make(styleOption)
	for k, v := range s {
		style[k] = v
	}
	return style
}

func WithAttr(a map[string]any) ElementOption {
	attr := make(attrOption)
	for k, v := range a {
		attr[k] = v
	}
	return attr
}

func WithListener(l map[string]js.Func) ElementOption {
	lsnr := make(lsnrOption)
	for k, v := range l {
		lsnr[k] = v
	}
	return lsnr
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
