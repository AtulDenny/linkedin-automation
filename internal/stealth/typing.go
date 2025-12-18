package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func TypeHuman(el *rod.Element, text string) error {
	for _, ch := range text {

		if err := el.Input(string(ch)); err != nil {
			return err
		}

		time.Sleep(time.Duration(rand.Intn(120)+40) * time.Millisecond)

		if rand.Float64() < 0.06 {
			time.Sleep(time.Duration(rand.Intn(400)+200) * time.Millisecond)
		}
	}

	return nil
}
