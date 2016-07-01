package maesp

import "github.com/jinseokYeom/esp"

// Multi-Agent Enforced SubPopulation
type MAESP struct {
	param        *Param               // Multi-Agent ESP parameter
	networks     []*esp.NNet          // neural networks for agents
	population   []*esp.Subpopulation // group of subpopulations
	bestScore    float64              // best team score
	bestNetworks []*esp.NNet          // best performing neural networks
}

func New(p *esp.Param) *MAESP {
	return &MAESP{
		param: p,
		networks: func() []*esp.NNet {
			nnets := make([]*esp.NNet, p.NumNetwork)
			for i, _ := range nnets {
				nnets[i] = esp.NewNNet(p.NumInput, p.NumOutput,
					p.NumNeuron, p.Response)
			}
			return nnets
		}(),
		population: func() []*esp.Subpopulation {
			numNeuron := p.NumNeuron * p.NumNetwork
			pop := make([]*esp.Subpopulation, numNeuron)
			length := p.NumInput + p.NumOutput
			biasLen := p.NumOutput + p.NumNeuron - 1
			for i := 0; i < p.NumNetwork; i++ {
				// every first neuron is a bias
				bias := i * p.NumNeuron
				pop[bias] = NewSubpopulation(p.SubpSize, biasLen)
				for j := bias + 1; j < p.NumNeuron; j++ {
					pop[j] = NewSubpopulation(p.SubpSize, length)
				}
			}
			return pop
		}(),
		bestScore: p.InitBestScore,
		bestNNet: NewNNet(p.NumInput, p.NumOutput,
			p.NumNeuron, p.Response),
	}
}

func (m *MAESP) Run(evalfunc func(nn []*esp.NNet) float64) {
	numEval := m.param.NumAvgEval * m.param.NumNeuron * m.param.SubpSize
	for i := 0; i < m.param.NumGeneration; i++ {
		for j := 0; j < numVal; j++ {
		}
	}
}
