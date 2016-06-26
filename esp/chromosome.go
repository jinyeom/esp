package esp

import (
	"math/rand"
)

// define a chromosome
type Chromosome struct {
	evaluated int       // number of evaluation times
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
				// based on Dave Moriarty's method
				g[i] = (rand.Float64() * 12.0) - 6.0
			}
			return g
		}(),
	}
}

// uniform crossover
func UCrossover(p1, p2 *Chromosome, m float64) (c1, c2 *Chromosome) {
	g1 := p1.Gene()
	g2 := p2.Gene()
	length := len(g1)
	for i := 0; i < length; i++ {
		rn := rand.Float64()
		if rn < m {
			// swap weight
			g1[i], g2[i] = g2[i], g1[i]
		}
	}
	return &Chromosome{
			evaluated: 0,
			score:     0.0,
			gene:      g1,
		}, &Chromosome{
			evaluated: 0,
			score:     0.0,
			gene:      g2,
		}
}

// get fitness average fitness score
func (c *Chromosome) Fitness() float64 {
	if c.evaluated == 0 {
		return 0.0
	}
	return c.score / float64(c.evaluated)
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
func (c *Chromosome) Mutate(r float64) {
	for i, _ := range c.gene {
		if rand.Float64() < r {
			mut := rand.NormFloat64()
			c.gene[i] += mut
		}
	}
}
