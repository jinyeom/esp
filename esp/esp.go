package esp

import (
	"fmt"
	"math/rand"
)

type ESP struct {
	param      *ESPParam        // ESP parameter
	network    *NNet            // neural network
	population []*Subpopulation // group of subpopulations
	bestScore  float64          // best score
	BestNNet   *NNet            // best performing neural network
}

func New(p *ESPParam) *ESP {
	return &ESP{
		param:   p,
		network: NewNNet(p.NumInput, p.NumOutput, p.NumNeuron),
		population: func() []*Subpopulation {
			pop := make([]*Subpopulation, p.NumNeuron)
			length := p.NumInput + p.NumOutput + 1
			for i := 0; i < p.NumNeuron; i++ {
				pop[i] = NewSubpopulation(p.SubpSize, length)
			}
			return pop
		}(),
		bestScore: 1000.0,
		BestNNet:  NewNNet(p.NumInput, p.NumOutput, p.NumNeuron),
	}
}

// update the best score and best performing nnet
func (e *ESP) updateBest(ns float64, c []*Chromosome) {
	if ns < e.bestScore {
		fmt.Printf("best score = %f\n", ns)
		e.bestScore = ns
		e.BestNNet = NewNNet(e.param.NumInput,
			e.param.NumOutput, e.param.NumNeuron)
		e.BestNNet.Build(c)
	}
}

// run ESP given an evaluation function
func (e *ESP) Run(evalfunc func(nn *NNet) float64) {
	indices := make([]int, e.param.NumNeuron)
	chroms := make([]*Chromosome, e.param.NumNeuron)
	numEval := e.param.NumAvgEval * e.param.NumNeuron * e.param.SubpSize
	for i := 0; i < e.param.NumGeneration; i++ {
		for j := 0; j < numEval; j++ {
			// select neurons
			for index, _ := range indices {
				// randomly select an index
				rn := rand.Intn(e.param.SubpSize)
				c := e.population[index].chromosomes[rn]
				chroms[index] = c
				indices[index] = rn
			}
			// create neural network and evaluate
			e.network.Build(chroms)
			score := evalfunc(e.network)
			// update the best neural network
			e.updateBest(score, chroms)
			for subpIndex, chromIndex := range indices {
				// evaluate each selected chromosome
				e.population[subpIndex].
					chromosomes[chromIndex].Evaluate(score)
			}
		}
		// crossover
		for i, subp := range e.population {
			p1 := subp.TSelect()
			p2 := subp.TSelect()
			parent1 := e.population[i].chromosomes[p1]
			parent2 := e.population[i].chromosomes[p2]
			child1, child2 :=
				UCrossover(parent1, parent2, e.param.CrossoverRate)
			// mutation
			child1.Mutate(e.param.MutationRate)
			child2.Mutate(e.param.MutationRate)
			// population update
			e.population[i].chromosomes[p1] = child1
			e.population[i].chromosomes[p2] = child2
		}
	}
}
