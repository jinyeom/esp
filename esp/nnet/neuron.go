package nnet

import (
	"fmt"
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

func (n *Neuron) Output(input []float64) ([]float64, error) {
	ni := len(input)       // number of inputs
	niw := len(n.iweights) // number of input weights
	now := len(n.oweights) // number of output weights

	// error check for given input size
	if ni != niw {
		err := fmt.Errorf("Unmatching inputs: %d != %d\n", ni, niw)
		return nil, err
	}

	outputs := make([]float64, now)
	for i, _ := range outputs {
		for j, in := range input {
			outputs[i] += in * n.iweights[j]
		}
	}

	return outputs, nil
}
