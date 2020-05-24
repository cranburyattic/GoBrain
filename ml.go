package main

import (
	"fmt"
	"github.com/cranburyattic/GoBrain/brain"
	"github.com/cranburyattic/GoBrain/random"
)

func trainTheBrain(p *brain.Perceptron, iterations int) {
	// create some points to be able to train the brain
	trainingPoints := [100]brain.Point{}

	for i := range trainingPoints {

		pt := brain.Point{}
		pt.SetValues(random.GenerateRamdomFloat64InRange(0, 300),
			random.GenerateRamdomFloat64InRange(0, 300))

		trainingPoints[i] = pt
	}

	// train the brain
	for i := 0; i < iterations; i++ {
		for _, pt := range trainingPoints {
			inputs := []float64{pt.X, pt.Y}
			p.Train(inputs, pt.Label)
		}
	}

}

func main() {

	// create the perceptron and set the starting weights
	p := brain.NewPerceptron()
	p.SetWeights(1.0)

	// actually train the brain - comment out to see the effect of not training the brain
	trainTheBrain(p, 20)

	correct := 0
	incorrect := 0
	// test the brain
	for i := 0; i < 30; i++ {

		pt := brain.NewPoint()
		pt.SetValues(random.GenerateRamdomFloat64InRange(0, 300),
			random.GenerateRamdomFloat64InRange(0, 300))

		inputs := []float64{pt.X, pt.Y}
		guess := p.Guess(inputs)
		if guess == pt.Label {
			correct++
		} else {
			incorrect++
		}

	}
	fmt.Printf("Correct %v : Incorrect %v\n", correct, incorrect)
}
