package brain

// Point is being tested
type Point struct {
	X     float64
	Y     float64
	Label int
}

// SetValues sets random x and y values of the point
// and adds a label
func (pt *Point) SetValues(x, y float64) (int, int) {

	pt.X = x
	pt.Y = y

	if pt.X > pt.Y {
		pt.Label = 1
		return 1, 0
	} else {
		pt.Label = -1
		return 0, 1
	}
}
