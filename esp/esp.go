package esp

import "math/rand"

type ESP struct {
	param      *Param           // ESP parameter
	network    *NNet            // neural network
	population []*Subpopulation // group of subpopulations
	bestScore  float64          // best score
	best       []*Chromosome    // best performing chromosomes
}

func New(p *Param) *ESP {
	return &ESP{
		param: p,
		network: NewNNet(p.NumInput, p.NumOutput,
			p.NumNeuron, p.Response),
		population: func() []*Subpopulation {
			pop := make([]*Subpopulation, p.NumNeuron)
			length := p.NumInput + p.NumOutput
			biasLen := p.NumOutput + p.NumNeuron - 1
			pop[0] = NewSubpopulation(p.SubpSize, biasLen)
			for i := 1; i < p.NumNeuron; i++ {
				pop[i] = NewSubpopulation(p.SubpSize, length)
			}
			return pop
		}(),
		bestScore: p.InitBestScore,
		best:      make([]*Chromosome, p.NumNeuron),
	}
}

// update the best score and best performing nnet
func (e *ESP) updateBest(ns float64, c []*Chromosome) {
	if ns < e.bestScore {
		//fmt.Printf("best score: %f\n", ns)
		e.bestScore = ns
		copy(e.best, c)
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
		// crossover/mutation
		e.update()
	}
}

// update population states
func (e *ESP) update() {
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

// get the best neural network
func (e *ESP) Best() []*Chromosome {
	return e.best
}
