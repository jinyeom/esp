package esp

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path"
	"strconv"
)

// ESP Parameter
type ESPParam struct {
	NumInput      int
	NumOutput     int
	NumNeuron     int
	SubpSize      int
	MutationRate  float64
	CrossoverRate float64
}

type ESP struct {
	NumOutput  int              // number of outputs
	NumNeuron  int              // number of neurons
	NNetwork   *NNet            // neural network
	SubpSize   int              // subpopulation size
	Population []*Subpopulation // group of subpopulations
	MRate      float64          // mutation rate
	XRate      float64          // crossover rate
}

func New(p *Param) (*ESP, error) {
	return &ESP{
		NNetwork: NewNNet(numInput, numOutput),
		Population: func() []*Subpopulation {
			pop := make([]*Subpopulation, p["numNeuron"])
			length := p["numInput"] + p["numOutput"]
			for i := 0; i < p["numNeuron"]; i++ {
				pop[i] = NewSubpopulation(p["subpSize"], length)
			}
			return pop
		}(),
	}, nil
}

func (e *ESP) Run() {

}
