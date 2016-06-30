/*  pole_balancing_test.go
    Jinseok Yeom
    June 2016

    **This pole balancing test experiment is based on David Moriarty's
    experiment in June 1994 from http://nn.cs.utexas.edu/?sanepolebalancing.

*/

package esp

import (
	"math/rand"
	"testing"
)

// physics sim constants
const (
	X_LIM        = 2.4    // x position limit [-2.4, 2.4]
	DX_LIM       = 1.0    // x velocity limit [-1.0, 1.0]
	TH_LIM       = 0.2    // theta limit [-0.2, 0.2]
	DTH_LIM      = 1.5    // angular velocity limit [-1.5, 1.5]
	GRAVITY      = 9.8    // gravity constant
	CART_MASS    = 1.0    // mass of the cart
	POLE_MASS    = 0.1    // mass of the pole
	LENGTH       = 0.5    // half length of pole
	FORCE_MAG    = 10.0   // force applied to the cart
	TAU          = 0.02   // seconds between state updates
	MAX_TIME     = 120000 // given time for each test
	RANDOM_START = true   // true if start randomly
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
		inputs[0] = (rand.Float64()%4800.0)/1000.0 - X_LIM
		inputs[1] = (rand.Float64()%2000.0)/1000.0 - DX_LIM
		inputs[2] = (rand.Float64()%400.0)/1000.0 - TH_LIM
		inputs[3] = (rand.Float64()%3000.0)/1000.0 - DTH_LIM
	}
	// play the game
	for i := 0; i < MAX_TIME; i++ {
		outputs := nn.Update(inputs)

		// check for failure; if so, return steps
	}
}

func TestPoleBalancing(t *testing.T) {

}
