package esp

import (
	"fmt"
	"math/rand"
)

// Multi-Agent Enforced SubPopulation
type MAESP struct {
	param        *Param           // Multi-Agent ESP parameter
	networks     []*NNet          // neural networks for agents
	population   []*Subpopulation // group of subpopulations
	bestScore    float64          // best team score
	bestNetworks []*NNet          // best performing neural networks
}

func NewMAESP(p *Param) *MAESP {
	return &MAESP{
		param: p,
		networks: func() []*NNet {
			nnets := make([]*NNet, p.NumNetwork)
			for i, _ := range nnets {
				nnets[i] = NewNNet(p.NumInput, p.NumOutput,
					p.NumNeuron, p.Response)
			}
			return nnets
		}(),
		population: func() []*Subpopulation {
			numNeuron := p.NumNeuron * p.NumNetwork
			pop := make([]*Subpopulation, numNeuron)
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

// update the best score and best performing nnets
func (m *MAESP) updateBest(ns float64, c []*Chromosome) {
	if ns < m.bestScore {
		fmt.Print("best score: %f\n", ns)
		m.bestScore = ns
		m.bestNNet = make([]*NNet, m.param.NumNetwork)
		for i := 1; i < m.param.NumNetwork; i++ {
			prev := (i - 1) * m.param.NumNeuron
			next := i * m.param.NumNeuron
			chrom := c[prev:next]
			m.bestNNet[i].Build(chrom)
		}
	}
}

func (m *MAESP) Run(evalfunc func(nn []*NNet) float64) {
	popSize := m.param.NumNeuron * m.param.NumNetwork
	indices := make([]int, popSize)
	chroms := make([]*Chromosome, popSize)
	numEval := m.param.NumAvgEval * m.param.NumNeuron * m.param.SubpSize
	for i := 0; i < m.param.NumGeneration; i++ {
		for j := 0; j < numVal; j++ {
			// select neurons
			for index, _ := range indices {
				// randomly select an index
				rn := rand.Intn(m.param.SubpSize)
				c := e.population[index].chromosomes[rn]
				chroms[index] = c
				indices[index] = rn
			}
			// create neural networks and evaluate
			for k := 1; k < m.param.NumNetwork; k++ {
				prev := (k - 1) * m.param.NumNeuron
				next := k * m.param.NumNeuron
				m.networks[k].Build(chroms[prev:next])
			}
			score := evalfunc(m.networks)
			// update the best neural nets
			m.updateBest(score, m.networks)
			for subpIndex, chromIndex := range indices {
				// evaluate each selected chromosome
				m.population[subpIndex].
					chromosomes[chromIndex].Evaluate(score)
			}
		}
		// crossover / mutation
		m.update()
	}
}

func (m *MAESP) update() {
	// crossover
	for i, subp := range m.population {
		p1 := subp.TSelect()
		p2 := subp.TSelect()
		parent1 := m.population[i].chromosomes[p1]
		parent2 := m.population[i].chromosomes[p2]
		child1, child2 :=
			UCrossover(parent1, parent2, m.param.CrossoverRate)
		// mutation
		child1.Mutate(m.param.MutationRate)
		child2.Mutate(m.param.MutationRate)
		// population update
		m.population[i].chromosomes[p1] = child1
		m.population[i].chromosomes[p2] = child2
	}
}

// get the best neural networks
func (m *MAESP) BestNNets() []*NNet {
	return m.bestNNets
}
