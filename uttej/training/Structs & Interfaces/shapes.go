package shape

type Rectangle struct {
	length  float64
	breadth float64
}

func (r Rectangle) Area() float64 {

	return r.length * r.breadth
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func Perimeter(l float64, b float64) float64 {
	return 2 * (l + b)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.length * rectangle.breadth
}
