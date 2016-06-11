package nnet

import (
	"math/rand"
	"testing"
)

const (
	N_INPUTS  = 5 // number of inputs
	N_OUTPUTS = 5 // number of outputs
)

// create a random gene
func RandSliceF64(n int) []float64 {
	s := make([]float64, n)
	for i, _ := range s {
		s[i] = rand.Float64()
	}
	return s
}

func TestNeuron(t *testing.T) {
	// test 1: empty chromosome
	gene := make([]float64, N_INPUTS)
	inputs := RandSliceF64(N_INPUTS)
	n := NewNeuron(N_INPUTS, N_OUTPUTS, gene, "sigmoid")

	o, err := n.Output(inputs)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%f\n", o)

}
