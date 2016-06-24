package esp

import (
	"math"
)

const BIAS = -1.0

// Neural Network
type NNet struct {
	numInput  int       // number of inputs
	numOutput int       // number of outputs
	numNeuron int       // number of neurons
	weights   []float64 // list of weights
}

// new neural network with no hidden neurons
func NewNNet(numInput, numOutput, numNeuron int) *NNet {
	return &NNet{
		numInput:  numInput,
		numOutput: numOutput,
		numNeuron: numNeuron,
		weights:   []float64{},
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

// sigmoid function
func sigmoid(x, resp float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x/resp))
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
		signal := sigmoid(netInput, 0.05)
		for j, _ := range outputs {
			outputs[j] += n.weights[counter] * signal
			counter++
		}
	}
	for i, output := range outputs {
		outputs[i] = sigmoid(output, 0.05)
	}
	return outputs
}
