package util

import (
	"math/rand"
	"time"
)

//generate random number

func MyRandom(num int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return r.Intn(num)
}
