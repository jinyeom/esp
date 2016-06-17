package esp

import ()

type ESP struct {
	NumInput  int   // number of inputs
	NumOutput int   // number of outputs
	NumNeuron int   // number of neurons
	NNetwork  *NNet // neural network

	SubpSize   int              // subpopulation size
	Population []*Subpopulation // group of subpopulations
	MRate      float64          // mutation rate
	XRate      float64          // crossover rate
}

// new esp from .esp config file

/* .esp FILE FORMAT                         */
/* ---------------------------------------- */
/* line 1: number of inputs                 */
/* line 2: number of outputs                */
/* line 3: number of neurons                */
/* line 4: subpoplation size                */
/* line 5: mutation rate (0.0 < x < 1.0)    */
/* line 6: crossover rate (0.0 < x < 1.0)   */

func New(filename string) (*ESP, error) {
	if path.Ext(filename) != ".esp" {
		return nil, errors.New("Config file is not .esp file")
	}

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	// file scanner
	fs := bufio.NewScanner(f)
	fd := make([]string, 6)
	for i := 0; i < 6 && fs.Scan(); i++ {
		fd[i] = fs.Text()
	}

	return &ESP{
		NumInput:   numInput,
		NumOutput:  numOutput,
		NumNeuron:  numNeuron,
		NNetwork:   NewNNet(numInput, numOutput, nil),
		Population: make([]*Subpopulation, numNeuron),
		MRate:      mRate,
		XRate:      xRate,
	}, nil
}

// initialize esp population
func (e *ESP) InitPopulation(size int) {
	for i := 0; i < e.NumNeuron; i++ {
		e.Population[i] = NewSubpopulation(size)
	}
}

func (e *ESP) Run() {

}
