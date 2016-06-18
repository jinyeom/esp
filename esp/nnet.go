package esp

import ()

// Neural Network
type NNet struct {
	numInputs  int       // number of inputs
	numOutputs int       // number of outputs
	neurons    []*Neuron // neurons in this nnet
}

// new neural network with no hidden neurons
func NewNNet(numInputs, numOutputs int) *NNet {
	return &NNet{
		numInputs:  numInputs,
		numOutputs: numOutputs,
		neurons:    nil,
	}
}

// provide hidden neurons
func (n *NNet) AddNeurons(neurons []*Neuron) {
	n.neurons = neurons
}

// update and return output
func (n *NNet) Update(inputs []float64) ([]float64, error) {
	outputs := make([]float64, n.numOutputs)
	for _, neuron := range n.neurons {
		out, err := neuron.Output(inputs)
		if err != nil {
			return nil, err
		}
		for i, _ := range outputs {
			outputs[i] += out[i]
		}
	}
	return outputs, nil
}
