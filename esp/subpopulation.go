package esp

import (
	"math/rand"
)

// define a subpopulation
type Subpopulation struct {
	subpSize    int           // size of subpopulation
	chromSize   int           // size of chromosome
	chromosomes []*Chromosome // gene pool for neurons
	fitnesses   []float64     // list of fitness scores in order
}

func NewSubpopulation(size, length int) *Subpopulation {
	return &Subpopulation{
		subpSize:  size,
		chromSize: length,
		chromosomes: func() []*Chromosome {
			c := make([]*Chromosome, size)
			for i, _ := range c {
				c[i] = NewChromosome(length)
			}
			return c
		}(),
		fitnesses: make([]float64, size),
	}
}

// binary tournament selection (return index)
func (s *Subpopulation) TSelect() int {
	best := rand.Intn(s.subpSize)
	for i := 1; i < s.subpSize; i++ {
		next := rand.Intn(s.subpSize)
		nfit := s.chromosomes[next].Fitness()
		bfit := s.chromosomes[best].Fitness()
		if nfit < bfit {
			best = next
		}
	}
	return best
}
