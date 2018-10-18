package brain

import (
	"testing"
)

func TestGuessShouldReturnCorrectValueGivenWeightAndInputs(t *testing.T) {
	tables := []struct {
		w float64
		i []float64
		r int
	}{
		{0.5, []float64{2, 1}, 1},
		{0.5, []float64{1, 2}, 1},
		{0.5, []float64{3, 3}, 1},
		{-0.5, []float64{2, 1}, -1},
		{-0.5, []float64{1, 2}, -1},
		{-0.5, []float64{3, 3}, -1},
	}

	for _, table := range tables {
		p := Perceptron{}
		p.SetWeights(table.w)
		res := p.Guess(table.i)
		if res != table.r {
			t.Errorf("Failed for %v %v %v", table.w, table.i, table.r)
		}
	}
}
