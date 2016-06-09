package nnet

import (
	"testing"
)

const (
	N_INPUTS  = 5 // number of inputs
	N_OUTPUTS = 5 // number of outputs
)

func TestNeuron(t *testing.T) error {
	// test 1: empty chromosome
	gene := []float64{}
	inputs := []float{}
	n := NewNeuron(N_INPUTS, N_OUTPUTS, gene)

	o, err := n.Output(input)
	if err != nil {
		log.Fatal(err)
		continue
	}
}
