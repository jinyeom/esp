package esp

import (
	"bufio"
	"errors"
	"os"
	"path"
	"strconv"
	"strings"
)

// ESP parameter
type ESPParam struct {
	NumInput      int
	NumOutput     int
	NumNeuron     int
	SubpSize      int
	NumGeneration int
	MutationRate  float64
	CrossoverRate float64
}

func NewESPParam(filename string) (*ESPParam, error) {
	if path.Ext(filename) != ".esp" {
		err := errors.New(".esp file required for parameter")
		return nil, err
	}
	f, err := os.Open(filename)
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
	return &ESPParam{
		NumInput:      int(fd["NUM_INPUT"]),
		NumOutput:     int(fd["NUM_OUTPUT"]),
		NumNeuron:     int(fd["NUM_NEURON"]),
		SubpSize:      int(fd["SUBP_SIZE"]),
		NumGeneration: int(fd["NUM_GENERATION"]),
		MutationRate:  fd["MUTATION_RATE"],
		CrossoverRate: fd["CROSSOVER_RATE"],
	}, nil
}
