package esp

import ()

type Neuron struct {
	inputWeights  []float64
	outputWeights []float64
}

func NewNeuron(i, o int, g []float64) *Neuron {
	return &Neuron{
		inputWeights:  g[:i],
		outputWeights: g[i:o],
	}
}
