package esp

import ()

type ESP struct {
	numInput   int
	numOutput  int
	network    *NNet
	population []*Subpopulation
}

// new ESP given number of neurons
func New(numInput, numNeuron, numOutput int) *ESP {
	return &ESP{
		numInput:   numInput,
		numOutput:  numOutput,
		network:    nil,
		population: make([]*Subpopulation, numNeuron),
	}
}

// set mutation rate
func MutationRate() {

}

// set crossover rate
func CrossoverRate() {

}

//
