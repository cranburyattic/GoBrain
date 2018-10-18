package brain

import (
	"testing"
)

func TestShouldSetLabelCorrectlyForXAndY(t *testing.T) {
	tables := []struct {
		x float64
		y float64
		n int
	}{
		{-1, 1, -1},
		{0, 1, -1},
		{1, 0, 1},
		{1, -1, 1},
		{0, 0, -1},
		{3, 3, -1},
	}

	for _, table := range tables {
		p := Point{}
		p.SetValues(table.x, table.y)
		if p.Label != table.n {
			t.Errorf("x=%v,y=%v has the incorrect Label=%v", table.x, table.y, p.Label)
		}
	}
}

func TestShoulReturnValuePairsForXAndY(t *testing.T) {
	tables := []struct {
		x  float64
		y  float64
		r1 int
		r2 int
	}{
		{-1, 1, 0, 1},
		{0, 1, 0, 1},
		{1, 0, 1, 0},
		{1, -1, 1, 0},
		{0, 0, 0, 1},
		{3, 3, 0, 1},
	}

	for _, table := range tables {
		p := Point{}
		r1, r2 := p.SetValues(table.x, table.y)
		if r1 != table.r1 {
			t.Errorf("x=%v,y=%v has the incorrect r1 value=%v", table.x, table.y, r1)
		}
		if r2 != table.r2 {
			t.Errorf("x=%v,y=%v has the incorrect r2 value=%v", table.x, table.y, r2)
		}
	}
}
