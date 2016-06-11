package nnet

import ()

// Neural Network
type NNet struct {
	NumInputs  int       // number of inputs
	NumOutputs int       // number of outputs
	Neurons    []*Neuron // neurons in this nnet
}

func New(numInputs, numOutputs int, neurons []*Neuron) *NNet {
	return &NNet{
		NumInputs:  numInputs,
		NumOutputs: numOutputs,
		Neurons:    neurons,
	}
}

// update and return output
func (n *NNet) Update(inputs []float64) ([]float64, error) {
	outputs := make([]float64, n.NumOutputs)
	for _, neuron := range n.Neurons {
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
