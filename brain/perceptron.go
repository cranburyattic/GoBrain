package brain

const LEARNING_RATE = 0.1

// Perceptron is the brain of the system
type Perceptron struct {
	weights []float64
}

func NewPerceptron() *Perceptron {
	return &Perceptron{}
}

// SetWeights sets up the starting weights for the Perceptron
func (p *Perceptron) SetWeights(startingWeight float64) {

	weights := make([]float64, 2)

	for i := range weights {
		weights[i] = startingWeight
	}
	p.weights = weights
}

// GetWeights returns the weights value array for inspection
func (p *Perceptron) GetWeights() []float64 {

	return p.weights
}

// Guess takes an input array containing two point values
// and returns the guess -1 or 1
func (p *Perceptron) Guess(inputs []float64) int {

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
		p.weights[i] += float64(error) * inputs[i] * LEARNING_RATE
	}
}

// the activation function
func sign(n float64) int {

	if n >= 0 {
		return 1
	}

	return -1
}
