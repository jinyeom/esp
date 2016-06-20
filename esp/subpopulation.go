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

// fitness-proportionate selection (not recommended)
func (s *Subpopulation) FPSelect() *Chromosome {
	best := 0.0
	for i := 0; i < s.SubpSize; i++ {
		score := s.Chromosomes[i].Fitness()
		s.Fitnesses[i] = score
		if score > best {
			best = score
		}
	}
	// stochastic acceptance
	for {
		i := rand.Intn(s.SubpSize)
		r := s.Fitnesses[i] / best
		if rand.Float64() < r {
			return s.Chromosomes[i]
		}
	}
}

// binary tournament selection
func (s *Subpopulation) TSelect() *Chromosome {
	best := s.Fitnesses[rand.Intn(s.SubpSize)]
	for i := 1; i < s.SubpSize; i++ {
		next := s.Fitnesses[rand.Intn(s.SubpSize)]
		if next > best {
			best = next
		}
	}
	return s.Chromosomes[best]
}

// n tournament selection
func (s *Subpopulation) TSelectn(n int) *Chromosome {

}

// one point crossover
func (s *Subpopulation) Crossover1P(c1, c2 int) {
}

// two point crossover
func (s *Subpopulation) Crossover2P(c1, c2 int) {
}

// uniform crossover
func (s *Subpopulation) UCrossover(c1, c2 int) {
}
