package main

import (
	"fmt"

	"github.com/cranburyattic/ml/brain"
)

func main() {

	// create the perceptron and set the weights
	p := brain.Perceptron{}
	p.SetWeights()

	// create some points to be able to train the brain
	points := [100]brain.Point{}
	var r1 int
	var r2 int

	for i := range points {

		pt := brain.Point{}
		n1, n2 := pt.SetValues()

		points[i] = pt

		r1 = r1 + n1
		r2 = r2 + n2
	}

	var correct int
	var wrong int

	for i := 0; i < 30; i++ {
		for _, pt := range points {
			inputs := []float64{pt.X, pt.Y}
			p.Train(inputs, pt.Label)
			guess := p.Guess(inputs)
			if guess == pt.Label {
				correct++
			} else {
				wrong++
			}
		}

		fmt.Println(r1, r2, correct, wrong)

		correct = 0
		wrong = 0
	}
}
