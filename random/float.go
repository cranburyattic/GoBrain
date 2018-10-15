package random

import (
	"math/rand"
	"time"
)

// GenerateRamdomFloat64InRange between two values
func GenerateRamdomFloat64InRange(min int, max int) float64 {
	rand.Seed(time.Now().UnixNano())
	return float64(min) + rand.Float64()*(float64(max)-float64(min))
}
