package esp

import ()

type ESP struct {
	numInput  int
	numOutput int

	network    *NNet
	population []*Subpopulation
}

// new ESP given number of neurons
func New(n int) *ESP {
	return &ESP{
		population: make([]*Subpopulation, n),
	}
}
