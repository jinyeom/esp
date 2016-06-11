package nnet

import (
	"testing"
)

const (
	NUM_INPUTS  = 5
	NUM_OUTPUTS = 5
)

func TestNNet(*testing.T) {
	// test 1: 0 neurons
	neurons := make([]*Neuron, 0)
	n := New(N_INPUTS, N_OUTPUTS, neurons)
	n.Update([]float64{1, 2, 3, 4, 5})
	// test 2: 1 neurons

	// test 3: n neurons
}
