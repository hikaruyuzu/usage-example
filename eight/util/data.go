package util

type Large interface {
	CalculateArea() int
}

type Rectangle struct {
	Side int
}

type Sequare struct {
	Width  int
	Length int
}
