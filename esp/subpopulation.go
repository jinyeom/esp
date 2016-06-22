package esp

import (
	"math/rand"
)

// define a subpopulation
type Subpopulation struct {
	SubpSize    int           // size of subpopulation
	ChromSize   int           // size of chromosome
	Chromosomes []*Chromosome // gene pool for neurons
	Fitnesses   []float64     // list of fitness scores in order
}

func NewSubpopulation(size, length int) *Subpopulation {
	return &Subpopulation{
		SubpSize:  size,
		ChromSize: length,
		Chromosomes: func() []*Chromosome {
			c := make([]*Chromosome, size)
			for i, _ := range c {
				c[i] = NewChromosome(length)
			}
			return c
		}(),
		Fitnesses: make([]float64, size),
	}
}

// binary tournament selection
func (s *Subpopulation) TSelect() *Chromosome {
	best := rand.Intn(s.SubpSize)
	for i := 1; i < s.SubpSize; i++ {
		next := rand.Intn(s.SubpSize)
		if s.Fitnesses[next] > s.Fitnesses[best] {
			best = next
		}
	}
	return s.Chromosomes[best]
}
