package random

import (
	"testing"
)

func TestRandomValueBetweenMinusOneAndOne(t *testing.T) {
	for i := 0; i < 1000; i++ {
		random := GenerateRamdomFloat64InRange(-1, 1)
		if random > 1 || random < -1 {
			t.Error("Failed")
		}
	}
}

func TestRandomValueBetweenMinusZeroAndTen(t *testing.T) {
	for i := 0; i < 1000; i++ {
		random := GenerateRamdomFloat64InRange(0, 10)
		if random > 10 || random < 0 {
			t.Error("Failed")
		}
	}
}

func TestRandomValueBetweenMinusTenAndTen(t *testing.T) {
	for i := 0; i < 1000; i++ {
		random := GenerateRamdomFloat64InRange(-10, 10)
		if random > 10 || random < -10 {
			t.Error("Failed")
		}
	}
}

func TestRandomValueBetweenValuesGivenASeedValue(t *testing.T) {

	tables := []struct {
		x int
		y int
		n float64
	}{
		{-1, 1, 0.6076896589950569},
		{0, 1, 0.8038448294975284},
		{0, 10, 8.038448294975284},
	}

	for _, table := range tables {
		random := GenerateRamdomFloat64InRangeUsingSeed(table.x, table.y, 5)
		if random != table.n {

			t.Error("Failed", random)
		}
	}
}
