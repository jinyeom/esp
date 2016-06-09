package evolution

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSubpopulation(t *testing.T) {
	// new sub-population test
	rand.Seed(10000)
	size, len := 10, 10

	sp := NewSubpopulation(size, len)

	for _, c := range sp.Chromosomes {
		fmt.Printf("%f\n", c.Gene())
	}
}
