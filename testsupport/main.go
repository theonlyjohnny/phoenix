package testsupport

import (
	"math/rand"
	"time"
)

var (
	r *rand.Rand
)

func init() {
	s := rand.NewSource(time.Now().UnixNano())
	r = rand.New(s)
}
