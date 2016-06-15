package esp

import ()

// define a subpopulation
type Subpopulation struct {
	SubpSize    int           // size of subpopulation
	ChromSize   int           // size of chromosome
	Chromosomes []*Chromosome // gene pool for neurons
}

func NewSubpopulation(s, l int) *Subpopulation {
	return &Subpopulation{
		SubpSize:  s,
		ChromSize: l,
		Chromosomes: func() []*Chromosome {
			c := make([]*Chromosome, s)
			for i, _ := range c {
				c[i] = NewChromosome(l)
			}
			return c
		}(),
	}
}

// fitness-proportionate selection
func (s *Subpopulation) FPSelect() *Chromosome {
}

// tournament selection
func (s *Subpopulation) TSelect() *Chromosome {
}

// one point crossover
func (s *Subpopulation) Crossover1P(c1, c2 int) {
}

// two point crossover

// uniform crossover
