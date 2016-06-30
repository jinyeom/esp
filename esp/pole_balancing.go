/*  pole_balancing.go
    Jinseok Yeom
    June 2016

    **This pole balancing test experiment is based on David Moriarty's
    experiment in June 1994 from http://nn.cs.utexas.edu/?sanepolebalancing.

*/

package esp

import "testing"

// physics sim constants
const (
	GRAVITY   = 9.8
	CART_MASS = 1.0
	POLE_MASS = 0.1
	LENGTH    = 0.5
	FORCE_MAG = 10.0
	TAU       = 0.02
)

var (
	TOTAL_MASS       = CART_MASS + POLE_MASS
	POLE_MASS_LENGTH = POLE_MASS * LENGTH
)

func poleBalancing() {

}

func TestPoleBalancing(t *testing.T) {

}
