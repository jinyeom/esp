package nnet

import (
	"fmt"
	"math"
)

type Neuron struct {
	iweights []float64
	oweights []float64
}

func NewNeuron(i, o int, g []float64) *Neuron {
	return &Neuron{
		iweights: g[:i],
		oweights: g[i:o],
	}
}

// get the neuron's input sum
func (n *Neuron) Input(input []float64) (float64, error) {
	ni := len(input)       // number of inputs
	niw := len(n.iweights) // number of input weights

	// error check for given input size
	if ni != niw {
		err := fmt.Errorf("Unmatching inputs: %d != %d\n", ni, niw)
		return 0.0, err
	}

	// return input sum
	return func() float64 {
		is := 0.0
		for i, in := range input {
			is += in * n.iweights[i]
		}
		return is
	}(), nil
}

// get the neuron's output values (post-activation)
func (n *Neuron) Output(input float64) []float64 {
	outputs := make([]float64, len(n.oweights))
	for i, w := range n.oweights {
		outputs[i] = w * input
	}
	return outputs
}

// create an activation function
func ActivationFunc(fn string) func(float64) float64 {
	switch fn {
	case "step":
		return step
	case "sigmoid":
		return sigmoid
	default:
		return nil
	}
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
