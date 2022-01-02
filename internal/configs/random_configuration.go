package configs

import (
	"math/rand"
	"time"
)

// ConfigureGlobalRandom Ensures that random values will differ based on run
func ConfigureGlobalRandom() {
	rand.Seed(time.Now().UnixNano())
}
