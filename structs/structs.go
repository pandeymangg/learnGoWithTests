package structs

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

type triangle struct {
	base   float64
	height float64
}

func (r rectangle) Area() float64 {
	return r.height * r.width
}

func (c circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (t triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

type Shape interface {
	Area() float64
}

func Perimeter(rect rectangle) float64 {
	return 2 * (rect.width + rect.height)
}
