package class

import "syscall/js"

type Class struct {
}

type Classes struct {
	js.Value
}

func (c Classes) List() string {
	return c.Value.String()
}

func (c Classes) Contains(class string) bool {
	return c.Call("contains", class).Bool()
}

func (c *Classes) Add(class string) bool {
	if c.Contains(class) {
		return false
	}
	c.Call("add", class)
	return true
}

func (c *Classes) Remove(class string) bool {
	if !c.Contains(class) {
		return false
	}
	c.Call("remove", class)
	return true
}

func (c *Classes) Toggle(class string) {
	if c.Contains(class) {
		c.Remove(class)
	} else {
		c.Add(class)
	}
}
