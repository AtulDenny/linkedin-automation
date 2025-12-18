package browser

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func Launch() (*rod.Browser, error) {
	rand.Seed(time.Now().UnixNano())

	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
	}

	l := launcher.New().
		Headless(false).
		Leakless(false).
		Set("disable-blink-features", "AutomationControlled").
		Set("disable-infobars").
		Set("no-sandbox").
		Set("disable-dev-shm-usage").
		Set("user-agent", userAgents[rand.Intn(len(userAgents))])

	u, err := l.Launch()
	if err != nil {
		return nil, err
	}

	browser := rod.New().ControlURL(u)
	err = browser.Connect()
	if err != nil {
		return nil, err
	}

	return browser, nil
}
