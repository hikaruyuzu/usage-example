package util

func CalculateLarge(v Large) int {
	return v.CalculateArea()
}

func (rectangle Rectangle) CalculateArea() int {
	return rectangle.Side * rectangle.Side
}

func (sequare Sequare) CalculateArea() int {
	return sequare.Length * sequare.Width
}
