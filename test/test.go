package main

import (
	"fmt"
	"github.com/jinseokYeom/esp"
	"math/rand"
)

// test
func main() {
	// new sub-population test
	rand.Seed(10000)
	size, len := 10, 10
	sp := esp.NewSubPopulation(size, len)
	fmt.Printf("%f\n", sp.GenePool)
}
