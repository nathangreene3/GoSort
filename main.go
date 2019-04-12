package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
}
