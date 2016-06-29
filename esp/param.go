package esp

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

// ESP parameter
type ESPParam struct {
	NumInput      int     // number of inputs
	NumOutput     int     // number of outputs
	NumNeuron     int     // number of neurons (+1 for bias)
	SubpSize      int     // subpopulation size
	NumGeneration int     // number of generations
	NumAvgEval    int     // number of average evaluations
	MutationRate  float64 // mutation rate
	CrossoverRate float64 // crossover rate
	Response      float64 // activation response for sigmoid function
	InitBestScore float64 // initial best score
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
		NumAvgEval:    int(fd["NUM_AVG_EVAL"]),
		MutationRate:  fd["MUTATION_RATE"],
		CrossoverRate: fd["CROSSOVER_RATE"],
		Response:      fd["RESPONSE"],
		InitBestScore: fd["INIT_BEST_SCORE"],
	}, nil
}

// print parameter
func (p *ESPParam) Show() {
	fmt.Printf("%-30s%d\n", "Number of inputs: ", p.NumInput)
	fmt.Printf("%-30s%d\n", "Number of outputs: ", p.NumOutput)
	fmt.Printf("%-30s%d\n", "Number of neurons: ", p.NumNeuron)
	fmt.Printf("%-30s%d\n", "Subpopulation size: ", p.SubpSize)
	fmt.Printf("%-30s%d\n", "Number of generations: ", p.NumGeneration)
	fmt.Printf("%-30s%d\n", "Number of avg. evaluation: ", p.NumAvgEval)
	fmt.Printf("%-30s%f\n", "Mutation rate: ", p.MutationRate)
	fmt.Printf("%-30s%f\n", "Crossover rate: ", p.CrossoverRate)
	fmt.Printf("%-30s%f\n", "Activation response: ", p.Response)
	fmt.Printf("%-30s%f\n", "Initial best score: ", p.InitBestScore)
}
