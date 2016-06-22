package esp

import (
	"fmt"
	"math"
)

type Neuron struct {
	inWeights  []float64
	outWeights []float64
}

func NewNeuron(in, out int, c *Chromosome) *Neuron {
	return &Neuron{
		inWeights:  c.Gene()[:in],
		outWeights: c.Gene()[in : in+out],
	}
}

// sigmoid function for activation
func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(x))
}

// get the neuron's output
func (n *Neuron) Output(input []float64) ([]float64, error) {
	ni := len(input)        // number of inputs
	niw := len(n.inWeights) // number of input weights
	// error check for given input size
	if ni != niw {
		err := fmt.Errorf("Unmatching inputs: %d != %d\n", ni, niw)
		return nil, err
	}
	// process signal
	signal := func() float64 {
		inputSum := 0.0
		for i, in := range input {
			inputSum += in * n.inWeights[i]
		}
		return sigmoid(inputSum)
	}()
	// get outputs
	outputs := make([]float64, len(n.outWeights))
	for i, w := range n.outWeights {
		outputs[i] = w * signal
	}
	return outputs, nil
}
