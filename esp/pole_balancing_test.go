/*  pole_balancing_test.go
    Jinseok Yeom
    June 2016

    **This pole balancing test experiment is based on David Moriarty's
    experiment in June 1994 from http://nn.cs.utexas.edu/?sanepolebalancing.

    Evaluation in this experiment is with MAX_TIME - time, not time itself,
    meaning, the lower the score is, the better.
*/

package esp

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

// physics sim constants
const (
	X_LIM        = 2.4      // x position limit [-2.4, 2.4]
	DX_LIM       = 1.0      // x velocity limit [-1.0, 1.0]
	TH_LIM       = 0.2      // theta limit [-0.2, 0.2]
	DTH_LIM      = 1.5      // angular velocity limit [-1.5, 1.5]
	GRAVITY      = 9.8      // gravity constant
	CART_MASS    = 1.0      // mass of the cart
	POLE_MASS    = 0.1      // mass of the pole
	LENGTH       = 0.5      // half length of pole
	FORCE_MAG    = 10.0     // force applied to the cart
	TAU          = 0.02     // seconds between state updates
	MAX_TIME     = 120000.0 // given time for each test
	RANDOM_START = false    // true if start randomly
)

var (
	TOTAL_MASS       = CART_MASS + POLE_MASS
	POLE_MASS_LENGTH = POLE_MASS * LENGTH
)

func poleBalancing(nn *NNet) float64 {
	// inputs[0]: (x) cart position in meters
	// inputs[1]: (x_dot) cart velocity
	// inputs[2]: (theta) pole angle in radians
	// inputs[3]: (theta_dot) pole angular velocity
	inputs := make([]float64, 4)
	if RANDOM_START {
		inputs[0] = float64(rand.Int31()%4800)/1000.0 - X_LIM
		inputs[1] = float64(rand.Int31()%2000)/1000.0 - DX_LIM
		inputs[2] = float64(rand.Int31()%400)/1000.0 - TH_LIM
		inputs[3] = float64(rand.Int31()%3000)/1000.0 - DTH_LIM
	}
	// play the game
	t := 0.0
	for t < MAX_TIME {
		outputs := nn.Update(inputs)
		y := (outputs[0] <= outputs[1])
		// apply action to the simulated cart-pole
		inputs = cartPole(y, inputs)
		// check for failure; if so, return steps
		x := inputs[0]  // x position
		th := inputs[2] // theta
		if x < -X_LIM || x > X_LIM ||
			th < -TH_LIM || th > TH_LIM {
			return MAX_TIME - t
		}
		t++
	}
	return MAX_TIME - t
}

// cart-pole simulation
func cartPole(action bool, inputs []float64) []float64 {
	force := FORCE_MAG
	if action {
		force = -FORCE_MAG
	}
	th := inputs[2]
	dth := inputs[3]
	cosTh := math.Cos(th)
	sinTh := math.Sin(th)
	temp := (force + POLE_MASS_LENGTH*
		dth*dth*sinTh) / TOTAL_MASS
	// angular acceleration
	ath := (GRAVITY*sinTh - cosTh*temp) /
		(LENGTH * (4.0/3.0 - POLE_MASS*cosTh*cosTh/TOTAL_MASS))
	// x acceleration
	ax := temp - POLE_MASS_LENGTH*ath*cosTh/TOTAL_MASS

	// update states
	inputs[0] += TAU * inputs[1]
	inputs[1] += TAU * ax
	inputs[2] += TAU * inputs[3]
	inputs[3] += TAU * ath
	return inputs
}

func TestPoleBalancing(t *testing.T) {
	s := time.Now().UnixNano()
	rand.Seed(s)
	param, err := NewParam("poletest.esp")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Seed: %d\n", s)
	param.Show()

	e := New(param)
	e.Run(poleBalancing)

	// test the best neural network
	nn := e.BestNNet()
	bestScore := poleBalancing(nn)
	fmt.Printf("Best time: %f\n", MAX_TIME-bestScore)
}
