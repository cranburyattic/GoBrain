package brain

import (
	"github.com/cranburyattic/ml/random"
)

// Perceptron is the brain of the system
type Perceptron struct {
	weights [2]float64
}

// SetWeights intializes random weights for the inputs
func (p *Perceptron) SetWeights() {
	for i := range p.weights {
		p.weights[i] = random.GenerateRamdomFloat64InRange(-1, 1)
	}
}

// Guess takes an input array containing two point values
// and retuns the guess -1 or 1
func (p Perceptron) Guess(inputs []float64) int {

	var sum float64

	for i := range p.weights {
		sum += inputs[i] * p.weights[i]
	}

	output := sign(sum)

	return output
}

// Train the Brain with a known data set
func (p *Perceptron) Train(inputs []float64, target int) {
	guess := p.Guess(inputs)
	error := target - guess

	// Tune all the weights
	for i := range p.weights {
		p.weights[i] += float64(error) * inputs[i] * 0.1
	}
}

// the activation function
func sign(n float64) int {
	if n >= 0 {
		return 1
	}
	return -1

}
