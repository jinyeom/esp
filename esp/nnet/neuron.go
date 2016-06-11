package nnet

import (
	"fmt"
	"math"
)

type Neuron struct {
	inWeights  []float64
	outWeights []float64
	activation func(float64) float64
}

func NewNeuron(in, out int, c *Chromosome, af string) *Neuron {
	return &Neuron{
		inWeights:  c.Gene()[:in],
		outWeights: c.Gene()[in:out],
		activation: func(fn string) func(float64) float64 {
			switch fn {
			case "step":
				return step
			case "sigmoid":
				return sigmoid
			default:
				return nil
			}
		}(af),
	}
}

// get the neuron's output
func (n *Neuron) Output(input []float64) ([]float64, error) {
	ni := len(input)        // number of inputs
	niw := len(n.InWeights) // number of input weights

	// error check for given input size
	if ni != niw {
		err := fmt.Errorf("Unmatching inputs: %d != %d\n", ni, niw)
		return nil, err
	}

	// process signal
	signal := func() float64 {
		inputSum := 0.0
		for i, in := range input {
			inputSum += in * n.InWeights[i]
		}
		return n.Activation(inputSum)
	}()

	// get outputs
	outputs := make([]float64, len(n.OutWeights))
	for i, w := range n.OutWeights {
		outputs[i] = w * signal
	}
	return outputs, nil
}

// step function
func step(x float64) float64 {
	if x < 0.0 {
		return 0.0
	}
	return 1.0
}

// sigmoid function
func sigmoid(x float64) float64 {
	return 1.0 / (1 + math.Exp(x))
}
