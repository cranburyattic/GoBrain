package brain

import (
	"github.com/cranburyattic/ml/random"
)

// Point is being tested
type Point struct {
	X     float64
	Y     float64
	Label int
}

// SetValues sets random x and y values of the point
// and adds a label
func (pt *Point) SetValues() (int, int) {

	pt.X = random.GenerateRamdomFloat64InRange(0, 300)
	pt.Y = random.GenerateRamdomFloat64InRange(0, 300)

	if pt.X > pt.Y {
		pt.Label = 1
		return 1, 0
	} else {
		pt.Label = -1
		return 0, 1
	}
}
