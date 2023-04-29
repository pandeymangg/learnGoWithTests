package structs

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rectangle) Area() float64 {
	return r.height * r.width
}

func (c circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

type Shape interface {
	Area() float64
}

func Perimeter(rect rectangle) float64 {
	return 2 * (rect.width + rect.height)
}
