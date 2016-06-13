package esp

import (
	"math/rand"
)

// define a chromosome
type Chromosome struct {
	fitness float64
	gene    []float64
}

func NewChromosome(len int) *Chromosome {
	return &Chromosome{
		fitness: 0,
		gene: func() []float64 {
			g := make([]float64, len)
			for i, _ := range g {
				g[i] = rand.Float64()
			}
			return g
		}(),
	}
}

// get fitness
func (c *Chromosome) Fitness() float64 {
	return c.fitness
}

// get gene
func (c *Chromosome) Gene() []float64 {
	return c.gene
}
