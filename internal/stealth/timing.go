package stealth

import (
	"math/rand"
	"time"
)

func HumanDelay(minMs, maxMs int) {
	if maxMs <= minMs {
		maxMs = minMs + 50
	}
	delay := rand.Intn(maxMs-minMs) + minMs
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
