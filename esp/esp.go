package esp

import ()

type ESP struct {
	network    *NNet
	population []*Subpopulation

	selection  func()
	crossover  func()
	mutation   func()
	evaluation func(...float64) float64
}

func New() *ESP {
	return &ESP{}
}

func Selection()
