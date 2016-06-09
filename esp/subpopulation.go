package esp

import (
	"math/rand"
)

// define a chromosome
type Chromosome struct {
	gene []float64
}

func NewChromosome(len int) *Chromosome {
	return &Chromosome{
		gene: func() {
			g := make([]float64, len)
			for i, _ := range g {
				g[i] = rand.Float64()
			}
			return g
		}(),
	}
}

// define a subpopulation
type Subpopulation struct {
	Size        int           // size of subpopulation
	ChromSize   int           // size of chromosome
	Chromosomes []*Chromosome // gene pool for neurons
}

func NewSubpopulation(s, l int) *Subpopulation {
	return &Subpopulation{
		Size:      s,
		ChromSize: l,
		Chromosomes: func() []*Chromosome {
			c := make([]*Chromosome, s)
			for i, _ := range gp {
				c[i] = NewChromosome(l)
			}
			return c
		}(),
	}
}

// one point crossover
func (s *Subpopulation) Crossover1P(c1, c2 int) {
	cut := rand.Intn(s.Size)
	for i := cut; i < s.Size; i++ {
		g1 := s.Chromosomes[c1].gene[i]
		g2 := s.Chromosomes[c2].gene[i]

		s.Chromosomes[c1].gene[i] = g2
		s.Chromosomes[c2].gene[i] = g1
	}
}
