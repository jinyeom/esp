package esp

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path"
	"strconv"
)

type ESP struct {
	NumInput   int              // number of inputs
	NumOutput  int              // number of outputs
	NumNeuron  int              // number of neurons
	NNetwork   *NNet            // neural network
	SubpSize   int              // subpopulation size
	Population []*Subpopulation // group of subpopulations
	MRate      float64          // mutation rate
	XRate      float64          // crossover rate
}

// load ESP parameter
func loadParam(filename string) (map[string]interface{}, error) {
	param := make(map[string]interface{})

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	// file scanner
	log.Printf("Loading ESP parameters: %s\n", filename)
	fs := bufio.NewScanner(f)
	fd := make([]string, 6)
	for i := 0; i < 6 && fs.Scan(); i++ {
		fd[i] = fs.Text()
	}

	numInput, err := strconv.ParseInt(fd[0], 10, 0)
	if err != nil {
		return nil, err
	}
	numOutput, err := strconv.ParseInt(fd[1], 10, 0)
	if err != nil {
		return nil, err
	}
	numNeuron, err := strconv.ParseInt(fd[2], 10, 0)
	if err != nil {
		return nil, err
	}
	subpSize, err := strconv.ParseInt(fd[3], 10, 0)
	if err != nil {
		return nil, err
	}
	mRate, err := strconv.ParseFloat(fd[4], 64)
	if err != nil {
		return nil, err
	}
	xRate, err := strconv.ParseFloat(fd[5], 64)
	if err != nil {
		return nil, err
	}

	param["numInput"] = numInput
	param["numOutput"] = numOutput
	param["numNeuron"] = numNeuron
	param["subpSize"] = subpSize
	param["mRate"] = mRate
	param["xRate"] = xRate

	return param, nil
}

func New(filename string) (*ESP, error) {
	if path.Ext(filename) != ".esp" {
		return nil, errors.New("Config file is not .esp file")
	}
	p, err := loadParam(filename)
	if err != nil {
		return nil, err
	}
	return &ESP{
		NumInput:  p["numInput"],
		NumOutput: p["numOutput"],
		NumNeuron: p["numNeuron"],
		NNetwork:  NewNNet(p["numInput"], p["numOutput"]),
		SubpSize:  p["subpSize"],
		Population: func() []*Subpopulation {
			pop := make([]*Subpopulation, p["numNeuron"])
			length := p["numInput"] + p["numOutput"]
			for i := 0; i < p["numNeuron"]; i++ {
				pop[i] = NewSubpopulation(p["subpSize"], length)
			}
			return pop
		}(),
		MRate: p["mRate"],
		XRate: p["xRate"],
	}, nil
}

func (e *ESP) Run() {

}
