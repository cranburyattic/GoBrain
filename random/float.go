package random

import (
	"math/rand"
	"time"
)

// GenerateRamdomFloat64InRange between two values
func GenerateRamdomFloat64InRange(min int, max int) float64 {
	return GenerateRamdomFloat64InRangeUsingSeed(min, max, time.Now().UnixNano())
}

// GenerateRamdomFloat64InRangeUsingSeed between two values with a seed
func GenerateRamdomFloat64InRangeUsingSeed(min int, max int, seed int64) float64 {
	rand.Seed(seed)
	return float64(min) + rand.Float64()*(float64(max)-float64(min))
}
