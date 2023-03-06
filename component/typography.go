package component

type Typography struct {
	*Element
}

func NewTypography(opts ...ElementOption) Typography {
	opts = append([]ElementOption{WithTag("p")}, opts...)
	e := NewElement(opts...)
	return Typography{e}
}
