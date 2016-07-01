package maesp

import (
	"bufio"
	"errors"
	"os"
	"path"
	"strconv"
	"strings"
)

type MAESPParam struct {
	NumInput      int     // number of inputs
	NumOutput     int     // number of outputs
	NumNeuron     int     // number of neurons
	NumAgent      int     // number of agents
	SubpSize      int     // subpopulation size
	NumGeneration int     // number of generations
	NumAvgEval    int     // number of average evaluations
	MutationRate  float64 // mutation rate
	CrossoverRate float64 // crossover rate
	Response      float64 // activation response for sigmoid function
	InitBestScore float64 // initial best score
}

func NewMAESPParam(filename string) (*MAESPParam, error) {
	if path.Ext(filename) != ".maesp" {
		err := errors.New(".maesp file required for parameter")
		return nil, err
	}
	f, err := os.Open(filaname)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	// file scanner
	fs := bufio.NewScanner(f)
	fd := make(map[string]float64)
	for fs.Scan() {
		line := fs.Text()
		tokens := strings.Split(line, "=")
		if len(tokens) == 2 {
			param := tokens[0]
			value, err := strconv.ParseFloat(tokens[1], 64)
			if err != nil {
				return nil, err
			}
			fd[param] = value
		}
	}
	return &MAESP{
		NumInput:      int(fd["NUM_INPUT"]),
		NumOutput:     int(fd["NUM_OUTPUT"]),
		NumNeuron:     int(fd["NUM_NEURON"]),
		NumAgent:      int(fd["NUM_AGENT"]),
		SubpSize:      int(fd["SUBP_SIZE"]),
		NumGeneration: int(fd["NUM_GENERATION"]),
		NumAvgEval:    int(fd["NUM_AVG_EVAL"]),
		MutationRate:  fd["MUTATION_RATE"],
		CrossoverRate: fd["CROSSOVER_RATE"],
		Response:      fd["RESPONSE"],
		InitBestScore: fd["INIT_BEST_SCORE"],
	}, nil
}
