package esp

import (
	"math"
)

// Neural Network
type NNet struct {
	numInput  int       // number of inputs
	numOutput int       // number of outputs
	numNeuron int       // number of neurons
	weights   []float64 // list of weights
	// sigmoid function for activation
	sigmoid func(float64) float64
}

// new neural network with no hidden neurons
func NewNNet(ni, no, nn int, resp float64) *NNet {
	return &NNet{
		numInput:  ni,
		numOutput: no,
		numNeuron: nn,
		weights:   []float64{},
		sigmoid: func(x float64) float64 {
			return 1.0 / (1.0 + math.Exp(-x/resp))
		},
	}
}

// build neural network with chromosomes
func (n *NNet) Build(c []*Chromosome) {
	n.weights = []float64{}
	for _, chrom := range c {
		gene := chrom.Gene()
		n.weights = append(n.weights, gene...)
	}
}

// update and return output
func (n *NNet) Update(inputs []float64) []float64 {
	outputs := make([]float64, n.numOutput)
	counter := 0
	for i := 0; i < n.numNeuron; i++ {
		netInput := 0.0
		for _, input := range inputs {
			netInput += input * n.weights[counter]
			counter++
		}
		signal := n.sigmoid(netInput)
		for j, _ := range outputs {
			outputs[j] += n.weights[counter] * signal
			counter++
		}
	}
	for i, output := range outputs {
		outputs[i] = n.sigmoid(output)
	}
	return outputs
}
