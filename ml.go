package main

import (
	"fmt"

	"github.com/cranburyattic/ml/brain"
	"github.com/cranburyattic/ml/random"
)

func trainTheBrain(p *brain.Perceptron, iterations int) {
	// create some points to be able to train the brain
	trainingPoints := [100]brain.Point{}
	var r1 int
	var r2 int

	for i := range trainingPoints {

		pt := brain.Point{}
		n1, n2 := pt.SetValues(random.GenerateRamdomFloat64InRange(0, 300),
			random.GenerateRamdomFloat64InRange(0, 300))

		trainingPoints[i] = pt

		r1 = r1 + n1
		r2 = r2 + n2
	}

	var correct int
	var incorrect int

	// train the brain
	for i := 0; i < iterations; i++ {
		for _, pt := range trainingPoints {
			inputs := []float64{pt.X, pt.Y}
			p.Train(inputs, pt.Label)
			guess := p.Guess(inputs)
			if guess == pt.Label {
				correct++
			} else {
				incorrect++
			}
		}

		//fmt.Println(r1, r2, correct, wrong)

		correct = 0
		incorrect = 0
	}

}

func main() {

	// create the perceptron and set the weights
	p := brain.Perceptron{}
	p.SetWeights()

	trainTheBrain(&p, 20)

	correct := 0
	incorrect := 0
	// test the brain
	for i := 0; i < 30; i++ {

		pt := brain.Point{}
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
