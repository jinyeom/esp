package esp

// Neural Network
type NNet struct {
	numInput  int       // number of inputs
	numOutput int       // number of outputs
	neurons   []*Neuron // neurons in this nnet
}

// new neural network with no hidden neurons
func NewNNet(numInput, numOutput int) *NNet {
	return &NNet{
		numInput:  numInput,
		numOutput: numOutput,
		neurons:   nil,
	}
}

// provide hidden neurons
func (n *NNet) AddNeurons(c []*Chromosome) {
	n.neurons = make([]*Neuron, len(c))
	for i, _ := range n.neurons {
		neuron := NewNeuron(n.numInput, n.numOutput, c[i])
		n.neurons[i] = neuron
	}
}

// update and return output
func (n *NNet) Update(inputs []float64) ([]float64, error) {
	outputs := make([]float64, n.numOutput)
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
