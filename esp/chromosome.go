package esp

import (
	"math/rand"
)

// define a chromosome
type Chromosome struct {
	evaluated float64   // number of evaluation times
	score     float64   // total fitness score
	gene      []float64 // input and output weights
}

func NewChromosome(len int) *Chromosome {
	return &Chromosome{
		evaluated: 0,
		score:     0,
		gene: func() []float64 {
			g := make([]float64, len)
			for i, _ := range g {
				g[i] = rand.Float64()
			}
			return g
		}(),
	}
}

// get fitness average fitness score
func (c *Chromosome) Fitness() float64 {
	if c.evaluated == 0 {
		return 0.0
	}
	return c.score / c.evaluated
}

// set fitness
func (c *Chromosome) Evaluate(f float64) {
	c.evaluated++
	c.score += f
}

// get gene
func (c *Chromosome) Gene() []float64 {
	return c.gene
}

// mutation with Gaussian Convolution
func (c *Chromosome) Mutate() {

}
