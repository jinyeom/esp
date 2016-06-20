package esp

import ()

type ESP struct {
	Param      *ESPParam        // ESP parameter
	NNetwork   *NNet            // neural network
	Population []*Subpopulation // group of subpopulations
}

func New(p *ESPParam) *ESP {
	return &ESP{
		Param:    p,
		NNetwork: NewNNet(p.NumInput, p.NumOutput),
		Population: func() []*Subpopulation {
			pop := make([]*Subpopulation, p.NumNeuron)
			length := p.NumInput + p.NumOutput
			for i := 0; i < p.NumNeuron; i++ {
				pop[i] = NewSubpopulation(p.SubpSize, length)
			}
			return pop
		}(),
	}
}
