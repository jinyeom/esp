package esp

import ()

// Neural Network
type NNet struct {
	inputs  int       // number of inputs
	outputs int       // number of outputs
	neurons []*Neuron // neurons in this nnet
}

func NewNNet(i, o int, n []*Neuron) *NNet {
	return &NNet{
		inputs:  i,
		outputs: o,
		neurons: n,
	}
}
