package maesp

import "github.com/jinseokYeom/esp"

// Multi-Agent Enforced SubPopulation
type MAESP struct {
	param        *MAESPParam          // Multi-Agent ESP parameter
	networks     []*esp.NNet          // neural networks for agents
	population   []*esp.Subpopulation // group of subpopulations
	bestScore    float64              // best team score
	bestNetworks []*esp.NNet          // best performing neural networks
}

func New() *MAESP {
	return &MAESP{}
}
